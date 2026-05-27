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

// Addenda98 is a Addendumer addenda record format for Notification OF Change(98)
// The field contents for Notification of Change Entries must match the field contents of the original Entries
type Addenda98 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda types code '98'
	TypeCode string `json:"typeCode"`
	// ChangeCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for a change Entry.
	// Must exist in changeCodeDict
	ChangeCode string `json:"changeCode"`
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace"`
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`
	// CorrectedData is the corrected data
	CorrectedData string `json:"correctedData"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// iatCorrectedData is a field containing the additional space allowed in IAT Correction Addenda records
	iatCorrectedData string

	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

var (
	changeCodeDict = map[string]*ChangeCode{}
)

func init() {
	// populate the changeCode map with lookup values
	changeCodeDict = makeChangeCodeDict()
}

// ChangeCode holds a change Code, Reason/Title, and Description
// table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type ChangeCode struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

// NewAddenda98 returns an reference to an instantiated Addenda98 with default values
func NewAddenda98() *Addenda98 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda98 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda98 *Addenda98) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always "98"

// 4-6

// 7-21

// 28-35

// 36-64

// 65-70 (Reserved for all except IAT Corrections)

// Reserved

// 80-94

// String writes the Addenda98 struct to a 94 character string
func (addenda98 *Addenda98) String() string { _ = "STUB: not implemented"; return "" }

// 6 char reserved field

// 15 char reserved field

// Validate verifies NACHA rules for Addenda98
func (addenda98 *Addenda98) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 98

// Addenda98 requires a valid ChangeCode

// Addenda98 Record must contain the corrected information corresponding to the Change Code used

// OriginalTraceField returns a zero padded OriginalTrace string
func (addenda98 *Addenda98) OriginalTraceField() string { _ = "STUB: not implemented"; return "" }

// OriginalDFIField returns a zero padded OriginalDFI string
func (addenda98 *Addenda98) OriginalDFIField() string { _ = "STUB: not implemented"; return "" }

// CorrectedDataField returns a space padded CorrectedData string
func (addenda98 *Addenda98) CorrectedDataField() string { _ = "STUB: not implemented"; return "" }

// IATCorrectedDataField returns a space padded CorrectedData string for IAT entries,
// which is a slightly larger field than typical CorrectedData values.
func (addenda98 *Addenda98) IATCorrectedDataField() string { _ = "STUB: not implemented"; return "" }

// TraceNumberField returns a zero padded traceNumber string
func (addenda98 *Addenda98) TraceNumberField() string { _ = "STUB: not implemented"; return "" }

func (addenda98 *Addenda98) ChangeCodeField() *ChangeCode { _ = "STUB: not implemented"; return nil }

// LookupChangeCode will return a struct representing the reason and description for
// the provided NACHA change code.
func LookupChangeCode(code string) *ChangeCode { _ = "STUB: not implemented"; return nil }

func makeChangeCodeDict() map[string]*ChangeCode { _ = "STUB: not implemented"; return nil }

// Change codes used when refusing a Notification of Change

// populate the map

func IsRefusedChangeCode(code string) bool { _ = "STUB: not implemented"; return false }

// CorrectedData is a struct returned from our helper method for parsing the NOC/COR
// corrected data from Addenda98 records.
//
// All fields are optional and a valid code may not have populated data in this struct.
type CorrectedData struct {
	AccountNumber   string
	RoutingNumber   string
	Name            string
	TransactionCode int
	Identification  string
}

type correctedDataOptions struct {
	ReturnPartialData bool
}

type correctedDataOption func(conf *correctedDataOptions)

func PartialCorrectedData() correctedDataOption {
	_ = "STUB: not implemented"
	return *new(correctedDataOption)
}

// ParseCorrectedData returns a struct with some fields filled in depending on the Addenda98's
// Code and CorrectedData. Fields are trimmed when populated in this struct.
func (addenda98 *Addenda98) ParseCorrectedData(options ...correctedDataOption) *CorrectedData {
	_ = "STUB: not implemented"
	return nil
}

// Incorrect DFI Account Number

// Incorrect Routing Number

// Incorrect Routing Number and Incorrect DFI Account Number

// Return partial data only if we're asked to

// Incorrect Individual Name

// Incorrect Transaction Code

// Incorrect DFI Account Number and Incorrect Transaction Code

// Return partial data only if we're asked to

// Incorrect Routing Number, Incorrect DFI Account Number, and Incorrect Tranaction Code

// Return nothing if we have extra data

// Accumulate part by part

// Return partial data only if we're asked to

// Incorrect Individual Identification Number

// The Code/Correction is either unsupported or wasn't parsed correctly

func first(size int, data string) string { _ = "STUB: not implemented"; return "" }

const correctedDataCharLength = 29

// ParseCorrectedData returns the string properlty formatted and justified for an
// Addenda98.CorrectedData field. The code must be an official NACHA change code.
func WriteCorrectionData(code string, data *CorrectedData) string {
	_ = "STUB: not implemented"
	return ""
}
