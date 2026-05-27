// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package server

import (
	"sync"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/base/log"
)

// Repository is the Service storage mechanism abstraction
type Repository interface {
	StoreFile(file *ach.File) error
	FindFile(id string) (*ach.File, error)
	FindAllFiles() []*ach.File
	DeleteFile(id string) error
	StoreBatch(fileID string, batch ach.Batcher) error
	FindBatch(fileID string, batchID string) (ach.Batcher, error)
	FindAllBatches(fileID string) []ach.Batcher
	DeleteBatch(fileID string, batchID string) error
}

type repositoryInMemory struct {
	mtx   sync.RWMutex
	files map[string]*ach.File

	ttl time.Duration

	logger log.Logger
}

// NewRepositoryInMemory is an in memory ach storage repository for files
func NewRepositoryInMemory(ttl time.Duration, logger log.Logger) Repository {
	_ = "STUB: not implemented"
	return *new(Repository)
}

// Don't run the cleanup if we've disabled the TTL

// Run our anon goroutine to cleanup old ACH files

func (r *repositoryInMemory) StoreFile(f *ach.File) error { _ = "STUB: not implemented"; return nil }

// FindFile retrieves a ach.File based on the supplied ID
func (r *repositoryInMemory) FindFile(id string) (*ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindAllFiles returns all files that have been saved in memory
func (r *repositoryInMemory) FindAllFiles() []*ach.File { _ = "STUB: not implemented"; return nil }

func (r *repositoryInMemory) DeleteFile(id string) error { _ = "STUB: not implemented"; return nil }

// TODO(adam): was copying ach.Batcher causing issues?
func (r *repositoryInMemory) StoreBatch(fileID string, batch ach.Batcher) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure the file does not already exist

// ensure the batch does not already exist

// Add the batch to the file

// FindBatch retrieves a ach.Batcher based on the supplied ID
func (r *repositoryInMemory) FindBatch(fileID string, batchID string) (ach.Batcher, error) {
	_ = "STUB: not implemented"
	return *new(ach.Batcher), nil
}

// FindAllBatches
func (r *repositoryInMemory) FindAllBatches(fileID string) []ach.Batcher {
	_ = "STUB: not implemented"
	return nil
}

func (r *repositoryInMemory) DeleteBatch(fileID string, batchID string) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanupOldFiles will iterate through r.files and delete entries which are older than
// the environmental variable ACH_FILE_TTL (parsed as a time.Duration).
func (r *repositoryInMemory) cleanupOldFiles() { _ = "STUB: not implemented"; return }

// YYMMDD
