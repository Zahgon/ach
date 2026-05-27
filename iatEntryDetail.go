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

// IATEntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient's financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type IATEntryDetail struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TransactionCode if the receivers account is:
	// Credit (deposit) to checking account '22'
	// Prenote for credit to checking account '23'
	// Debit (withdrawal) to checking account '27'
	// Prenote for debit to checking account '28'
	// Credit to savings account '32'
	// Prenote for credit to savings account '33'
	// Debit to savings account '37'
	// Prenote for debit to savings account '38'
	TransactionCode int `json:"transactionCode"`
	// RDFIIdentification is the RDFI's routing number without the last digit.
	// Receiving Depository Financial Institution
	RDFIIdentification string `json:"RDFIIdentification"`
	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string `json:"checkDigit"`
	// AddendaRecords is the number of Addenda Records
	AddendaRecords int `json:"addendaRecords"`
	// Amount Number of cents you are debiting/crediting this account
	Amount int `json:"amount"`
	// DFIAccountNumber is the receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string `json:"DFIAccountNumber"`
	// OFACScreeningIndicator - Leave blank
	OFACScreeningIndicator string `json:"OFACScreeningIndicator"`
	// SecondaryOFACScreeningIndicator - Leave blank
	SecondaryOFACScreeningIndicator string `json:"secondaryOFACScreeningIndicator"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one or more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator int `json:"addendaRecordIndicator"`
	// TraceNumber is assigned by the ODFI or software vendor and used as part of identification.
	//
	// The format of trace numbers is the first 8 digits of the ODFI's routing number followed by
	// 7 digits chosen by the ODFI or software vendor.
	//
	// Sequentual or random numbers can be chosen. The only requirement of Nacha is unique trace
	// numbers within a batch and file.
	//
	// Trace Numbers are included in each Entry Detail Record, Corporate Entry Detail Record,
	// and addenda Record.
	//
	// In association with the Batch Number, transmission (File Creation) Date,
	// and File ID Modifier, the Trace Number uniquely identifies an entry within a given file.
	//
	// For addenda Records, the Trace Number will be identical to the Trace Number
	// in the associated Entry Detail Record, since the Trace Number is associated
	// with an entry or item rather than a physical record.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// Addenda10 is mandatory for IAT entries
	//
	// The Addenda10 Record identifies the Receiver of the transaction and the dollar amount of
	// the payment.
	Addenda10 *Addenda10 `json:"addenda10"`
	// Addenda11 is mandatory for IAT entries
	//
	// The Addenda11 record identifies key information related to the Originator of
	// the entry.
	Addenda11 *Addenda11 `json:"addenda11"`
	// Addenda12 is mandatory for IAT entries
	//
	// The Addenda12 record identifies key information related to the Originator of
	// the entry.
	Addenda12 *Addenda12 `json:"addenda12"`
	// Addenda13 is mandatory for IAT entries
	//
	// The Addenda13 contains information related to the financial institution originating the entry.
	// For inbound IAT entries, the Fourth Addenda Record must contain information to identify the
	// foreign financial institution that is providing the funding and payment instruction for
	// the IAT entry.
	Addenda13 *Addenda13 `json:"addenda13"`
	// Addenda14 is mandatory for IAT entries
	//
	// The Addenda14 identifies the Receiving financial institution holding the Receiver's account.
	Addenda14 *Addenda14 `json:"addenda14"`
	// Addenda15 is mandatory for IAT entries
	//
	// The Addenda15 record identifies key information related to the Receiver.
	Addenda15 *Addenda15 `json:"addenda15"`
	// Addenda16 is mandatory for IAt entries
	//
	// Addenda16 record identifies additional key information related to the Receiver.
	Addenda16 *Addenda16 `json:"addenda16"`
	// Addenda17 is optional for IAT entries
	//
	// This is an optional Addenda Record used to provide payment-related data. There i a maximum of up to two of these
	// Addenda Records with each IAT entry.
	Addenda17 []*Addenda17 `json:"addenda17,omitempty"`
	// Addenda18 is optional for IAT entries
	//
	// This optional addenda record is used to provide information on each Foreign Correspondent Bank involved in the
	// processing of the IAT entry. If no Foreign Correspondent Bank is involved,the record should not be included.
	// A maximum of five Addenda18 records may be included with each IAT entry.
	Addenda18 []*Addenda18 `json:"addenda18,omitempty"`
	// Addenda98 for user with NOC
	Addenda98 *Addenda98 `json:"addenda98,omitempty"`
	// Addenda99 for use with Returns
	Addenda99 *Addenda99 `json:"addenda99,omitempty"`
	// Category defines if the entry is a Forward, Return, or NOC
	Category string `json:"category,omitempty"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters

	validateOpts *ValidateOpts
}

// NewIATEntryDetail returns a new IATEntryDetail with default values for non exported fields
func NewIATEntryDetail() *IATEntryDetail { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the EntryDetail values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (iatEd *IATEntryDetail) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// do nothing, ignore "6" record type

// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37

// 4-11 the RDFI's routing number without the last digit.

// 12-12 The last digit of the RDFI's routing number

// 13-16 Number of addenda records

// 17-29 reserved - Leave blank

// 30-39 Number of cents you are debiting/crediting this account

// 40-74 The foreign receiver's account number you are crediting/debiting

// 75-76 reserved Leave blank

// 77 OFACScreeningIndicator

// 78-78 Secondary SecondaryOFACScreeningIndicator

// 79-79 1 if addenda exists 0 if it does not

// 80-94 An internal identification (alphanumeric) that you use to uniquely identify
// this Entry Detail Record This number should be unique to the transaction and will
// help identify the transaction in case of an inquiry

// String writes the EntryDetail struct to a 94 character string.
func (iatEd *IATEntryDetail) String() string { _ = "STUB: not implemented"; return "" }

// SetValidation stores ValidateOpts on the EntryDetail which are to be used to override
// the default NACHA validation rules.
func (iatEd *IATEntryDetail) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (iatEd *IATEntryDetail) Validate() error { _ = "STUB: not implemented"; return nil }

// CheckDigit calculations

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (iatEd *IATEntryDetail) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

func (iatEd *IATEntryDetail) isCorrection() bool { _ = "STUB: not implemented"; return false }

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (iatEd *IATEntryDetail) SetRDFI(rdfi string) *IATEntryDetail {
	_ = "STUB: not implemented"
	return nil
}

// SetTraceNumber takes first 8 digits of ODFI and concatenates a sequence number onto the TraceNumber
func (iatEd *IATEntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	_ = "STUB: not implemented"
	return
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (iatEd *IATEntryDetail) RDFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// AddendaRecordsField returns a zero padded AddendaRecords string
func (iatEd *IATEntryDetail) AddendaRecordsField() string { _ = "STUB: not implemented"; return "" }

// AmountField returns a zero padded string of amount
func (iatEd *IATEntryDetail) AmountField() string { _ = "STUB: not implemented"; return "" }

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (iatEd *IATEntryDetail) DFIAccountNumberField() string { _ = "STUB: not implemented"; return "" }

// OFACScreeningIndicatorField gets the OFACScreeningIndicator
func (iatEd *IATEntryDetail) OFACScreeningIndicatorField() string {
	_ = "STUB: not implemented"
	return ""
}

// SecondaryOFACScreeningIndicatorField gets the SecondaryOFACScreeningIndicator
func (iatEd *IATEntryDetail) SecondaryOFACScreeningIndicatorField() string {
	_ = "STUB: not implemented"
	return ""
}

// TraceNumberField returns a zero padded TraceNumber string
func (iatEd *IATEntryDetail) TraceNumberField() string { _ = "STUB: not implemented"; return "" }

// AddAddenda17 appends an Addenda17 to the IATEntryDetail
func (iatEd *IATEntryDetail) AddAddenda17(addenda17 *Addenda17) { _ = "STUB: not implemented"; return }

// AddAddenda18 appends an Addenda18 to the IATEntryDetail
func (iatEd *IATEntryDetail) AddAddenda18(addenda18 *Addenda18) { _ = "STUB: not implemented"; return }

func (iatEd *IATEntryDetail) addendaCount() (n int) { _ = "STUB: not implemented"; return 0 }
