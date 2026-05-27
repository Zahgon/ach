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

// BatchTEL is a batch that handles SEC payment type Telephone-Initiated Entries (TEL)
// Telephone-Initiated Entries (TEL) are consumer debit transactions. The NACHA Operating Rules permit TEL entries when
// the Originator obtains the Receiver's authorization for the debit entry orally via the telephone.
// An entry based upon a Receiver's oral authorization must utilize the TEL (Telephone-Initiated Entry)
// Standard Entry Class (SEC) Code.
type BatchTEL struct {
	Batch
}

// NewBatchTEL returns a *BatchTEL
func NewBatchTEL(bh *BatchHeader) *BatchTEL { _ = "STUB: not implemented"; return nil }

// Validate ensures the batch meets NACHA rules specific to the SEC type TEL
func (batch *BatchTEL) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration and type specific based validation for this type.

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchTEL) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// Forward TEL batches can only have debit entries, but REVERSAL batches can only have credits

// Forward batch: can only have debits

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchTEL) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}
