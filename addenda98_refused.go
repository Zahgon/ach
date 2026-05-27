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

type Addenda98Refused struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`

	// TypeCode Addenda types code '98'
	TypeCode string `json:"typeCode"`

	// RefusedChangeCode is the code specifying why the Notification of Change is being refused.
	RefusedChangeCode string `json:"refusedChangeCode"`

	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace"`

	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the
	// forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`

	// CorrectedData is the corrected data
	CorrectedData string `json:"correctedData"`

	// ChangeCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for a change Entry.
	ChangeCode string `json:"changeCode"`

	// TraceSequenceNumber is the last seven digits of the TraceNumber in the original Notification of Change
	TraceSequenceNumber string `json:"traceSequenceNumber"`

	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber"`

	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda98Refused returns an reference to an instantiated Addenda98Refused with default values
func NewAddenda98Refused() *Addenda98Refused { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda98Refused values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda98Refused *Addenda98Refused) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Always "98"

// Positions 22-27 are Reserved

// Positions 75-79 are Reserved

// String writes the Addenda98 struct to a 94 character string
func (addenda98Refused *Addenda98Refused) String() string { _ = "STUB: not implemented"; return "" }

// Validate verifies NACHA rules for Addenda98
func (addenda98Refused *Addenda98Refused) Validate() error { _ = "STUB: not implemented"; return nil }

// Type Code must be 98

// RefusedChangeCode must be valid

// Addenda98 Record must contain the corrected information corresponding to the Change Code used

// ChangeCode must be valid

// TraceSequenceNumber must be valid

func (addenda98Refused *Addenda98Refused) RefusedChangeCodeField() *ChangeCode {
	_ = "STUB: not implemented"
	return nil
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (addenda98Refused *Addenda98Refused) OriginalTraceField() string {
	_ = "STUB: not implemented"
	return ""
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (addenda98Refused *Addenda98Refused) OriginalDFIField() string {
	_ = "STUB: not implemented"
	return ""
}

// CorrectedDataField returns a space padded CorrectedData string
func (addenda98Refused *Addenda98Refused) CorrectedDataField() string {
	_ = "STUB: not implemented"
	return ""
}

func (addenda98Refused *Addenda98Refused) ChangeCodeField() *ChangeCode {
	_ = "STUB: not implemented"
	return nil
}

func (addenda98Refused *Addenda98Refused) TraceSequenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

// TraceNumberField returns a zero padded traceNumber string
func (addenda98Refused *Addenda98Refused) TraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
