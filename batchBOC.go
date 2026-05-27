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

// BatchBOC holds the BatchHeader and BatchControl and all EntryDetail for BOC Entries.
//
// Back Office Conversion (BOC) A single entry debit initiated at the point of purchase
// or at a manned bill payment location to transfer funds through conversion to an
// ACH debit entry during back office processing.
//
// BOC allows retailers/billers, and ODFIs acting as Originators,
// to electronically convert checks received at the point-of-purchase as well as at a
// manned bill payment location into a single-entry ACH debit. The authorization to
// convert the check will be obtained through a notice at the checkout or manned bill
// payment location (e.g., loan payment at financial institution's teller window) and the
// receipt of the Receiver's check. The decision to process the check item as an ACH debit
// will be made in the “back office” instead of at the point-of-purchase. The customer's
// check will solely be used as a source document to obtain the routing number, account
// number and check serial number.
//
// Unlike ARC entries, BOC conversions require the customer to be present and a notice that
// checks may be converted to BOC ACH entries be posted.
type BatchBOC struct {
	Batch
}

// NewBatchBOC returns a *BatchBOC
func NewBatchBOC(bh *BatchHeader) *BatchBOC { _ = "STUB: not implemented"; return nil }

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchBOC) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration and type specific validation for this type.

// BOC detail entries can only be a debit, ServiceClassCode must allow debits

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchBOC) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// BOC detail entries must be a debit

// Amount must be 25,000 or less

// CheckSerialNumber underlying IdentificationNumber, must be defined

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchBOC) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// Additional steps specific to batch type
// ...
