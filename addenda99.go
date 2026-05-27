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

// When a Return Entry is prepared, the original Company/Batch Header Record, the original Entry Detail Record,
// and the Company/Batch Control Record are copied for return to the Originator.
//
// The Return Entry is a new Entry. These Entries must be assigned new batch and trace numbers, new identification numbers for the returning institution,
// appropriate transaction codes, etc., as required per format specifications.
//
// See Appendix Four: Return Entries in the NACHA Corporate

var (
	returnCodeDict = map[string]*ReturnCode{}
)

func init() {
	// populate the ReturnCode map with lookup values
	returnCodeDict = makeReturnCodeDict()
}

// Addenda99 utilized for Notification of Change Entry (COR) and Return types.
type Addenda99 struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`
	// ReturnCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for returning an Entry.
	// Must exist in returnCodeDict
	ReturnCode string `json:"returnCode"`
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace"`
	// DateOfDeath The field date of death is to be supplied on Entries being returned for reason of death (return reason codes R14 and R15). Format: YYMMDD (Y=Year, M=Month, D=Day)
	DateOfDeath string `json:"dateOfDeath"`
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`
	// AddendaInformation
	AddendaInformation string `json:"addendaInformation,omitempty"`
	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// ReturnCode holds a return Code, Reason/Title, and Description
//
// Table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type ReturnCode struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

// NewAddenda99 returns a new Addenda99 with default values for none exported fields
func NewAddenda99() *Addenda99 { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the Addenda99 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (Addenda99 *Addenda99) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 1-1 Always 7

// 2-3 Defines the specific explanation and format for the addenda information contained in the same record

// 4-6

// 7-21

// 22-27, might be a date or blank

// 28-35

// 36-79

// 80-94

// String writes the Addenda99 struct to a 94 character string
func (Addenda99 *Addenda99) String() string { _ = "STUB: not implemented"; return "" }

// Validate verifies NACHA rules for Addenda99
func (Addenda99 *Addenda99) Validate() error { _ = "STUB: not implemented"; return nil }

// Return Addenda requires a valid ReturnCode

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (Addenda99 *Addenda99) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// OriginalTraceField returns a zero padded OriginalTrace string
func (Addenda99 *Addenda99) OriginalTraceField() string { _ = "STUB: not implemented"; return "" }

// DateOfDeathField returns a space padded DateOfDeath string
func (Addenda99 *Addenda99) DateOfDeathField() string {
	_ = "STUB: not implemented"
	// Return space padded 6 characters if it is a zero value of DateOfDeath
	return ""
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (Addenda99 *Addenda99) OriginalDFIField() string { _ = "STUB: not implemented"; return "" }

// AddendaInformationField returns a space padded AddendaInformation string
func (Addenda99 *Addenda99) AddendaInformationField() string { _ = "STUB: not implemented"; return "" }

// IATPaymentAmount sets original forward entry payment amount characters 1-10 of underlying AddendaInformation
func (Addenda99 *Addenda99) IATPaymentAmount(s string) { _ = "STUB: not implemented"; return }

// IATAddendaInformation sets Addenda Information for IAT return items, characters 10-44 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATAddendaInformation(s string) { _ = "STUB: not implemented"; return }

// IATPaymentAmountField returns original forward entry payment amount int, characters 1-10 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATPaymentAmountField() int { _ = "STUB: not implemented"; return 0 }

// IATAddendaInformationField returns a space padded AddendaInformation string, characters 10-44 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATAddendaInformationField() string {
	_ = "STUB: not implemented"
	return ""
}

// TraceNumberField returns a zero padded TraceNumber string
func (Addenda99 *Addenda99) TraceNumberField() string { _ = "STUB: not implemented"; return "" }

// format: reserved (3), return trace number (15), return settlement date (3), return reason (2), addenda (21)
// Ref: https://www.nachaoperatingrulesonline.org/2.16334/s020
func (Addenda99 *Addenda99) SetDishonoredAddendaInformation(
	returnTraceNumber string,
	returnSettlementDate string,
	returnReasonCode string,
	addenda string,
) {
	_ = "STUB: not implemented"
	// This record drops the "R"
	return
}

// Ref: https://www.nachaoperatingrulesonline.org/2.16334/s020
func (Addenda99 *Addenda99) SetContestedAddendaInformation(
	originalSettlementDate string,
	returnTraceNumber string,
	returnSettlementDate string,
	returnReasonCode string,
	dishonoredReturnTraceNumber string,
	dishonoredReturnSettlementDate string,
	dishonoredReturnReasonCode string,
) {
	_ = "STUB: not implemented"
	// This record drops the "R"
	return
}

func (Addenda99 *Addenda99) AddendaInformationReturnTraceNumber() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99 *Addenda99) AddendaInformationReturnSettlementDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99 *Addenda99) AddendaInformationReturnReasonCode() string {
	_ = "STUB: not implemented"
	return ""
}

func (Addenda99 *Addenda99) AddendaInformationExtra() string { _ = "STUB: not implemented"; return "" }

func (Addenda99 *Addenda99) SetOriginalEntryReturnDate(date string) {
	_ = "STUB: not implemented"
	return
}

func (Addenda99 *Addenda99) OriginalEntryReturnDate() string { _ = "STUB: not implemented"; return "" }

// ReturnCodeField gives the ReturnCode struct for the given Addenda99 record
func (Addenda99 *Addenda99) ReturnCodeField() *ReturnCode { _ = "STUB: not implemented"; return nil }

// LookupReturnCode will return a struct representing the reason and description for
// the provided NACHA return code.
func LookupReturnCode(code string) *ReturnCode { _ = "STUB: not implemented"; return nil }

func makeReturnCodeDict() map[string]*ReturnCode { _ = "STUB: not implemented"; return nil }

// Return Reason Codes for RDFIs

// R03 may not be used to return ARC, BOC or POP entries solely because they do not contain an Individual Name.

// R07 Prohibited use for ARC, BOC, POP and RCK.

// Return Codes to be used for ENR entries and are initiated by a Federal Government Agency

// Return Codes to be used for RCK entries only and are initiated by a RDFI

// Return Codes to be used by the ODFI for dishonored return entries

// Return Codes to be used by the RDFI for contested dishonored return entries

//Return Codes to be used by Gateways for the return of international payments

// Additional Codes to be used by RDFI's for the Return of Entries

// populate the map
