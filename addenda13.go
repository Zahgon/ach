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

// Addenda13 is an addenda which provides business transaction information for Addenda Type
// Code 13 in a machine readable format. It is usually formatted according to ANSI, ASC, X13 Standard.
//
// # Addenda13 is mandatory for IAT entries
//
// The Addenda13 contains information related to the financial institution originating the entry.
// For inbound IAT entries, the Fourth Addenda Record must contain information to identify the
// foreign financial institution that is providing the funding and payment instruction for the IAT entry.
type Addenda13 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda13 types code '13'
	TypeCode string `json:"typeCode"`
	// Originating DFI Name
	// For Outbound IAT Entries, this field must contain the name of the U.S. ODFI.
	// For Inbound IATs: Name of the foreign bank providing funding for the payment transaction
	ODFIName string `json:"ODFIName"`
	// Originating DFI Identification Number Qualifier
	// For Inbound IATs: The 2-digit code that identifies the numbering scheme used in the
	// Foreign DFI Identification Number field:
	// 01 = National Clearing System
	// 02 = BIC Code
	// 03 = IBAN Code
	ODFIIDNumberQualifier string `json:"ODFIIDNumberQualifier"`
	// Originating DFI Identification
	// This field contains the routing number that identifies the U.S. ODFI initiating the entry.
	// For Inbound IATs: This field contains the bank ID number of the Foreign Bank providing funding
	// for the payment transaction.
	ODFIIdentification string `json:"ODFIIdentification"`
	// Originating DFI Branch Country Code
	// USb” = United States
	//(“b” indicates a blank space)
	// For Inbound IATs: This 3 position field contains a 2-character code as approved by the
	// International Organization for Standardization (ISO) used to identify the country in which
	// the branch of the bank that originated the entry is located. Values for other countries can
	// be found on the International Organization for Standardization website: www.iso.org.
	ODFIBranchCountryCode string `json:"ODFIBranchCountryCode"`
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

// NewAddenda13 returns a new Addenda13 with default values for none exported fields
func NewAddenda13() *Addenda13 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda13 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda13 *Addenda13) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 13

// 4-38 ODFIName

// 39-40 ODFIIDNumberQualifier

// 41-74 ODFIIdentification

// 75-77

// 78-87 reserved - Leave blank

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda13) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda13 struct to a 94 character string.
func (addenda13 *Addenda13) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda13 *Addenda13) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 13

// Valid ODFI Identification Number Qualifier

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda13 *Addenda13) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// ODFINameField gets the ODFIName field left padded
func (addenda13 *Addenda13) ODFINameField() string { _ = "STUB: not implemented"; return "" }

// ODFIIDNumberQualifierField gets the ODFIIDNumberQualifier field left padded
func (addenda13 *Addenda13) ODFIIDNumberQualifierField() string {
	_ = "STUB: not implemented"
	return ""
}

// ODFIIdentificationField gets the ODFIIdentificationCode field left padded
func (addenda13 *Addenda13) ODFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// ODFIBranchCountryCodeField gets the ODFIBranchCountryCode field left padded
func (addenda13 *Addenda13) ODFIBranchCountryCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda13 *Addenda13) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
