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

// Addenda10 is an addenda which provides business transaction information for Addenda Type
// Code 10 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda10 is mandatory for IAT entries
//
// The Addenda10 Record identifies the Receiver of the transaction and the dollar amount of
// the payment.
type Addenda10 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda10 types code '10'
	TypeCode string `json:"typeCode"`
	// Transaction Type Code Describes the type of payment:
	// ANN = Annuity, BUS = Business/Commercial, DEP = Deposit, LOA = Loan, MIS = Miscellaneous, MOR = Mortgage
	// PEN = Pension, RLS = Rent/Lease, REM = Remittance2, SAL = Salary/Payroll, TAX = Tax, TEL = Telephone-Initiated Transaction
	// WEB = Internet-Initiated Transaction, ARC = Accounts Receivable Entry, BOC = Back Office Conversion Entry,
	// POP = Point of Purchase Entry, RCK = Re-presented Check Entry
	TransactionTypeCode string `json:"transactionTypeCode"`
	// Foreign Payment Amount $$$$$$$$$$$$$$$$¢¢
	// For inbound IAT payments this field should contain the USD amount or may be blank.
	ForeignPaymentAmount int `json:"foreignPaymentAmount"`
	// Foreign Trace Number
	ForeignTraceNumber string `json:"foreignTraceNumber,omitempty"`
	// Receiving Company Name/Individual Name
	Name string `json:"name"`
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace number This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int `json:"entryDetailSequenceNumber"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
	// validateOpts defines optional overrides for record validation
	validateOpts *ValidateOpts
}

// NewAddenda10 returns a new Addenda10 with default values for none exported fields
func NewAddenda10() *Addenda10 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda10 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda10 *Addenda10) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 10

// 04-06 Describes the type of payment

// 07-24 Payment Amount	For inbound IAT payments this field should contain the USD amount or may be blank.

//  25-46 Insert blanks or zeros

// 47-81 Receiving Company Name/Individual Name

// 82-87 reserved - Leave blank

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda10) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda10 struct to a 94 character string.
func (addenda10 *Addenda10) String() string { _ = "STUB: not implemented"; return "" }

// TransactionTypeCode Validator

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda10 *Addenda10) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 10

// ToDo: Foreign Payment Amount blank ?

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda10 *Addenda10) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// ToDo:  Commented because it appears this value can be all 000 (maybe blank?)
/*	if addenda10.ForeignPaymentAmount == 0 {
	return fieldError( "ForeignPaymentAmount", ErrFieldRequired,  strconv.Itoa(addenda10.ForeignPaymentAmount))
}*/

// ForeignPaymentAmountField returns ForeignPaymentAmount zero padded
// ToDo: Review/Add logic for blank ?
func (addenda10 *Addenda10) ForeignPaymentAmountField() string {
	_ = "STUB: not implemented"
	return ""
}

// ForeignTraceNumberField gets the Foreign TraceNumber left padded
func (addenda10 *Addenda10) ForeignTraceNumberField() string { _ = "STUB: not implemented"; return "" }

// NameField gets the name field - Receiving Company Name/Individual Name left padded
func (addenda10 *Addenda10) NameField() string { _ = "STUB: not implemented"; return "" }

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda10 *Addenda10) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
