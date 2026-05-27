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

// BatchARC holds the BatchHeader and BatchControl and all EntryDetail for ARC Entries.
//
// Accounts Receivable Entry (ARC). A consumer check converted to a one-time ACH debit.
// The Accounts Receivable (ARC) Entry provides billers the opportunity to initiate single-entry ACH
// debits to customer accounts by converting checks at the point of receipt through the U.S. mail, at
// a drop box location or in-person for payment of a bill at a manned location. The biller is required
// to provide the customer with notice prior to the acceptance of the check that states the receipt of
// the customer's check will be deemed as the authorization for an ARC debit entry to the customer's
// account. The provision of the notice and the receipt of the check together constitute authorization
// for the ARC entry. The customer's check is solely be used as a source document to obtain the routing
// number, account number and check serial number.
//
// The difference between ARC and POP is that ARC can result from a check mailed in whereas POP is in-person.
type BatchARC struct {
	Batch
}

// NewBatchARC returns a *BatchARC
func NewBatchARC(bh *BatchHeader) *BatchARC { _ = "STUB: not implemented"; return nil }

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchARC) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration and type specific validation for this type.

// ARC detail entries can only be a debit, ServiceClassCode must allow debits

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchARC) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// ARC detail entries must be a debit

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
func (batch *BatchARC) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// Additional steps specific to batch type
// ...
