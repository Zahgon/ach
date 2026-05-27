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

// BatchTRX holds the BatchHeader and BatchControl and all EntryDetail for TRX Entries.
//
// Check Truncation Entries Exchange is used to identify a debit entry of a truncated checks (multiple).
type BatchTRX struct {
	Batch
}

// NewBatchTRX returns a *BatchTRX
func NewBatchTRX(bh *BatchHeader) *BatchTRX { _ = "STUB: not implemented"; return nil }

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchTRX) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration and type specific validation for this type.

// TRX detail entries can only be a debit, ServiceClassCode must allow debits

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchTRX) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// TRX detail entries must be a debit

// Trapping this error, as entry.CTXAddendaRecordsField() can not be greater than 9999

// validate CTXAddendaRecord Field is equal to the actual number of Addenda records
// use 0 value if there is no Addenda records

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchTRX) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// Additional steps specific to batch type
// ...
