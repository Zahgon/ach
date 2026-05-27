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

// BatchCOR COR - Automated Notification of Change (NOC) or Refused Notification of Change
// This Standard Entry Class Code is used by an RDFI or ODFI when originating a Notification of Change or Refused Notification of Change in automated format.
// A Notification of Change may be created by an RDFI to notify the ODFI that a posted Entry or Prenotification Entry contains invalid or erroneous information and should be changed.
type BatchCOR struct {
	Batch
}

// NewBatchCOR returns a *BatchCOR
func NewBatchCOR(bh *BatchHeader) *BatchCOR { _ = "STUB: not implemented"; return nil }

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchCOR) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration based validation for this type.
// COR Addenda must be Addenda98

// Add type specific validation.

// The Amount field must be zero
// batch.verify calls batch.isBatchAmount which ensures the batch.Control values are accurate.

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchCOR) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

/* COR TransactionCode must be a Return or NOC transaction Code
   Return/NOC
   Credit:  21, 31, 41, 51
   Debit: 26, 36, 46, 56
*/

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchCOR) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// isAddenda98 verifies that a Addenda98 exists for each EntryDetail and is Validated
func (batch *BatchCOR) isAddenda98() error { _ = "STUB: not implemented"; return nil }
