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

type Addenda99Dishonored struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`

	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`

	// DishonoredReturnReasonCode is the return code explaining the dishonorment
	DishonoredReturnReasonCode string `json:"dishonoredReturnReasonCode"`

	// OriginalEntryTraceNumber is the trace number specifieid in the initial entry
	OriginalEntryTraceNumber string `json:"originalEntryTraceNumber"`

	// OriginalReceivingDFIIdentification is the DFI Identification specifieid in the initial entry
	OriginalReceivingDFIIdentification string `json:"originalReceivingDFIIdentification"`

	// ReturnTraceNumber is the TraceNumber used when issuing the return
	ReturnTraceNumber string `json:"returnTraceNumber"`

	// ReturnSettlementDate is the date of return issuing
	ReturnSettlementDate string `json:"returnSettlementDate"`

	// ReturnReasonCode is the initial return code
	ReturnReasonCode string `json:"returnReasonCode"`

	// AddendaInformation is additional data
	AddendaInformation string `json:"addendaInformation"`

	// TraceNumber is the trace number for dishonorment
	TraceNumber string `json:"traceNumber"`

	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewAddenda99Dishonored returns a new Addenda99Dishonored with default values for none exported fields
func NewAddenda99Dishonored() *Addenda99Dishonored { _ = "STUB: not implemented"; return nil }

func (Addenda99Dishonored *Addenda99Dishonored) Parse(record string) {
	_ = "STUB: not implemented"
	return
}

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// reserved

// 36-38 reserved - Leave blank

func (Addenda99Dishonored *Addenda99Dishonored) String() string {
	_ = "STUB: not implemented"
	return ""
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (Addenda99Dishonored *Addenda99Dishonored) SetValidation(opts *ValidateOpts) {
	_ = "STUB: not implemented"
	return
}

func IsDishonoredReturnCode(code string) bool { _ = "STUB: not implemented"; return false }

// Validate verifies NACHA rules for Addenda99Dishonored
func (Addenda99Dishonored *Addenda99Dishonored) Validate() error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the DishonoredReturnReasonCode matches expected values

// We can validate the Dishonored ReturnCode

func (Addenda99Dishonored *Addenda99Dishonored) DishonoredReturnReasonCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

// OriginalEntryTraceNumberField returns a zero padded TraceNumber string
func (Addenda99Dishonored *Addenda99Dishonored) OriginalEntryTraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Dishonored *Addenda99Dishonored) OriginalReceivingDFIIdentificationField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnTraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnSettlementDateField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnReasonCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Dishonored *Addenda99Dishonored) AddendaInformationField() string {
	_ = "STUB: not implemented"
	return ""
}

// TraceNumberField returns a zero padded TraceNumber string
func (Addenda99Dishonored *Addenda99Dishonored) TraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
