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

// Addenda14 is an addenda which provides business transaction information for Addenda Type
// Code 14 in a machine readable format. It is usually formatted according to ANSI, ASC, X14 Standard.
//
// # Addenda14 is mandatory for IAT entries
//
// The Addenda14 identifies the Receiving financial institution holding the Receiver's account.
type Addenda14 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda14 types code '14'
	TypeCode string `json:"typeCode"`
	// Name of the Receiver's bank
	RDFIName string `json:"RDFIName"`
	// Receiving DFI Identification Number Qualifier
	// The 2-digit code that identifies the numbering scheme used in the
	// Receiving DFI Identification Number field:
	// 01 = National Clearing System
	// 02 = BIC Code
	// 03 = IBAN Code
	RDFIIDNumberQualifier string `json:"RDFIIDNumberQualifier"`
	// Receiving DFI Identification
	// This field contains the bank identification number of the DFI at which the
	// Receiver maintains his account.
	RDFIIdentification string `json:"RDFIIdentification"`
	// Receiving DFI Branch Country Code
	// USb” = United States
	//(“b” indicates a blank space)
	// This 3 position field contains a 2-character code as approved by the International
	// Organization for Standardization (ISO) used to identify the country in which the
	// branch of the bank that receives the entry is located. Values for other countries can
	// be found on the International Organization for Standardization website: www.iso.org
	RDFIBranchCountryCode string `json:"RDFIBranchCountryCode"`
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

// NewAddenda14 returns a new Addenda14 with default values for none exported fields
func NewAddenda14() *Addenda14 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda14 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda14 *Addenda14) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 14

// 4-38 RDFIName

// 39-40 RDFIIDNumberQualifier

// 41-74 RDFIIdentification

// 75-77

// 78-87 reserved - Leave blank

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda14) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda14 struct to a 94 character string.
func (addenda14 *Addenda14) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda14 *Addenda14) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 14

// Valid RDFI Identification Number Qualifier

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda14 *Addenda14) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// RDFINameField gets the RDFIName field left padded
func (addenda14 *Addenda14) RDFINameField() string { _ = "STUB: not implemented"; return "" }

// RDFIIDNumberQualifierField gets the RDFIIDNumberQualifier field left padded
func (addenda14 *Addenda14) RDFIIDNumberQualifierField() string {
	_ = "STUB: not implemented"
	return ""
}

// RDFIIdentificationField gets the RDFIIdentificationCode field left padded
func (addenda14 *Addenda14) RDFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// RDFIBranchCountryCodeField gets the RDFIBranchCountryCode field left padded
func (addenda14 *Addenda14) RDFIBranchCountryCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda14 *Addenda14) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
