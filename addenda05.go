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

// Addenda05 is a Addendumer addenda which provides business transaction information for Addenda Type
// Code 05 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
// It is used for the following StandardEntryClassCode: ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX.
type Addenda05 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda05 types code '05'
	TypeCode string `json:"typeCode"`
	// PaymentRelatedInformation
	PaymentRelatedInformation string `json:"paymentRelatedInformation"`
	// SequenceNumber is consecutively assigned to each Addenda05 Record following
	// an Entry Detail Record. The first addenda05 sequence number must always
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

// NewAddenda05 returns a new Addenda05 with default values for none exported fields
func NewAddenda05() *Addenda05 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda05 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda05 *Addenda05) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always 05

// 4-83 Based on the information entered (04-83) 80 alphanumeric

// 84-87 SequenceNumber is consecutively assigned to each Addenda05 Record following
// an Entry Detail Record

// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record

func (a *Addenda05) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the Addenda05 struct to a 94 character string.
func (addenda05 *Addenda05) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda05 *Addenda05) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 05

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda05 *Addenda05) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// PaymentRelatedInformationField returns a zero padded PaymentRelatedInformation string
func (addenda05 *Addenda05) PaymentRelatedInformationField() string {
	_ = "STUB: not implemented"
	return ""
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda05 *Addenda05) SequenceNumberField() string { _ = "STUB: not implemented"; return "" }

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda05 *Addenda05) EntryDetailSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
