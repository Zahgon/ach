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

// Addenda18 is an addenda which provides business transaction information for Addenda Type
// Code 18 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda18 is optional for IAT entries
//
// The Addenda18 record identifies information on each Foreign Correspondent Bank involved in the
// processing of the IAT entry. If no Foreign Correspondent Bank is involved,t he record should not be
// included. A maximum of five of these Addenda Records may be included with each IAT entry.
type Addenda18 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda18 types code '18'
	TypeCode string `json:"typeCode"`
	// ForeignCorrespondentBankName contains the name of the Foreign Correspondent Bank
	ForeignCorrespondentBankName string `json:"foreignCorrespondentBankName"`
	// Foreign Correspondent Bank Identification Number Qualifier contains a 2-digit code that
	// identifies the numbering scheme used in the Foreign Correspondent Bank Identification Number
	// field. Code values for this field are:
	// “01” = National Clearing System
	// “02” = BIC Code
	// “03” = IBAN Code
	ForeignCorrespondentBankIDNumberQualifier string `json:"foreignCorrespondentBankIDNumberQualifier"`
	// Foreign Correspondent Bank Identification Number contains the bank ID number of the Foreign
	// Correspondent Bank
	ForeignCorrespondentBankIDNumber string `json:"foreignCorrespondentBankIDNumber"`
	// Foreign Correspondent Bank Branch Country Code contains the two-character code, as approved by
	// the International Organization for Standardization (ISO), to identify the country in which the
	// branch of the Foreign Correspondent Bank is located. Values can be found on the International
	// Organization for Standardization website: www.iso.org
	ForeignCorrespondentBankBranchCountryCode string `json:"foreignCorrespondentBankBranchCountryCode"`
	// SequenceNumber is consecutively assigned to each Addenda18 Record following
	// an Entry Detail Record. The first addenda18 sequence number must always
	// be a "1".
	SequenceNumber int `json:"sequenceNumber"`
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

// NewAddenda18 returns a new Addenda18 with default values for none exported fields
func NewAddenda18() *Addenda18 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda18 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda18 *Addenda18) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 18

// 4-38 Based on the information entered (04-38) 35 alphanumeric

// 39-40  Based on the information entered (39-40) 2 alphanumeric
// “01” = National Clearing System
// “02” = BIC Code
// “03” = IBAN Code

// 41-74 Based on the information entered (41-74) 34 alphanumeric

// 75-77 Based on the information entered (75-77) 3 alphanumeric

// 78-83 - Blank space

// 84-87 SequenceNumber is consecutively assigned to each Addenda18 Record following an Entry Detail Record

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda18) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda18 struct to a 94 character string.
func (addenda18 *Addenda18) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda18 *Addenda18) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 18

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda18 *Addenda18) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// ForeignCorrespondentBankNameField returns a zero padded ForeignCorrespondentBankName string
func (addenda18 *Addenda18) ForeignCorrespondentBankNameField() string {
	_ = "STUB: not implemented"
	return ""
}

// ForeignCorrespondentBankIDNumberQualifierField returns a zero padded ForeignCorrespondentBankIDNumberQualifier string
func (addenda18 *Addenda18) ForeignCorrespondentBankIDNumberQualifierField() string {
	_ = "STUB: not implemented"
	return ""
}

// ForeignCorrespondentBankIDNumberField returns a zero padded ForeignCorrespondentBankIDNumber string
func (addenda18 *Addenda18) ForeignCorrespondentBankIDNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

// ForeignCorrespondentBankBranchCountryCodeField returns a zero padded ForeignCorrespondentBankBranchCountryCode string
func (addenda18 *Addenda18) ForeignCorrespondentBankBranchCountryCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda18 *Addenda18) SequenceNumberField() string { _ = "STUB: not implemented"; return "" }

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda18 *Addenda18) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
