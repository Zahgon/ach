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

// Batch holds the Batch Header and Batch Control and all Entry Records
type Batch struct {
	// id is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	// This field is managed by SetID() and ID().
	id string

	Header     *BatchHeader      `json:"batchHeader"`
	Entries    []*EntryDetail    `json:"entryDetails"`
	Control    *BatchControl     `json:"batchControl"`
	ADVEntries []*ADVEntryDetail `json:"advEntryDetails,omitempty"`
	ADVControl *ADVBatchControl  `json:"advBatchControl,omitempty"`

	// offset holds the information to build an EntryDetail record which
	// balances the batch by debiting or crediting the sum of amounts in the batch.
	offset *Offset

	// category defines if the entry is a Forward, Return, or NOC
	category string
	// Converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

const (
	// ACK ACH Payment Acknowledgment - A code that indicates acknowledgment of receipt of a corporate credit payment
	// (CCD).
	ACK = "ACK"
	// ADV Automated Accounting Advice – A code that provides accounting information regarding an Entry. It is an
	// optional service.
	ADV = "ADV"
	// ARC Accounts Receivable Entry – A code that indicates a consumer check converted to a one-time ACH debit.
	// The Accounts Receivable (ARC) Entry initiates a single-entry ACH debit to customer accounts by
	// converting checks at the point of receipt through the U.S. mail, at a drop box location or in-person for
	// payment of a bill at a manned location.
	ARC = "ARC"
	// ATX Financial EDI Acknowledgment -  A code that indicates acknowledgement by the Receiving Depository Financial
	// Institution (RDFI) that a Corporate Credit Exchange (CTX) has been received.
	ATX = "ATX"
	// BOC Back Office Conversion Entry - A code that indicates single entry debit initiated at the point of purchase
	// or at a manned bill payment location to transfer funds through conversion to an ACH debit entry during back
	// office processing.
	BOC = "BOC"
	// CCD Corporate Credit or Debit Entry - A code that indicates an entry initiated by an Organization to transfer
	// funds to or from an account of that Organization or another Organization. For commercial accounts only.
	CCD = "CCD"
	// CIE Customer Initiated Entry - A code that indicates a credit entry initiated on behalf of, and upon the
	// instruction of, a consumer to transfer funds to a non-consumer Receiver.
	CIE = "CIE"
	// COR Notification of Change or Refused Notification of Change - A code used by an RDFI or ODFI when originating a
	// Notification of Change or Refused Notification of Change in automated format.
	COR = "COR"
	// CTX Corporate Trade Exchange - A code that indicates the ability to collect and disburse funds and information
	// between companies. Generally it is used by businesses paying one another for goods or services.
	CTX = "CTX"
	// DNE Death Notification Entry - A code that United States Federal agencies (e.g. Social Security) use to notify
	// depository financial institutions that the recipient of government benefit payments has died.
	DNE = "DNE"
	// ENR Automated Enrollment Entry - A code indicating enrollment of a person with an agency of the US government
	// for a depository financial institution.
	ENR = "ENR"
	// IAT International ACH Transaction - A code IAT indicating a credit or debit ACH entry that is part of a payment
	// transaction involving a financial agency's office (i.e., depository financial institution or business issuing
	// money orders) that is not located in the territorial jurisdiction of the United States. IAT entries can be made
	// to or from a corporate or consumer account and must be accompanied by seven (7) mandatory addenda records
	// identifying the name and physical address of the Originator, name and physical address of the Receiver,
	// Receiver's account number, Receiver's bank identity and reason for the payment.
	IAT = "IAT"
	// MTE Machine Transfer Entry - A code that indicates when a consumer uses their debit card at an Automated Teller
	// Machine (ATM) to withdraw cash. MTE transactions cannot be aggregated together under a single Entry.
	MTE = "MTE"
	// POP Point of Purchase Entry - A code that indicates a check presented in-person to a merchant for purchase
	// is presented as an ACH entry instead of a physical check.
	POP = "POP"
	// POS Point of Sale Entry - A code that indicates a debit entry initiated at an "electronic terminal" to a
	// consumer account of the receiver to pay an obligation incurred in a point-of-sale transaction, or to effect a
	// point-of-sale terminal cash withdrawal.
	POS = "POS"
	// PPD Prearranged Payment and Deposit Entry - A code that indicates an entry initiated by an organization based
	// on a standing or a single entry authorization to transfer funds.
	PPD = "PPD"
	// RCK Re-presented Check Entry - A code that indicates a physical check that was presented but returned because of
	//// insufficient funds may be represented as an ACH entry.
	RCK = "RCK"
	// SHR Shared Network Transaction - A code that indicates a debit Entry initiated at an "electronic terminal," as
	// that term is defined in Regulation E, to a Consumer Account of the Receiver to pay an obligation incurred in a
	// point-of-sale transaction, or to effect a point-of-sale terminal cash withdrawal. Also an adjusting or other
	// credit Entry related to such debit Entry, transfer of funds, or obligation. SHR Entries are initiated in a
	// shared network where the ODFI and RDFI have an agreement in addition to these Rules to process such Entries.
	SHR = "SHR"
	// TEL Telephone Initiated Entry - A code indicating a Telephone-Initiated consumer debit transaction. The NACHA
	// Operating Rules permit TEL entries when the originator obtains the Receiver's authorization for the debit entry
	// orally via the telephone. An entry based upon a Receiver's oral authorization must utilize the TEL
	// Standard Entry Class (SEC) Code.
	TEL = "TEL"
	// TRC Check Truncation Entry - is a code used to identify a debit entry of a truncated check.
	TRC = "TRC"
	// TRX Check Truncation Entries Exchange - used to identify a debit entry exchange of truncated checks (multiple).
	TRX = "TRX"
	// WEB Internet-Initiated/Mobile Entry - A code indicating an entry submitted pursuant to an authorization obtained
	// solely via the Internet or a mobile network. For consumer accounts only.
	WEB = "WEB"
	// XCK Destroyed Check Entry - A code indicating a debit entry initiated for destroyed check eligible items
	XCK = "XCK"
)

func (batch *Batch) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (batch *Batch) UnmarshalJSON(p []byte) error { _ = "STUB: not implemented"; return nil }

// blank out the fields of our Batch before reading

// NewBatch takes a BatchHeader and returns a matching SEC code batch type that is a batcher. Returns an error if the SEC code is not supported.
func NewBatch(bh *BatchHeader) (Batcher, error) {
	_ = "STUB: not implemented"
	return *new(Batcher), nil
}

// ConvertBatchType will take a batch object and convert it into one of the correct batch type
func ConvertBatchType(b Batch) Batcher { _ = "STUB: not implemented"; return *new(Batcher) }

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *Batch) Create() error { _ = "STUB: not implemented"; return nil }

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *Batch) Validate() error { _ = "STUB: not implemented"; return nil }

// ValidateTotals performs checks on: 1. Batch entry count 2. Batch credit/debit totals of the 3. Batch entry hash
// ValidateTotals will never modify the Batch.
//
// The first error encountered is returned.
func (batch *Batch) ValidateTotals() error { _ = "STUB: not implemented"; return nil }

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (batch *Batch) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (batch *Batch) verify() error {
	_ = "STUB: not implemented"
	// No entries in batch
	return nil
}

// verify field inclusion in all the records of the batch.

// convert the field error in to a batch error for a consistent api

// validate batch header and control codes are the same

// Company Identification in the batch header and control must match if bypassCompanyIdentificationMatch is not enabled.

// Control ODFIIdentification must be the same as batch header

// batch number header and control must match

// Control ODFIIdentification must be the same as batch header

// batch number header and control must match

// Build creates valid batch by building sequence numbers and batch control. An error is returned if
// the batch being built has invalid records.
func (batch *Batch) build() error {
	_ = "STUB: not implemented"
	// Requires a valid BatchHeader
	return nil
}

// Create record sequence numbers

// Add a sequenced TraceNumber if one is not already set. Have to keep original trance number Return and NOC entries

// Automatically set the TraceNumber if we are validating Origin and don't have custom trace numbers

// sequences don't exist in NOC or Return addenda

// build a BatchControl record

// Set Sequence Number

// build a BatchADVControl record

// SetHeader appends an BatchHeader to the Batch
func (batch *Batch) SetHeader(batchHeader *BatchHeader) { _ = "STUB: not implemented"; return }

// GetHeader returns the current Batch header
func (batch *Batch) GetHeader() *BatchHeader { _ = "STUB: not implemented"; return nil }

// SetControl appends an BatchControl to the Batch
func (batch *Batch) SetControl(batchControl *BatchControl) { _ = "STUB: not implemented"; return }

// GetControl returns the current Batch Control
func (batch *Batch) GetControl() *BatchControl { _ = "STUB: not implemented"; return nil }

// SetADVControl appends an BatchADVControl to the Batch
func (batch *Batch) SetADVControl(batchADVControl *ADVBatchControl) {
	_ = "STUB: not implemented"
	return
}

// GetADVControl returns the current Batch ADVControl
func (batch *Batch) GetADVControl() *ADVBatchControl { _ = "STUB: not implemented"; return nil }

// GetEntries returns a slice of entry details for the batch
func (batch *Batch) GetEntries() []*EntryDetail { _ = "STUB: not implemented"; return nil }

// AddEntry appends an EntryDetail to the Batch
func (batch *Batch) AddEntry(entry *EntryDetail) { _ = "STUB: not implemented"; return }

// DeleteEntries deletes all Entries from the Batch where del() == true
func (batch *Batch) DeleteEntries(del func(e *EntryDetail) bool) { _ = "STUB: not implemented"; return }

// AddADVEntry appends an ADV EntryDetail to the Batch
func (batch *Batch) AddADVEntry(entry *ADVEntryDetail) { _ = "STUB: not implemented"; return }

// DeleteADVEntries deletes all ADV Entries from the Batch where del() == true
func (batch *Batch) DeleteADVEntries(del func(e *ADVEntryDetail) bool) {
	_ = "STUB: not implemented"
	return
}

// GetADVEntries returns a slice of entry details for the batch
func (batch *Batch) GetADVEntries() []*ADVEntryDetail { _ = "STUB: not implemented"; return nil }

// Category returns batch category
func (batch *Batch) Category() string { _ = "STUB: not implemented"; return "" }

// If an Entry has NOC or Return that's the Batch's category

// ID returns the id of the batch
func (batch *Batch) ID() string {
	_ = "STUB: not implemented"

	// SetID sets the batch id
	return ""
}

func (batch *Batch) SetID(id string) {
	_ = "STUB: not implemented"

	// isFieldInclusion iterates through all the records in the batch and verifies against default fields
	return
}

func (batch *Batch) isFieldInclusion() error { _ = "STUB: not implemented"; return nil }

// Some SEC codes require IndividualName is non-blank (and non-zeros)

// Verify IndividualName is populated

// ADV File/Batch

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *Batch) isBatchEntryCount() error { _ = "STUB: not implemented"; return nil }

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *Batch) isBatchAmount() error { _ = "STUB: not implemented"; return nil }

// ToDo: Consider going back to one function for calculating BatchAmounts, but I'm not sure I want to have
// calculateBatchAmounts with ADV TransactionCodes.  In addition the smaller functions help keep the -over for
// gocyclo lower, although since we are currently at 25 (originally it was 18 or 19) it probably won't matter now
// in this case.  Based on what I see in other github go code, I'm not sure 25 is a high enough number either.
// Balancing easy to understand functions without having to create functions just for the purpose of meeting the
// -over number convinces me that it should be higher than 25.

func (batch *Batch) calculateBatchAmounts() (credit int, debit int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (batch *Batch) calculateADVBatchAmounts() (credit int, debit int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *Batch) isSequenceAscending() error { _ = "STUB: not implemented"; return nil }

// isEntryHash validates the hash by recalculating the result
func (batch *Batch) isEntryHash() error { _ = "STUB: not implemented"; return nil }

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *Batch) calculateEntryHash() int { _ = "STUB: not implemented"; return 0 }

// EntryHash is essentially the sum of all the RDFI routing numbers in the batch. If the sum exceeds 10 digits
// (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there
// are only 10.

// "Only an agency of the United States Government may originate a DNE entry" - NACHA Operating Rules
// Origination code '2' is for government agencies. Codes 21, 23, 31, and 33 are the only transaction codes
// allowed for DNEs. Tranaction codes 21 and 31 are just for returns or NOCs of the 23 and 33 codes.
// So we check that the Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *Batch) isOriginatorDNE() error { _ = "STUB: not implemented"; return nil }

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header ODFI
func (batch *Batch) isTraceNumberODFI() error { _ = "STUB: not implemented"; return nil }

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *Batch) isAddendaSequence() error { _ = "STUB: not implemented"; return nil }

// addenda without indicator flag of 1

// check if sequence is ascending

// sequences don't exist in NOC or Return addenda

// check that we are in the correct Entry Detail

// isCategory verifies that a Forward and Return Category are not in the same batch
func (batch *Batch) isCategory() error { _ = "STUB: not implemented"; return nil }

// addendaFieldInclusion verifies Addenda* Field Inclusion based on entry.Category and
// batchHeader.StandardEntryClassCode
// Forward Entries:
// MTE, POS, and SHR can only have Addenda02
// ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX can only have Addenda05
// ARC, BOC, POP, RCK, TEL, TRC, XCK cannot have Addenda02 or Addenda05
// Notification of Change:
// COR and Addenda98
// Return:
// Addenda99, Addenda99Dishonored, Addenda99Contested
func (batch *Batch) addendaFieldInclusion(entry *EntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// addendaFieldInclusionForward verifies Addenda* Field Inclusion for entry.Category Forward
func (batch *Batch) addendaFieldInclusionForward(entry *EntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// ACK, ATX, CCD, CIE, CTX, DNE, ENR WEB, PPD, TRX can only have Addenda05

// addendaFieldInclusionNOC verifies Addenda* Field Inclusion for entry.Category NOC
func (batch *Batch) addendaFieldInclusionNOC(entry *EntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// addendaFieldInclusionReturn verifies Addenda* Field Inclusion for entry.Category Return
func (batch *Batch) addendaFieldInclusionReturn(entry *EntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// do nothing, these SEC codes allow multiple Addenda05 records alongside the DishonoredReturn addenda

// do nothing, CTX allows Addenda05 records for Return

// Offset entries within a Return batch will not have an Addenda99 record as they might be
// used to zero accounting entries.
//
// See: https://github.com/moov-io/ach/issues/1010

// IsADV determines if a batch is batch type ADV - BatchADV
func (batch *Batch) IsADV() bool { _ = "STUB: not implemented"; return false }

func (batch *Batch) ValidAmountForCodes(entry *EntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// NOC entries will have a zero'd amount value

// Returned prenotes can have a zero amount, so allow returns through

// If the entry is a PRENOTE force it's amount to be zero

// ValidTranCodeForServiceClassCode validates a TransactionCode is valid for a ServiceClassCode
func (batch *Batch) ValidTranCodeForServiceClassCode(entry *EntryDetail) error {
	_ = "STUB: not implemented"
	// ADV should use ADVEntryDetail
	return nil
}

// We're unable to validate the ServiceClassCode with custom TransactionCode validation.

// Equal returns true only if two Batch (or any Batcher) objects are equal. Equality is determined by
// many of the ACH Batch and EntryDetail properties.
func (batch *Batch) Equal(other Batcher) bool {
	_ = "STUB: not implemented"
	// Some fields are intentionally not compared as they could vary between batches that would otherwise be the same.
	return false
}

// skip to next EntryDetail

// skip to next EntryDetail

// skip to next EntryDetail

// skip to next EntryDetail

// skip to next EntryDetail

// skip to next EntryDetail

// skip to next EntryDetail

// skip to next EntryDetail

// WithOffset sets the Offset information onto a Batch so that during Create a balanced offset record(s) at the end of each batch.
//
// If there are debits, there is a credit offset matching the sum of the debits. If there are credits, there is a debit offset matching
// the sum of the credits. They are mutually exclusive.
func (b *Batch) WithOffset(off *Offset) { _ = "STUB: not implemented"; return }

const offsetIndividualName = "OFFSET"

func (b *Batch) upsertOffsets() error { _ = "STUB: not implemented"; return nil }

// remove any Offset records already on the batch

// TODO(adam): Should we remove this based on checking the last element is
// debit/credit and sums to all the other elements (which are mutually exclusive to
// the last record being debit or credit)?
// See: https://github.com/moov-io/ach/issues/540

// fixup BatchControl records for our conditional after this for loop

// remove the EntryDetail

// Make sure the offset account type is valid

// Create our debit offset EntryDetail

// zero out so we don't add an empty OFFSET EntryDetail

// Create our credit offset EntryDetail

// zero out so we don't add an empty OFFSET EntryDetail

// Add both EntryDetails to our Batch and recalculate some fields

func createOffsetEntryDetail(off *Offset, batch *Batch) *EntryDetail {
	_ = "STUB: not implemented"
	return nil
}

// left empty

// aba8 returns the first 8 digits of an ABA routing number.
// If the input is invalid then an empty string is returned.
func aba8(rtn string) string { _ = "STUB: not implemented"; return "" }

// ACH server will prefix with space, 0, or 1

func lastTraceNumber(entries []*EntryDetail) int { _ = "STUB: not implemented"; return 0 }
