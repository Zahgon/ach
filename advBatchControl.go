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

// ADVBatchControl contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type ADVBatchControl struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// This should be the same as BatchHeader ServiceClassCode for ADV: AutomatedAccountingAdvices.
	ServiceClassCode int `json:"serviceClassCode"`
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda
	// Record processed, within either the batch or file as appropriate.
	EntryAddendaCount int `json:"entryAddendaCount"`
	// validate the Receiving DFI Identification in each Entry Detail Record is hashed
	// to provide a check against inadvertent alteration of data contents due
	// to hardware failure or program error
	//
	// In this context the Entry Hash is the sum of the corresponding fields in the
	// Entry Detail Records on the file.
	EntryHash int `json:"entryHash"`
	// TotalDebitEntryDollarAmount Contains accumulated Entry debit totals within the batch.
	TotalDebitEntryDollarAmount int `json:"totalDebit"`
	// TotalCreditEntryDollarAmount Contains accumulated Entry credit totals within the batch.
	TotalCreditEntryDollarAmount int `json:"totalCredit"`
	// ACHOperatorData is an alphanumeric code used to identify an ACH Operator
	ACHOperatorData string `json:"achOperatorData"`
	// ODFIIdentification the routing number is used to identify the DFI originating entries within a given branch.
	ODFIIdentification string `json:"ODFIIdentification"`
	// BatchNumber this number is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by record.
	BatchNumber int `json:"batchNumber"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
	// validateOpts defines optional overrides for record validation
	validateOpts *ValidateOpts
}

// Parse takes the input record string and parses the EntryDetail values
func (bc *ADVBatchControl) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always "8"

// 2-4 This is the same as the "Service code" field in previous Batch Header Record

// 5-10 Total number of Entry Detail Record in the batch

// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10

// 21-32 Number of cents of debit entries within the batch

// 33-44 Number of cents of credit entries within the batch

// 45-54 ACH Operator Data

// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record

// 88-94 This is the same as the "Batch number" field in previous Batch Header Record

func (a *ADVBatchControl) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// NewADVBatchControl returns a new ADVBatchControl with default values for none exported fields
func NewADVBatchControl() *ADVBatchControl { _ = "STUB: not implemented"; return nil }

// String writes the ADVBatchControl struct to a 94 character string.
func (bc *ADVBatchControl) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *ADVBatchControl) Validate() error { _ = "STUB: not implemented"; return nil }

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bc *ADVBatchControl) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// EntryAddendaCountField gets a string of the addenda count zero padded
func (bc *ADVBatchControl) EntryAddendaCountField() string { _ = "STUB: not implemented"; return "" }

// EntryHashField get a zero padded EntryHash
func (bc *ADVBatchControl) EntryHashField() string { _ = "STUB: not implemented"; return "" }

// TotalDebitEntryDollarAmountField get a zero padded Debit Entry Amount
func (bc *ADVBatchControl) TotalDebitEntryDollarAmountField() string {
	_ = "STUB: not implemented"
	return ""
}

// TotalCreditEntryDollarAmountField get a zero padded Credit Entry Amount
func (bc *ADVBatchControl) TotalCreditEntryDollarAmountField() string {
	_ = "STUB: not implemented"
	return ""
}

// ACHOperatorDataField get the ACHOperatorData right padded
func (bc *ADVBatchControl) ACHOperatorDataField() string { _ = "STUB: not implemented"; return "" }

// ODFIIdentificationField get the odfi number zero padded
func (bc *ADVBatchControl) ODFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// BatchNumberField gets a string of the batch number zero padded
func (bc *ADVBatchControl) BatchNumberField() string { _ = "STUB: not implemented"; return "" }
