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

package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/raft"
	command "github.com/weaviate/weaviate/cloud/proto/cluster"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/usecases/sharding"
	"google.golang.org/protobuf/proto"
)

type DB interface {
	AddClass(pl command.AddClassRequest) error
	UpdateClass(req command.UpdateClassRequest) error
	DeleteClass(string) error
	AddProperty(string, command.AddPropertyRequest) error
	AddTenants(class string, req *command.AddTenantsRequest) error
	UpdateTenants(class string, req *command.UpdateTenantsRequest) error
	DeleteTenants(class string, req *command.DeleteTenantsRequest) error
}

type Parser interface {
	ParseClass(class *models.Class) error
}

type Config struct {
	WorkDir  string // raft working directory
	NodeID   string
	Host     string
	RaftPort string
	DB       DB
	Parser   Parser
}

type Store struct {
	raft         *raft.Raft
	applyTimeout time.Duration
	raftDir      string
	nodeID       string
	host         string
	raftPort     string
	schema       *schema
	db           DB
	parser       Parser
}

func New(cfg Config) Store {
	return Store{
		raftDir:      cfg.WorkDir,
		applyTimeout: time.Second * 20,
		nodeID:       cfg.NodeID,
		host:         cfg.Host,
		raftPort:     cfg.RaftPort,
		schema:       NewSchema(cfg.NodeID),
		db:           cfg.DB,
		parser:       cfg.Parser,
	}
}

func (f *Store) SetDB(db DB) {
	f.db = db
}

func (f *Store) AddClass(cls *models.Class, ss *sharding.State) error {
	req := command.AddClassRequest{Class: cls, State: ss}
	subCommand, err := json.Marshal(&req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	cmd := &command.Command{
		Type:       command.Command_TYPE_ADD_CLASS,
		Class:      cls.Class,
		SubCommand: subCommand,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) UpdateClass(cls *models.Class, ss *sharding.State) error {
	req := command.UpdateClassRequest{Class: cls, State: ss}
	subCommand, err := json.Marshal(&req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	cmd := &command.Command{
		Type:       command.Command_TYPE_UPDATE_CLASS,
		Class:      cls.Class,
		SubCommand: subCommand,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) DeleteClass(name string) error {
	cmd := &command.Command{
		Type:  command.Command_TYPE_DELETE_CLASS,
		Class: name,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) AddProperty(class string, p *models.Property) error {
	req := command.AddPropertyRequest{Property: p}
	subCommand, err := json.Marshal(&req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	cmd := &command.Command{
		Type:       command.Command_TYPE_ADD_PROPERTY,
		Class:      class,
		SubCommand: subCommand,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) AddTenants(class string, req *command.AddTenantsRequest) error {
	subCommand, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	cmd := &command.Command{
		Type:       command.Command_TYPE_ADD_TENANT,
		Class:      class,
		SubCommand: subCommand,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) UpdateTenants(class string, req *command.UpdateTenantsRequest) error {
	subCommand, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	cmd := &command.Command{
		Type:       command.Command_TYPE_UPDATE_TENANT,
		Class:      class,
		SubCommand: subCommand,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) DeleteTenants(class string, req *command.DeleteTenantsRequest) error {
	subCommand, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	cmd := &command.Command{
		Type:       command.Command_TYPE_DELETE_TENANT,
		Class:      class,
		SubCommand: subCommand,
	}
	cmdBytes, err := proto.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("marshal command: %w", err)
	}

	fut := f.raft.Apply(cmdBytes, f.applyTimeout)
	if err := fut.Error(); err != nil {
		if errors.Is(err, raft.ErrNotLeader) {
			return ErrNotLeader
		}
		return err
	}
	return nil
}

func (f *Store) SchemaReader() *schema {
	return f.schema
}

type Response struct {
	Error error
	Data  interface{}
}

var _ raft.FSM = &Store{}

func removeNilTenants(tenants []*command.Tenant) []*command.Tenant {
	n := 0
	for i := range tenants {
		if tenants[i] != nil {
			tenants[n] = tenants[i]
			n++
		}
	}
	return tenants[:n]
}
