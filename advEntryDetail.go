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

// ADVEntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient's financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type ADVEntryDetail struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TransactionCode representing Accounting Entries
	// Credit for ACH debits originated - 81
	// Debit for ACH credits originated - 82
	// Credit for ACH credits received 83
	// Debit for ACH debits received 84
	// Credit for ACH credits in rejected batches 85
	// Debit for ACH debits in rejected batches - 86
	// Summary credit for respondent ACH activity - 87
	// Summary debit for respondent ACH activity - 88
	TransactionCode int `json:"transactionCode"`
	// RDFIIdentification is the RDFI's routing number without the last digit.
	// Receiving Depository Financial Institution
	RDFIIdentification string `json:"RDFIIdentification"`
	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string `json:"checkDigit"`
	// DFIAccountNumber is the receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string `json:"DFIAccountNumber"`
	// Amount Number of cents you are debiting/crediting this account
	Amount int `json:"amount"`
	// AdviceRoutingNumber
	AdviceRoutingNumber string `json:"adviceRoutingNumber"`
	// FileIdentification
	FileIdentification string `json:"fileIdentification,omitempty"`
	// ACHOperatorData
	ACHOperatorData string `json:"achOperatorData,omitempty"`
	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string `json:"individualName"`
	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distinct one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	DiscretionaryData string `json:"discretionaryData,omitempty"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator int `json:"addendaRecordIndicator"`
	// ACHOperatorRoutingNumber
	ACHOperatorRoutingNumber string `json:"achOperatorRoutingNumber"`
	// JulianDay
	JulianDay int `json:"julianDay"`
	// SequenceNumber
	SequenceNumber int `json:"sequenceNumber"`
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
	// validateOpts defines optional overrides for record validation
	validateOpts *ValidateOpts
}

const (
	// ADV Transaction Code Values
	// These transaction codes represent accounting entries

	// CreditForDebitsOriginated is an accounting entry credit for ACH debits originated
	CreditForDebitsOriginated = 81
	// CreditForCreditsReceived is an accounting entry credits for ACH credits received
	CreditForCreditsReceived = 83
	// CreditForCreditsRejected is an accounting entry credit for ACH credits in rejected batches
	CreditForCreditsRejected = 85
	// CreditSummary is an accounting entry for summary credit for respondent ACH activity
	CreditSummary = 87

	// DebitForCreditsOriginated is an accounting entry debit for ACH credits originated
	DebitForCreditsOriginated = 82
	// DebitForDebitsReceived is an accounting entry debit for ACH debits received
	DebitForDebitsReceived = 84
	// DebitForDebitsRejectedBatches is an accounting entry debit for ACH debits in rejected batches
	DebitForDebitsRejectedBatches = 86
	// DebitSummary is an accounting entry for summary debit for respondent ACH activity
	DebitSummary = 88
)

// NewADVEntryDetail returns a new ADVEntryDetail with default values for non exported fields
func NewADVEntryDetail() *ADVEntryDetail { _ = "STUB: not implemented"; return nil }

// Parse takes the input record string and parses the ADVEntryDetail values
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm
// successful parsing and data validity.

// Parse ADVEntryDetail
func (ed *ADVEntryDetail) Parse(record string) { _ = "STUB: not implemented"; return }

// 1-1 Always "6"
// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37

// 4-11 the RDFI's routing number without the last digit.

// 12-12 The last digit of the RDFI's routing number

// 13-27 The receiver's bank account number you are crediting/debiting

// 28-39 Number of cents you are debiting/crediting this account

// 40-48 Advice Routing Number

// 49-53 File Identification

// 54-54 ACH Operator Data

// 55-76 Individual Name

// 77-78 allows ODFIs to include codes of significance only to them, normally blank

// 79-79 1 if addenda exists 0 if it does not

// 80-87

// 88-90

// 91-94

func (a *ADVEntryDetail) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// String writes the ADVEntryDetail struct to a 94 character string.
func (ed *ADVEntryDetail) String() string { _ = "STUB: not implemented"; return "" }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *ADVEntryDetail) Validate() error { _ = "STUB: not implemented"; return nil }

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (ed *ADVEntryDetail) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (ed *ADVEntryDetail) SetRDFI(rdfi string) *ADVEntryDetail {
	_ = "STUB: not implemented"
	return nil
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (ed *ADVEntryDetail) RDFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (ed *ADVEntryDetail) DFIAccountNumberField() string { _ = "STUB: not implemented"; return "" }

// AmountField returns a zero padded string of amount
func (ed *ADVEntryDetail) AmountField() string { _ = "STUB: not implemented"; return "" }

// AdviceRoutingNumberField gets the AdviceRoutingNumber with zero padding
func (ed *ADVEntryDetail) AdviceRoutingNumberField() string { _ = "STUB: not implemented"; return "" }

// FileIdentificationField returns a space padded string of FileIdentification
func (ed *ADVEntryDetail) FileIdentificationField() string { _ = "STUB: not implemented"; return "" }

// ACHOperatorDataField returns a space padded string of ACHOperatorData
func (ed *ADVEntryDetail) ACHOperatorDataField() string { _ = "STUB: not implemented"; return "" }

// IndividualNameField returns a space padded string of IndividualName
func (ed *ADVEntryDetail) IndividualNameField() string { _ = "STUB: not implemented"; return "" }

// DiscretionaryDataField returns a space padded string of DiscretionaryData
func (ed *ADVEntryDetail) DiscretionaryDataField() string { _ = "STUB: not implemented"; return "" }

// ACHOperatorRoutingNumberField returns a space padded string of ACHOperatorRoutingNumber
func (ed *ADVEntryDetail) ACHOperatorRoutingNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

// JulianDateDayField returns a zero padded string of JulianDay
func (ed *ADVEntryDetail) JulianDateDayField() string { _ = "STUB: not implemented"; return "" }

// SequenceNumberField returns a zero padded string of SequenceNumber
func (ed *ADVEntryDetail) SequenceNumberField() string { _ = "STUB: not implemented"; return "" }
