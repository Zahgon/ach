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

package ach

import (
	"errors"
)

var (
	ErrFlattenChangedEntryCount   = errors.New("Flatten operation changed entry and addenda count")
	ErrFlattenChangedDebitAmount  = errors.New("Flatten operation changed total debit entry amount")
	ErrFlattenChangedCreditAmount = errors.New("Flatten operation changed total credit entry amount")
)

// Flatten returns a flattened version of a File, where batches with similar batch headers are consolidated.
//
// Two batches are eligible to be combined if:
//   - their headers match, excluding the batch number (which isn't used in return matching and reflects
//     the final composition of the file.)
//   - they don't contain any entries with common trace numbers, since trace numbers must be unique
//     within a batch.
func Flatten(originalFile *File) (*File, error) { _ = "STUB: not implemented"; return nil, nil }

// Convert batches and IAT batches to "mergeables" for consistent flattening logic

// Considering bigger batches first allows for the least number of flattened batches

// Merge each original batch into a new batch

// Create a new file containing each of our new batches

// Sort batches by original batch number to roughly maintain batch order in the flattened file

// Sanity checks; this is kind of a scary operation!

// FlattenBatches flattens the file's batches by consolidating batches with the same BatchHeader data into one Batch.
// Entries within each flattened batch will be sorted by their TraceNumber field.
func (f *File) FlattenBatches() (*File, error) {
	_ = "STUB: not implemented"

	// Determine if two batches can be combined (ie, have the same header and no common trace numbers)
	return nil, nil
}

func canMerge(a mergeable, b mergeable) bool { _ = "STUB: not implemented"; return false }

// Represents either a "normal" batch or an IAT batch
type mergeable interface {
	GetHeaderSignature() string
	GetTraceNumbers() map[string]bool
	Consume(mergeable) error
	GetBatch() interface{}
	GetBatchNumber() int
	Copy() mergeable
	GetEntryCount() int
	AddToFile(*File) error
}

type mergeableBatcher struct {
	batcher      Batcher
	traceNumbers map[string]bool
}

// Batch header excluding the batch number, which isn't important to preserve
func (b mergeableBatcher) GetHeaderSignature() string { _ = "STUB: not implemented"; return "" }
func (b mergeableBatcher) GetBatch() interface{}      { _ = "STUB: not implemented"; return nil }
func (b mergeableBatcher) GetEntryCount() int         { _ = "STUB: not implemented"; return 0 }
func (b mergeableBatcher) GetBatchNumber() int        { _ = "STUB: not implemented"; return 0 }

func (b mergeableBatcher) GetTraceNumbers() map[string]bool { _ = "STUB: not implemented"; return nil }

func (m mergeableBatcher) Consume(mergeableToConsume mergeable) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep the lower of the two batch numbers, to roughly maintain batch order in the flattened file

func (m mergeableBatcher) Copy() mergeable { _ = "STUB: not implemented"; return *new(mergeable) }

func (m mergeableBatcher) AddToFile(file *File) error {
	_ = "STUB: not implemented"
	// Sort entries by trace number
	return nil
}

// Inherit validation options from file if batch doesn't have them

// Also set validation on all entries

type mergeableIATBatch struct {
	iatBatch     *IATBatch
	traceNumbers map[string]bool
}

// Batch header excluding the batch number, which isn't important to preserve
func (b mergeableIATBatch) GetHeaderSignature() string { _ = "STUB: not implemented"; return "" }
func (b mergeableIATBatch) GetBatch() interface{}      { _ = "STUB: not implemented"; return nil }
func (b mergeableIATBatch) GetEntryCount() int         { _ = "STUB: not implemented"; return 0 }
func (b mergeableIATBatch) GetBatchNumber() int        { _ = "STUB: not implemented"; return 0 }

func (b mergeableIATBatch) GetTraceNumbers() map[string]bool { _ = "STUB: not implemented"; return nil }

func (m mergeableIATBatch) Consume(mergeableToConsume mergeable) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep the lower of the two batch numbers, to roughly maintain batch order in the flattened file

func (m mergeableIATBatch) Copy() mergeable { _ = "STUB: not implemented"; return *new(mergeable) }

func (m mergeableIATBatch) AddToFile(file *File) error {
	_ = "STUB: not implemented"
	// Sort entries by trace number
	return nil
}

// Inherit validation options from file

// Also set validation on all entries
