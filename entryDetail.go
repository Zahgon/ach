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

// EntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient's financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetail struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TransactionCode if the receivers account is checking, savings, general ledger (GL) or loan.
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
	// IdentificationNumber an internal identification (alphanumeric) that
	// you use to uniquely identify this Entry Detail Record
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string `json:"individualName"`
	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distinct one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	//
	// WEB and TEL batches use the Discretionary Data Field as the Payment Type Code
	//
	// Refused ACK AND ATX entries will use this field for the Refused Acknowledgement Code
	DiscretionaryData string `json:"discretionaryData,omitempty"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
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
	// Addenda02 for use with StandardEntryClassCode MTE, POS, and SHR
	Addenda02 *Addenda02 `json:"addenda02,omitempty"`
	// Addenda05 for use with StandardEntryClassCode: ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX.
	Addenda05 []*Addenda05 `json:"addenda05,omitempty"`
	// Addenda98 for user with Notification of Change
	Addenda98 *Addenda98 `json:"addenda98,omitempty"`
	// Addenda98 for user with Refused Notification of Change
	Addenda98Refused *Addenda98Refused `json:"addenda98Refused,omitempty"`
	// Addenda99 for use with Returns
	Addenda99 *Addenda99 `json:"addenda99,omitempty"`
	// Addenda99Contested for use with Contested Dishonored Returns
	Addenda99Contested *Addenda99Contested `json:"addenda99Contested,omitempty"`
	// Addenda99Dishonored for use with Dishonored Returns
	Addenda99Dishonored *Addenda99Dishonored `json:"addenda99Dishonored,omitempty"`
	// Category defines if the entry is a Forward, Return, or NOC
	Category string `json:"category,omitempty"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters

	validateOpts *ValidateOpts
	secCode      string
}

const (
	// CategoryForward defines the entry as being sent to the receiving institution
	CategoryForward = "Forward"
	// CategoryReturn defines the entry as being a return of a forward entry back to the originating institution
	CategoryReturn = "Return"
	// CategoryNOC defines the entry as being a notification of change of a forward entry to the originating institution
	CategoryNOC = "NOC"
	// CategoryDishonoredReturn defines the entry as being a dishonored return initiated by the ODFI to the RDFI that
	// submitted the return entry
	CategoryDishonoredReturn = "DishonoredReturn"
	// CategoryDishonoredReturnContested defines the entry as a contested dishonored return initiated by the RDFI to
	// the ODFI that submitted the dishonored return
	CategoryDishonoredReturnContested = "DishonoredReturnContested"

	// TransactionCode Values

	// CheckingCredit is a credit to the receiver's checking account
	CheckingCredit = 22
	// CheckingReturnNOCCredit is a return that credits the receiver's checking account
	CheckingReturnNOCCredit = 21
	// CheckingPrenoteCredit is a pre-notification of a credit to the receiver's checking account
	CheckingPrenoteCredit = 23
	// CheckingZeroDollarRemittanceCredit is a zero dollar remittance data credit to a checking account for CCD, CTX,
	// ACK, and ATX entries
	CheckingZeroDollarRemittanceCredit = 24
	// CheckingDebit is a debit to the receivers checking account
	CheckingDebit = 27
	// CheckingReturnNOCDebit is a return that debits the receiver's checking account
	CheckingReturnNOCDebit = 26
	// CheckingPrenoteDebit is a pre-notification of a debit to the receiver's checking account
	CheckingPrenoteDebit = 28
	// CheckingZeroDollarRemittanceDebit is a zero dollar remittance data debit to a checking account for CCD, CTX,
	// ACK, and ATX entries
	CheckingZeroDollarRemittanceDebit = 29

	// SavingsCredit is a credit to the receiver's savings account
	SavingsCredit = 32
	// SavingsReturnNOCCredit is a return that credits the receiver's savings account
	SavingsReturnNOCCredit = 31
	// SavingsPrenoteCredit is a pre-notification of a credit to the receiver's savings account
	SavingsPrenoteCredit = 33
	// SavingsZeroDollarRemittanceCredit is a zero dollar remittance data credit to a savings account for CCD
	// and CTX entries
	SavingsZeroDollarRemittanceCredit = 34
	// SavingsDebit is a debit to the receivers savings account
	SavingsDebit = 37
	// SavingsReturnNOCDebit is a return that debits the receiver's savings account
	SavingsReturnNOCDebit = 36
	// SavingsPrenoteDebit is a pre-notification of a debit to the receiver's savings account
	SavingsPrenoteDebit = 38
	// SavingsZeroDollarRemittanceDebit is a zero dollar remittance data debit to a savings account for CCD
	// and CTX entries
	SavingsZeroDollarRemittanceDebit = 39

	// GLCredit is a credit to the receiver's general ledger (GL) account
	GLCredit = 42
	// GLReturnNOCCredit is a return that credits the receiver's general ledger (GL) account
	GLReturnNOCCredit = 41
	// GLPrenoteCredit is a pre-notification of a credit to the receiver's general ledger (GL) account
	GLPrenoteCredit = 43
	// GLZeroDollarRemittanceCredit is a zero dollar remittance data credit to the receiver's general ledger (GL) account
	GLZeroDollarRemittanceCredit = 44
	// GLDebit is a debit to the receiver's general ledger (GL) account
	GLDebit = 47
	// GLReturnNOCDebit is a return that debits the receiver's general ledger (GL) account
	GLReturnNOCDebit = 46
	// GLPrenoteDebit is a pre-notification of a debit to the receiver's general ledger (GL) account
	GLPrenoteDebit = 48
	// GLZeroDollarRemittanceDebit is a zero dollar remittance data debit to the receiver's general ledger (GL) account
	GLZeroDollarRemittanceDebit = 49

	// LoanCredit is a credit to the receiver's loan account
	LoanCredit = 52
	// LoanReturnNOCCredit is a return that credits the receiver's loan account
	LoanReturnNOCCredit = 51
	// LoanPrenoteCredit is a pre-notification of a credit to the receiver's loan account
	LoanPrenoteCredit = 53
	// LoanZeroDollarRemittanceCredit is a zero dollar remittance data credit to the receiver's loan account
	LoanZeroDollarRemittanceCredit = 54
	// LoanDebit is a debit (Reversal's Only) to the receiver's loan account
	LoanDebit = 55
	// LoanReturnNOCDebit is a return that debits the receiver's loan account
	LoanReturnNOCDebit = 56
	// LoanPrenoteDebit is N/A
	// LoanZeroDollarRemittanceDebit is N/A

	// End of TransactionCode Values
)

// NewEntryDetail returns a new EntryDetail with default values for non exported fields
func NewEntryDetail() *EntryDetail { _ = "STUB: not implemented"; return nil }

func (ed *EntryDetail) SetSECCode(code string) { _ = "STUB: not implemented"; return }

// Parse takes the input record string and parses the EntryDetail values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (ed *EntryDetail) Parse(record string) { _ = "STUB: not implemented"; return }

// Precompute byte positions for all rune boundaries (0 to 94)

// Extract fields using precomputed byte positions

// String writes the EntryDetail struct to a 94 character string.
func (ed *EntryDetail) String() string { _ = "STUB: not implemented"; return "" }

// SetValidation stores ValidateOpts on the EntryDetail which are to be used to override
// the default NACHA validation rules.
func (ed *EntryDetail) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *EntryDetail) Validate() error { _ = "STUB: not implemented"; return nil }

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (ed *EntryDetail) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

const (
	// NachaEntryAmountLimit is the maximum amount allowed by the Nacha format for an entry (10 digits)
	NachaEntryAmountLimit = 99_999_999_99
)

func (ed *EntryDetail) amountOverflowsField() error { _ = "STUB: not implemented"; return nil }

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (ed *EntryDetail) SetRDFI(rdfi string) *EntryDetail { _ = "STUB: not implemented"; return nil }

// SetTraceNumber takes first 8 digits of ODFI and concatenates a sequence number onto the TraceNumber
func (ed *EntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	_ = "STUB: not implemented"
	return
}

// Populate TraceNumber of addenda records that should match the Entry's trace number

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (ed *EntryDetail) RDFIIdentificationField() string { _ = "STUB: not implemented"; return "" }

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (ed *EntryDetail) DFIAccountNumberField() string { _ = "STUB: not implemented"; return "" }

// AmountField returns a zero padded string of amount
func (ed *EntryDetail) AmountField() string { _ = "STUB: not implemented"; return "" }

// IdentificationNumberField returns a space padded string of IdentificationNumber
func (ed *EntryDetail) IdentificationNumberField() string { _ = "STUB: not implemented"; return "" }

// CheckSerialNumberField is used in RCK, ARC, BOC files but returns
// a space padded string of the underlying IdentificationNumber field
func (ed *EntryDetail) CheckSerialNumberField() string { _ = "STUB: not implemented"; return "" }

// SetCheckSerialNumber setter for RCK, ARC, BOC CheckSerialNumber
// which is underlying IdentificationNumber
func (ed *EntryDetail) SetCheckSerialNumber(s string) { _ = "STUB: not implemented"; return }

// SetPOPCheckSerialNumber setter for POP CheckSerialNumber
// which is characters 1-9 of underlying CheckSerialNumber \ IdentificationNumber
func (ed *EntryDetail) SetPOPCheckSerialNumber(s string) { _ = "STUB: not implemented"; return }

// SetPOPTerminalCity setter for POP Terminal City
// which is characters 10-13 of underlying CheckSerialNumber \ IdentificationNumber
func (ed *EntryDetail) SetPOPTerminalCity(s string) { _ = "STUB: not implemented"; return }

// SetPOPTerminalState setter for POP Terminal State
// which is characters 14-15 of underlying CheckSerialNumber \ IdentificationNumber
func (ed *EntryDetail) SetPOPTerminalState(s string) { _ = "STUB: not implemented"; return }

// POPCheckSerialNumberField is used in POP, characters 1-9 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func (ed *EntryDetail) POPCheckSerialNumberField() string { _ = "STUB: not implemented"; return "" }

// POPTerminalCityField is used in POP, characters 10-13 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func (ed *EntryDetail) POPTerminalCityField() string { _ = "STUB: not implemented"; return "" }

// POPTerminalStateField is used in POP, characters 14-15 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func (ed *EntryDetail) POPTerminalStateField() string { _ = "STUB: not implemented"; return "" }

// SetSHRCardExpirationDate format MMYY is used in SHR, characters 1-4 of underlying
// IdentificationNumber
func (ed *EntryDetail) SetSHRCardExpirationDate(s string) { _ = "STUB: not implemented"; return }

// SetSHRDocumentReferenceNumber format int is used in SHR, characters 5-15 of underlying
// IdentificationNumber
func (ed *EntryDetail) SetSHRDocumentReferenceNumber(s string) { _ = "STUB: not implemented"; return }

// SetSHRIndividualCardAccountNumber format int is used in SHR, underlying
// IndividualName
func (ed *EntryDetail) SetSHRIndividualCardAccountNumber(s string) {
	_ = "STUB: not implemented"
	return
}

// SHRCardExpirationDateField format MMYY is used in SHR, characters 1-4 of underlying
// IdentificationNumber
func (ed *EntryDetail) SHRCardExpirationDateField() string { _ = "STUB: not implemented"; return "" }

// SHRDocumentReferenceNumberField format int is used in SHR, characters 5-15 of underlying
// IdentificationNumber
func (ed *EntryDetail) SHRDocumentReferenceNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

// SHRIndividualCardAccountNumberField format int is used in SHR, underlying
// IndividualName
func (ed *EntryDetail) SHRIndividualCardAccountNumberField() string {
	_ = "STUB: not implemented"
	return ""
}

// IndividualNameField returns a space padded string of IndividualName
func (ed *EntryDetail) IndividualNameField() string { _ = "STUB: not implemented"; return "" }

// ReceivingCompanyField is used in CCD files but returns the underlying IndividualName field
func (ed *EntryDetail) ReceivingCompanyField() string { _ = "STUB: not implemented"; return "" }

// SetReceivingCompany setter for CCD ReceivingCompany which is underlying IndividualName
func (ed *EntryDetail) SetReceivingCompany(s string) { _ = "STUB: not implemented"; return }

// OriginalTraceNumberField is used in ACK and ATX files but returns the underlying IdentificationNumber field
func (ed *EntryDetail) OriginalTraceNumberField() string { _ = "STUB: not implemented"; return "" }

// SetOriginalTraceNumber setter for ACK and ATX OriginalTraceNumber which is underlying IdentificationNumber
func (ed *EntryDetail) SetOriginalTraceNumber(s string) { _ = "STUB: not implemented"; return }

// SetCATXAddendaRecords setter for CTX and ATX AddendaRecords characters 1-4 of underlying IndividualName
func (ed *EntryDetail) SetCATXAddendaRecords(i int) { _ = "STUB: not implemented"; return }

// SetCATXReceivingCompany setter for CTX and ATX ReceivingCompany characters 5-20 underlying IndividualName
// Position 21-22 of underlying Individual Name are reserved blank space for CTX "  "
func (ed *EntryDetail) SetCATXReceivingCompany(s string) { _ = "STUB: not implemented"; return }

// CATXAddendaRecordsField is used in CTX and ATX files, characters 1-4 of underlying IndividualName field
func (ed *EntryDetail) CATXAddendaRecordsField() string { _ = "STUB: not implemented"; return "" }

// CATXReceivingCompanyField is used in CTX and ATX files, characters 5-20 of underlying IndividualName field
func (ed *EntryDetail) CATXReceivingCompanyField() string { _ = "STUB: not implemented"; return "" }

// CATXReservedField is used in CTX and ATX files, characters 21-22 of underlying IndividualName field
func (ed *EntryDetail) CATXReservedField() string { _ = "STUB: not implemented"; return "" }

// DiscretionaryDataField returns a space padded string of DiscretionaryData
func (ed *EntryDetail) DiscretionaryDataField() string { _ = "STUB: not implemented"; return "" }

// PaymentTypeField returns the DiscretionaryData field used in WEB and TEL batch files
func (ed *EntryDetail) PaymentTypeField() string {
	_ = "STUB: not implemented"
	// because DiscretionaryData can be changed outside of PaymentType we reset the value for safety
	return ""
}

// SetPaymentType as R (Recurring) all other values will result in S (single).
// This is used for WEB and TEL batch files in-place of DiscretionaryData.
func (ed *EntryDetail) SetPaymentType(t string) { _ = "STUB: not implemented"; return }

// SetProcessControlField setter for TRC Process Control Field characters 1-6 of underlying IndividualName
func (ed *EntryDetail) SetProcessControlField(s string) { _ = "STUB: not implemented"; return }

// SetItemResearchNumber setter for TRC Item Research Number characters 7-22 of underlying IndividualName
func (ed *EntryDetail) SetItemResearchNumber(s string) { _ = "STUB: not implemented"; return }

// SetItemTypeIndicator setter for TRC Item Type Indicator which is underlying Discretionary Data
func (ed *EntryDetail) SetItemTypeIndicator(s string) { _ = "STUB: not implemented"; return }

// ProcessControlField getter for TRC Process Control Field characters 1-6 of underlying IndividualName
func (ed *EntryDetail) ProcessControlField() string { _ = "STUB: not implemented"; return "" }

// ItemResearchNumber getter for TRC Item Research Number characters 7-22 of underlying IndividualName
func (ed *EntryDetail) ItemResearchNumber() string { _ = "STUB: not implemented"; return "" }

// ItemTypeIndicator getter for TRC Item Type Indicator which is underlying Discretionary Data
func (ed *EntryDetail) ItemTypeIndicator() string { _ = "STUB: not implemented"; return "" }

// TraceNumberField returns a zero padded TraceNumber string
func (ed *EntryDetail) TraceNumberField() string { _ = "STUB: not implemented"; return "" }

// CreditOrDebit returns a "C" for credit or "D" for debit based on the entry TransactionCode
func (ed *EntryDetail) CreditOrDebit() string { _ = "STUB: not implemented"; return "" }

// take the second number in the TransactionCode

// AddAddenda05 appends an Addenda05 to the EntryDetail
func (ed *EntryDetail) AddAddenda05(addenda05 *Addenda05) { _ = "STUB: not implemented"; return }

// addendaCount returns the count of Addenda records added onto this EntryDetail
func (ed *EntryDetail) addendaCount() (n int) { _ = "STUB: not implemented"; return 0 }
