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

// Addenda02 is a Addendumer addenda which provides business transaction information for Addenda Type
// Code 02 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
// It is used for following StandardEntryClassCode: MTE, POS, and SHR.
type Addenda02 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda02 type code '02'
	TypeCode string `json:"typeCode"`
	// ReferenceInformationOne may be used for additional reference numbers, identification numbers,
	// or codes that the merchant needs to identify the particular transaction or customer.
	ReferenceInformationOne string `json:"referenceInformationOne,omitempty"`
	// ReferenceInformationTwo  may be used for additional reference numbers, identification numbers,
	// or codes that the merchant needs to identify the particular transaction or customer.
	ReferenceInformationTwo string `json:"referenceInformationTwo,omitempty"`
	// TerminalIdentificationCode identifies an Electronic terminal with a unique code that allows
	// a terminal owner and/or switching network to identify the terminal at which an Entry originated.
	TerminalIdentificationCode string `json:"terminalIdentificationCode"`
	// TransactionSerialNumber is assigned by the terminal at the time the transaction is originated.  The
	// number, with the Terminal Identification Code, serves as an audit trail for the transaction and is
	// usually assigned in ascending sequence.
	TransactionSerialNumber string `json:"transactionSerialNumber"`
	// TransactionDate expressed MMDD identifies the date on which the transaction occurred.
	TransactionDate string `json:"transactionDate"`
	// AuthorizationCodeOrExpireDate indicates the code that a card authorization center has
	// furnished to the merchant.
	AuthorizationCodeOrExpireDate string `json:"authorizationCodeOrExpireDate,omitempty"`
	// Terminal Location identifies the specific location of a terminal (i.e., street names of an
	// intersection, address, etc.) in accordance with the requirements of Regulation E.
	TerminalLocation string `json:"terminalLocation"`
	// TerminalCity Identifies the city in which the electronic terminal is located.
	TerminalCity string `json:"terminalCity"`
	// TerminalState Identifies the state in which the electronic terminal is located
	TerminalState string `json:"terminalState"`
	// TraceNumber Standard Entry Detail Trace Number
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
	// validateOpts defines optional overrides for record validation
	validateOpts *ValidateOpts
}

// NewAddenda02 returns a new Addenda02 with default values for none exported fields
func NewAddenda02() *Addenda02 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda02 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda02 *Addenda02) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 02

// 4-10 Based on the information entered (04-10) 7 alphanumeric

// 11-13 Based on the information entered (11-13) 3 alphanumeric

// 14-19

// 20-25

// 26-29

// 30-35

// 36-62

// 63-77

// 78-79

// 80-94

func (a *Addenda02) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda02 struct to a 94 character string.
func (addenda02 *Addenda02) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda02 *Addenda02) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 02

// TransactionDate Addenda02 ACH File format is MMDD. Validate MM is 01-12 and day for the
// month 01-31 depending on month.

// fieldInclusion validate mandatory fields are not default values  and required fields are defined. If fields are
// invalid the ACH transfer will be returned.

func (addenda02 *Addenda02) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// Required Fields

// ReferenceInformationOneField returns a space padded ReferenceInformationOne string
func (addenda02 *Addenda02) ReferenceInformationOneField() string {
	_ = "STUB: not implemented"
	return ""
}

// ReferenceInformationTwoField returns a space padded ReferenceInformationTwo string
func (addenda02 *Addenda02) ReferenceInformationTwoField() string {
	_ = "STUB: not implemented"
	return ""
}

// TerminalIdentificationCodeField returns a space padded TerminalIdentificationCode string
func (addenda02 *Addenda02) TerminalIdentificationCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

// TransactionSerialNumberField returns a zero padded TransactionSerialNumber string
func (addenda02 *Addenda02) TransactionSerialNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

// TransactionDateField returns TransactionDate MMDD string
func (addenda02 *Addenda02) TransactionDateField() string { _ = "STUB: not implemented"; return "" }

// AuthorizationCodeOrExpireDateField returns a space padded AuthorizationCodeOrExpireDate string
func (addenda02 *Addenda02) AuthorizationCodeOrExpireDateField() string {
	_ = "STUB: not implemented"
	return ""
}

// TerminalLocationField returns a space padded TerminalLocation string
func (addenda02 *Addenda02) TerminalLocationField() string { _ = "STUB: not implemented"; return "" }

// TerminalCityField returns a space padded TerminalCity string
func (addenda02 *Addenda02) TerminalCityField() string { _ = "STUB: not implemented"; return "" }

// TerminalStateField returns a space padded TerminalState string
func (addenda02 *Addenda02) TerminalStateField() string { _ = "STUB: not implemented"; return "" }

// TraceNumberField returns a space padded TraceNumber string
func (addenda02 *Addenda02) TraceNumberField() string { _ = "STUB: not implemented"; return "" }
