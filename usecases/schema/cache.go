//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/usecases/sharding"
)

var (
	errClassNotFound = errors.New("class not found")
	errShardNotFound = errors.New("shard not found")
)

// State is a cached copy of the schema that can also be saved into a remote
// storage, as specified by Repo
type State struct {
	ObjectSchema  *models.Schema `json:"object"`
	ShardingState map[string]*sharding.State
}

// NewState returns a new state with room for nClasses classes
func NewState(nClasses int) State {
	return State{
		ObjectSchema: &models.Schema{
			Classes: make([]*models.Class, 0, nClasses),
		},
		ShardingState: make(map[string]*sharding.State, nClasses),
	}
}

type schemaCache struct {
	sync.RWMutex
	State
}

// ShardOwner returns the node owner of the specified shard
func (s *schemaCache) ShardOwner(class, shard string) (string, error) {
	s.RLock()
	defer s.RUnlock()
	cls := s.ShardingState[class]
	if cls == nil {
		return "", errClassNotFound
	}
	x, ok := cls.Physical[shard]
	if !ok {
		return "", errShardNotFound
	}
	if len(x.BelongsToNodes) < 1 || x.BelongsToNodes[0] == "" {
		return "", fmt.Errorf("owner node not found")
	}
	return x.BelongsToNodes[0], nil
}

// ShardOwner returns the node owner of the specified shard
func (s *schemaCache) ShardReplicas(class, shard string) ([]string, error) {
	s.RLock()
	defer s.RUnlock()
	cls := s.ShardingState[class]
	if cls == nil {
		return nil, errClassNotFound
	}
	x, ok := cls.Physical[shard]
	if !ok {
		return nil, errShardNotFound
	}
	return x.BelongsToNodes, nil
}

// TenantShard returns shard name for the provided tenant
func (s *schemaCache) TenantShard(class, tenant string) string {
	s.RLock()
	defer s.RUnlock()
	ss := s.ShardingState[class]
	if ss == nil {
		return ""
	}
	return ss.Shard(tenant, "")
}

// ShardFromUUID returns shard name of the provided uuid
func (s *schemaCache) ShardFromUUID(class string, uuid []byte) string {
	s.RLock()
	defer s.RUnlock()
	ss := s.ShardingState[class]
	if ss == nil {
		return ""
	}
	return ss.PhysicalShard(uuid)
}

func (s *schemaCache) CopyShardingState(className string) *sharding.State {
	s.RLock()
	defer s.RUnlock()
	pst := s.ShardingState[className]
	if pst != nil {
		st := pst.DeepCopy()
		pst = &st
	}

	return pst
}

// LockGuard provides convenient mechanism for owning mutex by function which mutates the state
func (s *schemaCache) LockGuard(mutate func()) {
	s.Lock()
	defer s.Unlock()
	mutate()
}

// RLockGuard provides convenient mechanism for owning mutex function which doesn't mutates the state
func (s *schemaCache) RLockGuard(reader func() error) error {
	s.RLock()
	defer s.RUnlock()
	return reader()
}

func (s *schemaCache) isEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.State.ObjectSchema == nil || len(s.State.ObjectSchema.Classes) == 0
}

func (s *schemaCache) setState(st State) {
	s.Lock()
	defer s.Unlock()
	s.State = st
}

func (s *schemaCache) detachClass(name string) bool {
	s.Lock()
	defer s.Unlock()
	schema, ci := s.ObjectSchema, -1
	for i, cls := range schema.Classes {
		if cls.Class == name {
			ci = i
			break
		}
	}
	if ci == -1 {
		return false
	}

	// update all at once to prevent race condition with concurrent readers
	xs := make([]*models.Class, len(schema.Classes)-1)
	copy(xs, schema.Classes[:ci])
	copy(xs, schema.Classes[ci+1:])
	schema.Classes = xs
	return true
}

func (s *schemaCache) deleteClassState(name string) {
	s.Lock()
	defer s.Unlock()
	delete(s.ShardingState, name)
}

func (s *schemaCache) unsafeFindClassIf(pred func(*models.Class) bool) *models.Class {
	for _, cls := range s.ObjectSchema.Classes {
		if pred(cls) {
			return cls
		}
	}
	return nil
}

func (s *schemaCache) unsafeFindClass(className string) *models.Class {
	for _, c := range s.ObjectSchema.Classes {
		if c.Class == className {
			return c
		}
	}
	return nil
}

func (s *schemaCache) AddClass(c *models.Class, ss *sharding.State) {
	s.Lock()
	defer s.Unlock()

	// update all at once to prevent race condition with concurrent readers
	cs := make([]*models.Class, len(s.ObjectSchema.Classes)+1)
	copy(cs, s.ObjectSchema.Classes)
	cs[len(s.ObjectSchema.Classes)] = c
	s.ObjectSchema.Classes = cs

	s.ShardingState[c.Class] = ss
}

func (s *schemaCache) AddProperty(class string, p *models.Property) ([]byte, error) {
	s.Lock()
	defer s.Unlock()

	c := s.unsafeFindClass(class)
	if c == nil {
		return nil, errClassNotFound
	}

	// update all at once to prevent race condition with concurrent readers
	src := c.Properties
	dest := make([]*models.Property, len(src)+1)
	copy(dest, src)
	dest[len(src)] = p
	c.Properties = dest
	metadata, err := json.Marshal(&class)
	if err != nil {
		c.Properties = src
		return nil, fmt.Errorf("marshal class %s: %w", class, err)
	}
	return metadata, nil
}
