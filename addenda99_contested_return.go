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

type Addenda99Contested struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`

	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`

	// ContestedReturnCode is the return code explaining the contested dishonorment
	ContestedReturnCode string `json:"contestedReturnCode"`

	// OriginalEntryTraceNumber is the trace number specifieid in the initial entry
	OriginalEntryTraceNumber string `json:"originalEntryTraceNumber"`

	// DateOriginalEntryReturned is the original entry's date
	DateOriginalEntryReturned string `json:"dateOriginalEntryReturned"`

	// OriginalReceivingDFIIdentification is the DFI Identification specifieid in the initial entry
	OriginalReceivingDFIIdentification string `json:"originalReceivingDFIIdentification"`

	// OriginalSettlementDate is the initial date of settlement
	OriginalSettlementDate string `json:"originalSettlementDate"`

	// ReturnTraceNumber is the original returns trace number
	ReturnTraceNumber string `json:"returnTraceNumber"`

	// ReturnSettlementDate is the original return's settlement date
	ReturnSettlementDate string `json:"returnSettlementDate"`

	// ReturnReasonCode is the original return's code
	ReturnReasonCode string `json:"returnReasonCode"`

	// DishonoredReturnTraceNumber is the dishonorment's trace number
	DishonoredReturnTraceNumber string `json:"dishonoredReturnTraceNumber"`

	// DishonoredReturnSettlementDate is the dishonorment's settlement date
	DishonoredReturnSettlementDate string `json:"dishonoredReturnSettlementDate"`

	// DishonoredReturnReasonCode is the dishonorment's return code
	DishonoredReturnReasonCode string `json:"dishonoredReturnReasonCode"`

	// TraceNumber is the trace number for contesting
	TraceNumber string `json:"traceNumber"`

	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewAddenda99Contested returns a new Addenda99Contested with default values for none exported fields
func NewAddenda99Contested() *Addenda99Contested { _ = "STUB: not implemented"; return nil }

func (Addenda99Contested *Addenda99Contested) Parse(record string) {
	_ = "STUB: not implemented"
	return
}

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// reserved

func (Addenda99Contested *Addenda99Contested) String() string { _ = "STUB: not implemented"; return "" }

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (Addenda99Contested *Addenda99Contested) SetValidation(opts *ValidateOpts) {
	_ = "STUB: not implemented"
	return
}

// Validate verifies NACHA rules for Addenda99Contested
func (Addenda99Contested *Addenda99Contested) Validate() error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the ContestedReturnReasonCode matches expected values

// We can validate the Contested ReturnCode

func IsContestedReturnCode(code string) bool { _ = "STUB: not implemented"; return false }

func (Addenda99Contested *Addenda99Contested) ContestedReturnCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) OriginalEntryTraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) DateOriginalEntryReturnedField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) OriginalReceivingDFIIdentificationField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) OriginalSettlementDateField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) ReturnTraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) ReturnSettlementDateField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) ReturnReasonCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) DishonoredReturnTraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) DishonoredReturnSettlementDateField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) DishonoredReturnReasonCodeField() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99Contested *Addenda99Contested) TraceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}
