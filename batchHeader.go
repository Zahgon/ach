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

import (
	"time"
)

// BatchHeader identifies the originating entity and the type of transactions
// contained in the batch (i.e., the standard entry class, PPD for consumer, CCD
// or CTX for corporate). This record also contains the effective date, or desired
// settlement date, for all entries contained in this batch. The settlement date
// field is not entered as it is determined by the ACH operator
type BatchHeader struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`

	// ServiceClassCode ACH Mixed Debits and Credits '200'
	// ACH Credits Only '220'
	// ACH Debits Only '225'
	ServiceClassCode int `json:"serviceClassCode"`

	// CompanyName the company originating the entries in the batch
	CompanyName string `json:"companyName"`

	// CompanyDiscretionaryData allows Originators and/or ODFIs to include codes (one or more),
	// of significance only to them, to enable specialized handling of all
	// subsequent entries in that batch. There will be no standardized
	// interpretation for the value of the field. This field must be returned
	// intact on any return entry.
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"`

	// CompanyIdentification The 9 digit FEIN number (proceeded by a predetermined
	// alpha or numeric character) of the entity in the company name field
	CompanyIdentification string `json:"companyIdentification"`

	// StandardEntryClassCode
	// Identifies the payment type (product) found within an ACH batch-using a 3-character code.
	// The SEC Code pertains to all items within batch.
	// Determines format of the detail records.
	// Determines addenda records (required or optional PLUS one or up to 9,999 records).
	// Determines rules to follow (return time frames).
	// Some SEC codes require specific data in predetermined fields within the ACH record
	StandardEntryClassCode string `json:"standardEntryClassCode"`

	// CompanyEntryDescription A description of the entries contained in the batch
	//
	//The Originator establishes the value of this field to provide a
	// description of the purpose of the entry to be displayed back to
	// the receive For example, "GAS BILL," "REG. SALARY," "INS. PREM,"
	// "SOC. SEC.," "DTC," "TRADE PAY," "PURCHASE," etc.
	//
	// This field must contain the word "REVERSAL" (left justified) when the
	// batch contains reversing entries.
	//
	// This field must contain the word "RECLAIM" (left justified) when the
	// batch contains reclamation entries.
	//
	// This field must contain the word "NONSETTLED" (left justified) when the
	// batch contains entries which could not settle.
	CompanyEntryDescription string `json:"companyEntryDescription,omitempty"`

	// CompanyDescriptiveDate currently, the Rules provide that the “Originator establishes this field as the date it
	// would like to see displayed to the Receiver for descriptive purposes.” NACHA recommends that, as desired,
	// the content of this field be formatted using the convention “SDHHMM”, where the “SD” in positions 64- 65 denotes
	// the intent for same-day settlement, and the hours and minutes in positions 66-69 denote the desired settlement
	// time using a 24-hour clock. When electing to use this convention, the ODFI would validate that the field
	// contains either.
	//
	// ODFIs at their discretion may require their Originators to further show intent for
	// same-day settlement using an optional, yet standardized, same-day indicator in the Company Descriptive Date
	// field. The Company Descriptive Date field (5 record, field 8) is an optional field with 6 positions available
	// (positions 64-69).
	CompanyDescriptiveDate string `json:"companyDescriptiveDate,omitempty"`

	// EffectiveEntryDate the date on which the entries are to settle. Format: YYMMDD (Y=Year, M=Month, D=Day)
	EffectiveEntryDate string `json:"effectiveEntryDate,omitempty"`

	// SettlementDate Leave blank, this field is inserted by the ACH operator
	SettlementDate string `json:"settlementDate,omitempty"`

	// OriginatorStatusCode refers to the ODFI initiating the Entry.
	// 0 ADV File prepared by an ACH Operator.
	// 1 This code identifies the Originator as a depository financial institution.
	// 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int `json:"originatorStatusCode"`

	//ODFIIdentification First 8 digits of the originating DFI transit routing number
	ODFIIdentification string `json:"ODFIIdentification"`

	// BatchNumber is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by
	// record.
	BatchNumber int `json:"batchNumber"`

	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// validator is composed for data validation
	validator

	// converters is composed for ACH to golang Converters
	converters

	validateOpts *ValidateOpts
}

const (
	// BatchHeader.ServiceClassCode and BatchControl.ServiceClassCode

	// MixedDebitsAndCredits indicates a batch can have debit and credit ACH entries
	MixedDebitsAndCredits = 200
	// CreditsOnly indicates a batch can only have credit ACH entries
	CreditsOnly = 220
	// DebitsOnly indicates a batch can only have debit ACH entries
	DebitsOnly = 225
	// AutomatedAccountingAdvices indicates a batch can only have Automated Accounting Advices (debit and credit)
	AutomatedAccountingAdvices = 280
)

// NewBatchHeader returns a new BatchHeader with default values for non exported fields
func NewBatchHeader() *BatchHeader { _ = "STUB: not implemented"; return nil }

// Prepared by a financial institution

// Parse takes the input record string and parses the BatchHeader values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (bh *BatchHeader) Parse(record string) { _ = "STUB: not implemented"; return }

// We're going to process the record rune-by-rune and at each field cutoff save the value.

// Append rune to buffer

// At each cutoff save the buffer and reset

// 2-4 MixedCreditsAnDebits (200), CreditsOnly (220), DebitsOnly (225)

// 5-20 Your company's name. This name may appear on the receivers' statements prepared by the RDFI.

// 21-40 Optional field you may use to describe the batch for internal accounting purposes

// 41-50 A 10-digit number assigned to you by the ODFI once they approve you to
// originate ACH files through them. This is the same as the "Immediate origin" field in File Header Record

// 51-53 If the entries are PPD (credits/debits towards consumer account), use PPD.
// If the entries are CCD (credits/debits towards corporate account), use CCD.
// The difference between the 2 SEC codes are outside of the scope of this post.

// 54-63 Your description of the transaction. This text will appear on the receivers' bank statement.
// For example: "Payroll   "

// 64-69 The date you choose to identify the transactions in YYMMDD format.
// This date may be printed on the receivers' bank statement by the RDFI

// 70-75 Date transactions are to be posted to the receivers' account.
// You almost always want the transaction to post as soon as possible, so put tomorrow's date in YYMMDD format

// 76-78 Always blank if creating batches (just fill with spaces).
// Set to file value when parsing. Julian day format.

// 79-79 Always 1

// 80-87 Your ODFI's routing number without the last digit. The last digit is simply a
// checksum digit, which is why it is not necessary

// 88-94 Sequential number of this Batch Header Record
// For example, put "1" if this is the first Batch Header Record in the file

// String writes the BatchHeader struct to a 94 character string.
func (bh *BatchHeader) String() string { _ = "STUB: not implemented"; return "" }

// Equal returns true only if two BatchHeaders are equal.
// Equality is determined by the Nacha defined fields of each record.
func (bh *BatchHeader) Equal(other *BatchHeader) bool { _ = "STUB: not implemented"; return false }

// SetValidation stores ValidateOpts on the BatchHeader which are to be used to override
// the default NACHA validation rules.
func (bh *BatchHeader) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bh *BatchHeader) Validate() error { _ = "STUB: not implemented"; return nil }

// Ensure the ServiceClassCode follows NACHA standards if we have no TransactionCode
// validation overrides. Custom TransactionCode's don't allow for standard validation.

// Originator status code 0 is used for ADV batches only

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bh *BatchHeader) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// CompanyNameField get the CompanyName left padded
func (bh *BatchHeader) CompanyNameField() string { _ = "STUB: not implemented"; return "" }

// CompanyDiscretionaryDataField get the CompanyDiscretionaryData left padded
func (bh *BatchHeader) CompanyDiscretionaryDataField() string { _ = "STUB: not implemented"; return "" }

// CompanyIdentificationField get the CompanyIdentification left padded
func (bh *BatchHeader) CompanyIdentificationField() string { _ = "STUB: not implemented"; return "" }

// CompanyEntryDescriptionField get the CompanyEntryDescription left padded
func (bh *BatchHeader) CompanyEntryDescriptionField() string { _ = "STUB: not implemented"; return "" }

// CompanyDescriptiveDateField get the CompanyDescriptiveDate left padded
func (bh *BatchHeader) CompanyDescriptiveDateField() string { _ = "STUB: not implemented"; return "" }

// EffectiveEntryDateField get the EffectiveEntryDate in YYMMDD format
func (bh *BatchHeader) EffectiveEntryDateField() string {
	_ = "STUB: not implemented"
	// ENR records require EffectiveEntryDate to be space filled. NACHA Page OR108
	return ""
}

// YYMMDD

// ODFIIdentificationField get the odfi number zero padded
func (bh *BatchHeader) ODFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// BatchNumberField get the batch number zero padded
func (bh *BatchHeader) BatchNumberField() string { _ = "STUB: not implemented"; return "" }

func (bh *BatchHeader) SettlementDateField() string { _ = "STUB: not implemented"; return "" }

func (bh *BatchHeader) LiftEffectiveEntryDate() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// YYMMDD
