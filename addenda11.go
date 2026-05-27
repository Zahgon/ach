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

// Addenda11 is an addenda which provides business transaction information for Addenda Type
// Code 11 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda11 is mandatory for IAT entries
//
// The Addenda11 record identifies key information related to the Originator of
// the entry.
type Addenda11 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda11 types code '11'
	TypeCode string `json:"typeCode"`
	// Originator Name contains the originators name (your company name / name)
	OriginatorName string `json:"originatorName"`
	// Originator Street Address Contains the originators street address (your company's address / your address)
	OriginatorStreetAddress string `json:"originatorStreetAddress"`
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

// NewAddenda11 returns a new Addenda11 with default values for none exported fields
func NewAddenda11() *Addenda11 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda11 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda11 *Addenda11) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 11

// 4-38

// 39-73

// 74-87 reserved - Leave blank

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda11) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda11 struct to a 94 character string.
func (addenda11 *Addenda11) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda11 *Addenda11) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 11

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda11 *Addenda11) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// OriginatorNameField gets the OriginatorName field - Originator Company Name/Individual Name left padded
func (addenda11 *Addenda11) OriginatorNameField() string { _ = "STUB: not implemented"; return "" }

// OriginatorStreetAddressField gets the OriginatorStreetAddress field - Originator Street Address left padded
func (addenda11 *Addenda11) OriginatorStreetAddressField() string {
	_ = "STUB: not implemented"
	return ""
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda11 *Addenda11) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
