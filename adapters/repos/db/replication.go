//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package db

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/go-openapi/strfmt"
	"github.com/weaviate/weaviate/entities/additional"
	"github.com/weaviate/weaviate/entities/multi"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/entities/storobj"
	"github.com/weaviate/weaviate/usecases/objects"
	"github.com/weaviate/weaviate/usecases/replica"
	"github.com/weaviate/weaviate/usecases/replica/hashtree"
)

type Replicator interface {
	ReplicateObject(ctx context.Context, shardName, requestID string,
		object *storobj.Object) replica.SimpleResponse
	ReplicateObjects(ctx context.Context, shardName, requestID string,
		objects []*storobj.Object) replica.SimpleResponse
	ReplicateUpdate(ctx context.Context, shard, requestID string,
		doc *objects.MergeDocument) replica.SimpleResponse
	ReplicateDeletion(ctx context.Context, shardName, requestID string,
		uuid strfmt.UUID) replica.SimpleResponse
	ReplicateDeletions(ctx context.Context, shardName, requestID string,
		uuids []strfmt.UUID, dryRun bool, schemaVersion uint64) replica.SimpleResponse
	ReplicateReferences(ctx context.Context, shard, requestID string,
		refs []objects.BatchReference) replica.SimpleResponse
	CommitReplication(shard,
		requestID string) interface{}
	AbortReplication(shardName,
		requestID string) interface{}
}

func (db *DB) ReplicateObject(ctx context.Context, class,
	shard, requestID string, object *storobj.Object,
) replica.SimpleResponse {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.ReplicateObject(ctx, shard, requestID, object)
}

func (db *DB) ReplicateObjects(ctx context.Context, class,
	shard, requestID string, objects []*storobj.Object, schemaVersion uint64,
) replica.SimpleResponse {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.ReplicateObjects(ctx, shard, requestID, objects, schemaVersion)
}

func (db *DB) ReplicateUpdate(ctx context.Context, class,
	shard, requestID string, mergeDoc *objects.MergeDocument,
) replica.SimpleResponse {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.ReplicateUpdate(ctx, shard, requestID, mergeDoc)
}

func (db *DB) ReplicateDeletion(ctx context.Context, class,
	shard, requestID string, uuid strfmt.UUID,
) replica.SimpleResponse {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.ReplicateDeletion(ctx, shard, requestID, uuid)
}

func (db *DB) ReplicateDeletions(ctx context.Context, class,
	shard, requestID string, uuids []strfmt.UUID, dryRun bool, schemaVersion uint64,
) replica.SimpleResponse {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.ReplicateDeletions(ctx, shard, requestID, uuids, dryRun, schemaVersion)
}

func (db *DB) ReplicateReferences(ctx context.Context, class,
	shard, requestID string, refs []objects.BatchReference,
) replica.SimpleResponse {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.ReplicateReferences(ctx, shard, requestID, refs)
}

func (db *DB) CommitReplication(class,
	shard, requestID string,
) interface{} {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return nil
	}

	return index.CommitReplication(shard, requestID)
}

func (db *DB) AbortReplication(class,
	shard, requestID string,
) interface{} {
	index, pr := db.replicatedIndex(class)
	if pr != nil {
		return *pr
	}

	return index.AbortReplication(shard, requestID)
}

func (db *DB) replicatedIndex(name string) (idx *Index, resp *replica.SimpleResponse) {
	if !db.StartupComplete() {
		return nil, &replica.SimpleResponse{Errors: []replica.Error{
			*replica.NewError(replica.StatusNotReady, name),
		}}
	}

	if idx = db.GetIndex(schema.ClassName(name)); idx == nil {
		return nil, &replica.SimpleResponse{Errors: []replica.Error{
			*replica.NewError(replica.StatusClassNotFound, name),
		}}
	}
	return
}

func (i *Index) writableShard(name string) (ShardLike, *replica.SimpleResponse) {
	localShard, err := i.getOrInitLocalShard(context.Background(), name)
	if err != nil {
		return nil, &replica.SimpleResponse{Errors: []replica.Error{
			{Code: replica.StatusShardNotFound, Msg: name},
		}}
	}
	if localShard.isReadOnly() {
		return nil, &replica.SimpleResponse{Errors: []replica.Error{{
			Code: replica.StatusReadOnly, Msg: name,
		}}}
	}
	return localShard, nil
}

func (i *Index) ReplicateObject(ctx context.Context, shard, requestID string, object *storobj.Object) replica.SimpleResponse {
	localShard, pr := i.writableShard(shard)
	if pr != nil {
		return *pr
	}
	return localShard.preparePutObject(ctx, requestID, object)
}

func (i *Index) ReplicateUpdate(ctx context.Context, shard, requestID string, doc *objects.MergeDocument) replica.SimpleResponse {
	localShard, pr := i.writableShard(shard)
	if pr != nil {
		return *pr
	}
	return localShard.prepareMergeObject(ctx, requestID, doc)
}

func (i *Index) ReplicateDeletion(ctx context.Context, shard, requestID string, uuid strfmt.UUID) replica.SimpleResponse {
	localShard, pr := i.writableShard(shard)
	if pr != nil {
		return *pr
	}
	return localShard.prepareDeleteObject(ctx, requestID, uuid)
}

func (i *Index) ReplicateObjects(ctx context.Context, shard, requestID string, objects []*storobj.Object, schemaVersion uint64) replica.SimpleResponse {
	localShard, pr := i.writableShard(shard)
	if pr != nil {
		return *pr
	}
	return localShard.preparePutObjects(ctx, requestID, objects)
}

func (i *Index) ReplicateDeletions(ctx context.Context, shard, requestID string, uuids []strfmt.UUID, dryRun bool, schemaVersion uint64) replica.SimpleResponse {
	localShard, pr := i.writableShard(shard)
	if pr != nil {
		return *pr
	}
	return localShard.prepareDeleteObjects(ctx, requestID, uuids, dryRun)
}

func (i *Index) ReplicateReferences(ctx context.Context, shard, requestID string, refs []objects.BatchReference) replica.SimpleResponse {
	localShard, pr := i.writableShard(shard)
	if pr != nil {
		return *pr
	}
	return localShard.prepareAddReferences(ctx, requestID, refs)
}

func (i *Index) CommitReplication(shard, requestID string) interface{} {
	localShard, err := i.getOrInitLocalShard(context.Background(), shard)
	// TODO-RAFT no error response on error?
	if err != nil {
		return nil
	}
	return localShard.commitReplication(context.Background(), requestID, &i.backupMutex)
}

func (i *Index) AbortReplication(shard, requestID string) interface{} {
	localShard, err := i.getOrInitLocalShard(context.Background(), shard)
	if err != nil {
		return replica.SimpleResponse{Errors: []replica.Error{
			{Code: replica.StatusShardNotFound, Msg: shard},
		}}
	}
	return localShard.abortReplication(context.Background(), requestID)
}

func (i *Index) IncomingFilePutter(ctx context.Context, shardName,
	filePath string,
) (io.WriteCloser, error) {
	localShard, err := i.getOrInitLocalShard(context.Background(), shardName)
	if err != nil {
		return nil, fmt.Errorf("shard %q does not exist locally", shardName)
	}

	return localShard.filePutter(ctx, filePath)
}

func (i *Index) IncomingCreateShard(ctx context.Context, className string, shardName string) error {
	if _, err := i.getOrInitLocalShard(ctx, shardName); err != nil {
		return fmt.Errorf("incoming create shard: %w", err)
	}
	return nil
}

func (i *Index) IncomingReinitShard(ctx context.Context,
	shardName string,
) error {
	shard, err := i.getOrInitLocalShard(ctx, shardName)
	if err != nil {
		return fmt.Errorf("shard %q does not exist locally", shardName)
	}

	return shard.reinit(ctx)
}

func (s *Shard) filePutter(ctx context.Context,
	filePath string,
) (io.WriteCloser, error) {
	// TODO: validate file prefix to rule out that we're accidentally writing
	// into another shard
	finalPath := filepath.Join(s.Index().Config.RootPath, filePath)
	dir := path.Dir(finalPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("create parent folder for %s: %w", filePath, err)
	}

	f, err := os.Create(finalPath)
	if err != nil {
		return nil, fmt.Errorf("open file %q for writing: %w", filePath, err)
	}

	return f, nil
}

func (s *Shard) reinit(ctx context.Context) error {
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown shard: %w", err)
	}

	if err := s.initNonVector(ctx, nil); err != nil {
		return fmt.Errorf("reinit non-vector: %w", err)
	}

	if s.hasTargetVectors() {
		if err := s.initTargetVectors(ctx); err != nil {
			return fmt.Errorf("reinit vector: %w", err)
		}
	} else {
		if err := s.initLegacyVector(ctx); err != nil {
			return fmt.Errorf("reinit vector: %w", err)
		}
	}

	s.initCycleCallbacks()
	s.initDimensionTracking()

	return nil
}

// OverwriteObjects if their state didn't change in the meantime
// It returns nil if all object have been successfully overwritten
// and otherwise a list of failed operations.
func (i *Index) OverwriteObjects(ctx context.Context,
	shard string, updates []*objects.VObject,
) ([]replica.RepairResponse, error) {
	result := make([]replica.RepairResponse, 0, len(updates)/2)
	s, err := i.getOrInitLocalShard(ctx, shard)
	if err != nil {
		return nil, fmt.Errorf("shard %q not found locally", shard)
	}
	for i, u := range updates {
		// Just in case but this should not happen
		data := u.LatestObject
		if data == nil || data.ID == "" {
			msg := fmt.Sprintf("received nil object or empty uuid at position %d", i)
			result = append(result, replica.RepairResponse{Err: msg})
			continue
		}
		// valid update
		found, err := s.ObjectByID(ctx, data.ID, nil, additional.Properties{})
		var curUpdateTime int64 // 0 means object doesn't exist on this node
		if found != nil {
			curUpdateTime = found.LastUpdateTimeUnix()
		}
		r := replica.RepairResponse{ID: data.ID.String(), UpdateTime: curUpdateTime}
		switch {
		case err != nil:
			r.Err = "not found: " + err.Error()
		case curUpdateTime == u.StaleUpdateTime:
			// the stored object is not the most recent version. in
			// this case, we overwrite it with the more recent one.
			err := s.PutObject(ctx, storobj.FromObject(data, u.Vector, u.Vectors))
			if err != nil {
				r.Err = fmt.Sprintf("overwrite stale object: %v", err)
			}
		case curUpdateTime != data.LastUpdateTimeUnix:
			// object changed and its state differs from recent known state
			r.Err = "conflict"
		}

		if r.Err != "" { // include only unsuccessful responses
			result = append(result, r)
		}
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}

func (i *Index) IncomingOverwriteObjects(ctx context.Context,
	shardName string, vobjects []*objects.VObject,
) ([]replica.RepairResponse, error) {
	return i.OverwriteObjects(ctx, shardName, vobjects)
}

func (i *Index) DigestObjects(ctx context.Context,
	shardName string, ids []strfmt.UUID,
) (result []replica.RepairResponse, err error) {
	result = make([]replica.RepairResponse, len(ids))
	s, err := i.getOrInitLocalShard(ctx, shardName)
	if err != nil {
		return nil, fmt.Errorf("shard %q not found locally", shardName)
	}

	multiIDs := make([]multi.Identifier, len(ids))
	for j := range multiIDs {
		multiIDs[j] = multi.Identifier{ID: ids[j].String()}
	}

	objs, err := s.MultiObjectByID(ctx, multiIDs)
	if err != nil {
		return nil, fmt.Errorf("shard objects digest: %w", err)
	}

	for j := range objs {
		if objs[j] == nil {
			deleted, err := s.WasDeleted(ctx, ids[j])
			if err != nil {
				return nil, err
			}
			result[j] = replica.RepairResponse{
				ID:      ids[j].String(),
				Deleted: deleted,
				// TODO: use version when supported
				Version: 0,
			}
		} else {
			result[j] = replica.RepairResponse{
				ID:         objs[j].ID().String(),
				UpdateTime: objs[j].LastUpdateTimeUnix(),
				// TODO: use version when supported
				Version: 0,
			}
		}
	}

	return
}

func (i *Index) IncomingDigestObjects(ctx context.Context,
	shardName string, ids []strfmt.UUID,
) (result []replica.RepairResponse, err error) {
	return i.DigestObjects(ctx, shardName, ids)
}

func (i *Index) HashTreeLevel(ctx context.Context,
	shardName string, level int, discriminant *hashtree.Bitset,
) (digests []hashtree.Digest, err error) {
	s, err := i.getOrInitLocalShard(ctx, shardName)
	if err != nil {
		return nil, fmt.Errorf("shard %q not found locally", shardName)
	}
	return s.HashTreeLevel(ctx, level, discriminant)
}

func (i *Index) IncomingHashTreeLevel(ctx context.Context,
	shardName string, level int, discriminant *hashtree.Bitset,
) (digests []hashtree.Digest, err error) {
	return i.HashTreeLevel(ctx, shardName, level, discriminant)
}

func (i *Index) FetchObject(ctx context.Context,
	shardName string, id strfmt.UUID,
) (objects.Replica, error) {
	shard, err := i.getOrInitLocalShard(ctx, shardName)
	if err != nil {
		return objects.Replica{}, fmt.Errorf("shard %q does not exist locally", shardName)
	}

	obj, err := shard.ObjectByID(ctx, id, nil, additional.Properties{})
	if err != nil {
		return objects.Replica{}, fmt.Errorf("shard %q read repair get object: %w", shard.ID(), err)
	}

	if obj == nil {
		deleted, err := shard.WasDeleted(ctx, id)
		if err != nil {
			return objects.Replica{}, err
		}
		return objects.Replica{
			ID:      id,
			Deleted: deleted,
		}, nil
	}

	return objects.Replica{
		Object: obj,
		ID:     obj.ID(),
	}, nil
}

func (i *Index) FetchObjects(ctx context.Context,
	shardName string, ids []strfmt.UUID,
) ([]objects.Replica, error) {
	shard, err := i.getOrInitLocalShard(ctx, shardName)
	if err != nil {
		return nil, fmt.Errorf("shard %q does not exist locally", shardName)
	}

	objs, err := shard.MultiObjectByID(ctx, wrapIDsInMulti(ids))
	if err != nil {
		return nil, fmt.Errorf("shard %q replication multi get objects: %w", shard.ID(), err)
	}

	resp := make([]objects.Replica, len(ids))

	for j, obj := range objs {
		if obj == nil {
			deleted, err := shard.WasDeleted(ctx, ids[j])
			if err != nil {
				return nil, err
			}
			resp[j] = objects.Replica{
				ID:      ids[j],
				Deleted: deleted,
			}
		} else {
			resp[j] = objects.Replica{
				Object: obj,
				ID:     obj.ID(),
			}
		}
	}

	return resp, nil
}
