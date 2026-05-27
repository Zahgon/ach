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

// Addenda12 is an addenda which provides business transaction information for Addenda Type
// Code 12 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda12 is mandatory for IAT entries
//
// The Addenda12 record identifies key information related to the Originator of
// the entry.
type Addenda12 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda12 types code '12'
	TypeCode string `json:"typeCode"`
	// Originator City & State / Province
	// Data elements City and State / Province  should be separated with an asterisk (*) as a delimiter
	// and the field should end with a backslash (\).
	// For example: San Francisco*CA\
	OriginatorCityStateProvince string `json:"originatorCityStateProvince"`
	// Originator Country & Postal Code
	// Data elements must be separated by an asterisk (*) and must end with a backslash (\)
	// For example: US*10036\
	OriginatorCountryPostalCode string `json:"originatorCountryPostalCode"`
	// Originator Date of Birth
	// Example: YYYY-MM-DD
	OriginatorDateOfBirth string `json:"originatorDateOfBirth"`
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

// NewAddenda12 returns a new Addenda12 with default values for none exported fields
func NewAddenda12() *Addenda12 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda12 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda12 *Addenda12) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 12

// 4-38

// 39-73

// 74-83

// 84-87 reserved - Leave blank

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda12) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda12 struct to a 94 character string.
func (addenda12 *Addenda12) String() string { _ = "STUB: not implemented"; return "" }

// ToDo Validator for backslash

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda12 *Addenda12) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 12

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda12 *Addenda12) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// OriginatorCityStateProvinceField gets the OriginatorCityStateProvinceField left padded
func (addenda12 *Addenda12) OriginatorCityStateProvinceField() string {
	_ = "STUB: not implemented"
	return ""
}

// OriginatorCountryPostalCodeField gets the OriginatorCountryPostalCode field left padded
func (addenda12 *Addenda12) OriginatorCountryPostalCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

// OriginatorDateOfBirthField formats the OriginatorDateOfBirth field
func (addenda12 *Addenda12) OriginatorDateOfBirthField() string {
	_ = "STUB: not implemented"
	return ""
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda12 *Addenda12) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
