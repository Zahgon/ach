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

// BatchXCK holds the BatchHeader and BatchControl and all EntryDetail for XCK Entries.
//
// Destroyed Check Entry identifies a debit entry initiated for a XCK eligible items.
type BatchXCK struct {
	Batch
}

// NewBatchXCK returns a *BatchXCK
func NewBatchXCK(bh *BatchHeader) *BatchXCK { _ = "STUB: not implemented"; return nil }

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchXCK) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration and type specific validation for this type.

// XCK detail entries can only be a debit, ServiceClassCode must allow debits

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchXCK) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// XCK detail entries must be a debit

// Amount must be 2,500 or less

// ProcessControlField underlying IdentificationNumber, must be defined

// ItemResearchNumber underlying IdentificationNumber, must be defined

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchXCK) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// Additional steps specific to batch type
// ...
