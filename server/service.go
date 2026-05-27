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
	"errors"
	"io"
	"time"

	"github.com/moov-io/ach"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

// Service is a REST interface for interacting with ACH file structures
// TODO: Add ctx to function parameters to pass the client security token
type Service interface {
	// CreateFile creates a new ach file record and returns a resource ID
	CreateFile(f *ach.FileHeader) (string, error)
	// AddFile retrieves a file based on the File id
	GetFile(id string) (*ach.File, error)
	// GetFiles retrieves all files accessible from the client.
	GetFiles() []*ach.File
	// BuildFile tabulates file values according to the Nacha spec
	BuildFile(id string) (*ach.File, error)
	// DeleteFile takes a file resource ID and deletes it from the store
	DeleteFile(id string) error
	// GetFileContents creates a valid plaintext file in memory assuming it has a FileHeader and at least one Batch record.
	GetFileContents(id string, opts *ach.WriteOpts) (io.Reader, error)
	// ValidateFile
	ValidateFile(id string, opts *ach.ValidateOpts) error
	// BalanceFile will apply a given offset record to the file
	BalanceFile(fileID string, off *ach.Offset) (*ach.File, error)
	// SegmentFileID segments an ach file
	SegmentFileID(id string, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error)
	// SegmentFile segments an ach file
	SegmentFile(file *ach.File, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error)
	// FlattenBatches will minimize the ach.Batch objects in a file by consolidating EntryDetails under distinct batch headers
	FlattenBatches(id string) (*ach.File, error)
	// CreateBatch creates a new batch within and ach file and returns its resource ID
	CreateBatch(fileID string, bh ach.Batcher) (string, error)
	// GetBatch retrieves a batch based oin the file id and batch id
	GetBatch(fileID string, batchID string) (ach.Batcher, error)
	// GetBatches retrieves all batches associated with the file id.
	GetBatches(fileID string) []ach.Batcher
	// DeleteBatch takes a fileID and BatchID and removes the batch from the file
	DeleteBatch(fileID string, batchID string) error
	// MergeFiles will combine all the given files together
	MergeFiles(fileIDs []string, files []*ach.File, conditions *ach.Conditions) ([]*ach.File, error)
	// ReverseFile creates a NACHA compliant reversal of the ACH file
	ReverseFile(fileID string, effectiveEntryDate time.Time) (*ach.File, error)
}

// service a concrete implementation of the service.
type service struct {
	store Repository
}

// NewService creates a new concrete service
func NewService(r Repository) Service { _ = "STUB: not implemented"; return *new(Service) }

// CreateFile add a file to storage
// TODO(adam): the HTTP endpoint accepts malformed bodies (and missing data)
func (s *service) CreateFile(fh *ach.FileHeader) (string, error) {
	_ = "STUB: not implemented"
	// create a new file
	return "", nil
}

// set resource id's

// GetFile returns a files based on the supplied id
func (s *service) GetFile(id string) (*ach.File, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *service) GetFiles() []*ach.File { _ = "STUB: not implemented"; return nil }

// BuildFile tabulates file values according to the Nacha spec
func (s *service) BuildFile(id string) (*ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone the file to avoid mutating the original in the repository

func (s *service) DeleteFile(id string) error { _ = "STUB: not implemented"; return nil }

func (s *service) GetFileContents(id string, opts *ach.WriteOpts) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// Clone the file to avoid mutating the original in the repository

func (s *service) ValidateFile(id string, opts *ach.ValidateOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *service) CreateBatch(fileID string, batch ach.Batcher) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *service) GetBatch(fileID string, batchID string) (ach.Batcher, error) {
	_ = "STUB: not implemented"
	return *new(ach.Batcher), nil
}

func (s *service) GetBatches(fileID string) []ach.Batcher { _ = "STUB: not implemented"; return nil }

func (s *service) DeleteBatch(fileID string, batchID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *service) BalanceFile(fileID string, off *ach.Offset) (*ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone the file to avoid mutating the original in the repository

// Apply the Offset to each Batch and then re-create (to tabulate new EntryDetail records)

// overwrite the ID so it's new and unique

// Save our new file

// SegmentFileID takes an ACH FileID and segments the files into a credit ACH File and debit ACH File and adds to in memory storage.
func (s *service) SegmentFileID(fileID string, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Clone the file to avoid mutating the original in the repository

// SegmentFile takes an ACH File and segments the files into a credit ACH File and debit ACH File and adds to in memory storage.
func (s *service) SegmentFile(file *ach.File, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error) {
	_ = "STUB: not implemented"
	// Build/tabulate file in the case it is malformed.
	return nil, nil, nil
}

// FlattenBatches consolidates batches that have the same BatchHeader
func (s *service) FlattenBatches(fileID string) (*ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone the file to avoid mutating the original in the repository

// File Create in the case a file is malformed.

func (s *service) MergeFiles(fileIDs []string, files []*ach.File, conditions *ach.Conditions) ([]*ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hash(data []byte) string { _ = "STUB: not implemented"; return "" }

// ReverseFile creates a NACHA compliant reversal of the ACH file
func (s *service) ReverseFile(fileID string, effectiveEntryDate time.Time) (*ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone the file to avoid modifying the original in the repository

// new ID for reversed file

// cloneFile creates a deep copy of the file via JSON serialization.
// This prevents mutations to the returned file from affecting the original in the repository.
// JSON is used instead of ACH Writer/Reader because it can handle files that haven't been built yet.
func cloneFile(f *ach.File) (*ach.File, error) { _ = "STUB: not implemented"; return nil, nil }

// Use SkipAll to avoid validation during cloning - we just want an exact copy

// Restore original validation options (or clear the SkipAll we used for cloning)
