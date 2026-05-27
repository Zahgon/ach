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

// BatchPOP holds the BatchHeader and BatchControl and all EntryDetail for POP Entries.
//
// Point-of-Purchase. A check presented in-person to a merchant for purchase is presented
// as an ACH entry instead of a physical check.
//
// This ACH debit application is used by originators as a method of payment for the
// in-person purchase of goods or services by consumers. These Single Entry debit
// entries are initiated by the originator based on a written authorization and
// account information drawn from the source document (a check) obtained from the
// consumer at the point-of-purchase. The source document, which is voided by the
// merchant and returned to the consumer at the point-of-purchase, is used to
// collect the consumer's routing number, account number and check serial number that
// will be used to generate the debit entry to the consumer's account.
//
// The difference between POP and ARC is that ARC can result from a check mailed in whereas POP is in-person.
type BatchPOP struct {
	Batch
}

// NewBatchPOP returns a *BatchPOP
func NewBatchPOP(bh *BatchHeader) *BatchPOP { _ = "STUB: not implemented"; return nil }

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchPOP) Validate() error { _ = "STUB: not implemented"; return nil }

// basic verification of the batch before we validate specific rules.

// Add configuration and type specific validation for this type.

// POP detail entries can only be a debit, ServiceClassCode must allow debits

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchPOP) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// POP detail entries must be a debit

// Amount must be 25,000 or less

// CheckSerialNumber, Terminal City, Terminal State underlying IdentificationNumber, must be defined

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchPOP) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// Additional steps specific to batch type
// ...
