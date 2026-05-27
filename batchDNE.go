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
	"time"
)

// BatchDNE is a batch file that handles SEC code Death Notification Entry (DNE)
// United States Federal agencies (e.g. Social Security) use this to notify depository
// financial institutions that the recipient of government benefit payments has died.
//
// Notes:
//   - Date of death always in positions 18-23
//   - SSN (positions 38-46) are zero if no SSN
//   - Beneficiary payment starts at position 55
type BatchDNE struct {
	Batch
}

// NewBatchDNE returns a *BatchDNE
func NewBatchDNE(bh *BatchHeader) *BatchDNE { _ = "STUB: not implemented"; return nil }

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchDNE) Validate() error { _ = "STUB: not implemented"; return nil }

// SEC code

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchDNE) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// Range over Entries

// DNE must have one Addenda05

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchDNE) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

type DNEPaymentInformation struct {
	DateOfDeath time.Time
	CustomerSSN string

	// Amount is a two-decimal float value formatted as a string
	// Example: 123.45
	Amount string
}

// ParseDNEPaymentInformation returns an DNEPaymentInformation for a given Addenda05 record. The information is parsed from the addenda's
// PaymentRelatedInformation field.
//
// The returned information is not validated for correctness.
func ParseDNEPaymentInformation(addenda05 *Addenda05) (*DNEPaymentInformation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (info DNEPaymentInformation) String() string { _ = "STUB: not implemented"; return "" }
