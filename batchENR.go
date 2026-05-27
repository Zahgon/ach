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

// BatchENR is a non-monetary entry that enrolls a person with an agency of the US government
// for a depository financial institution.
//
// Allowed TransactionCode values: 22 Demand Credit, 27 Demand Debit, 32 Savings Credit, 37 Savings Debit
type BatchENR struct {
	Batch
}

// NewBatchENR returns a *BatchENR
func NewBatchENR(bh *BatchHeader) *BatchENR { _ = "STUB: not implemented"; return nil }

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchENR) Validate() error { _ = "STUB: not implemented"; return nil }

// Batch Header checks

// return the first invalid entry's error

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchENR) InvalidEntries() []InvalidEntry { _ = "STUB: not implemented"; return nil }

// Range over Entries

// nothing

// reurns, do nothing

// Verify the Amount is valid for SEC code and TransactionCode

// Verify the TransactionCode is valid for a ServiceClassCode

// ENR must have one Addenda05
// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchENR) Create() error {
	_ = "STUB: not implemented"
	// generates sequence numbers and batch control
	return nil
}

// ENRPaymentInformation structure
type ENRPaymentInformation struct {
	// TransactionCode is the Transaction Code of the holder's account
	// Values: 22 (Demand  Credit), 27 (Demand Debit), 32 (Savings Credit), 37 (Savings Debit)
	TransactionCode int

	// RDFIIdentification is the Receiving Depository Identification Number. Typically the first 8 of their ABA routing number.
	RDFIIdentification string

	// CheckDigit is the last digit from an ABA routing number.
	CheckDigit string

	// DFIAccountNumber contains the holder's account number.
	DFIAccountNumber string

	// IndividualIdentification contains the customer's Social Security Number (SSN) for automated enrollments and the
	// taxpayer ID for companies.
	IndividualIdentification string

	// IndividualName is the account holders full name.
	IndividualName string

	// EnrolleeClassificationCode (also called Representative Payee Indicator) returns a code from a specific Addenda05 record.
	// These codes represent:
	//  0: (no)  - Initiated by beneficiary
	//  1: (yes) - Initiated by someone other than named beneficiary
	//  A: Enrollee is a consumer
	//  b: Enrollee is a company
	EnrolleeClassificationCode string
}

func (info ENRPaymentInformation) String() string {
	_ = "STUB: not implemented"
	// Stretch the companies name across two fields
	return ""
}

// First fifteen characters are added

// Add on a second field if needed

// Format the Individual's name by Surname first

// Surname comes fist

// ParseENRPaymentInformation returns an ENRPaymentInformation for a given Addenda05 record. The information is parsed from the addenda's
// PaymentRelatedInformation field.
//
// The returned information is not validated for correctness.
func ParseENRPaymentInformation(addenda05 *Addenda05) (*ENRPaymentInformation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PaymentRelatedInformation is terminated by '\'

// Business Names can be fill two field lengths
