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

// First position of all Record Types. These codes are uniquely assigned to
// the first byte of each row in a file.
const (
	fileHeaderPos   = "1"
	batchHeaderPos  = "5"
	entryDetailPos  = "6"
	entryAddendaPos = "7"
	batchControlPos = "8"
	fileControlPos  = "9"

	// RecordLength character count of each line representing a letter in a file
	RecordLength = 94
)

// FileError is an error describing issues validating a file
type FileError struct {
	FieldName string
	Value     string
	Msg       string
}

func (e FileError) Error() string { _ = "STUB: not implemented"; return "" }

// File contains the structures of a parsed ACH File.
type File struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a file.
	ID string `json:"id"`

	Header     FileHeader     `json:"fileHeader"`
	Batches    []Batcher      `json:"batches"`
	IATBatches []IATBatch     `json:"IATBatches"`
	Control    FileControl    `json:"fileControl"`
	ADVControl ADVFileControl `json:"fileADVControl"`

	// NotificationOfChange (Notification of change) is a slice of references to BatchCOR in file.Batches
	NotificationOfChange []Batcher `json:"NotificationOfChange"`

	// ReturnEntries is a slice of references to file.Batches that contain return entries
	ReturnEntries []Batcher `json:"ReturnEntries"`

	validateOpts *ValidateOpts
}

// NewFile constructs a file template.
func NewFile() *File { _ = "STUB: not implemented"; return nil }

type file struct {
	ID string `json:"id"`
}

type fileHeader struct {
	Header FileHeader `json:"fileHeader"`
}

type fileControl struct {
	Control FileControl `json:"fileControl"`
}

type advFileControl struct {
	ADVControl ADVFileControl `json:"advFileControl"`
}

// FileFromJSON attempts to return a *File object assuming the input is valid JSON.
//
// Callers should always check for a nil-error before using the returned file.
//
// The File returned may not be valid and an error may be returned from validation.
// Invalid files may be rejected by Financial Institutions or ACH tools.
//
// Date and Time fields in formats: RFC 3339 and ISO 8601 will be parsed and rewritten
// as their YYMMDD (year, month, day) or hhmm (hour, minute) formats.
func FileFromJSON(bs []byte) (*File, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadJSONFile will consume the specified filepath and parse the contents as a JSON formatted ACH file.
func ReadJSONFile(path string) (*File, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadJSONFileWith will consume the specified filepath and parse the contents
// as a JSON formatted ACH file with custom ValidateOpts.
func ReadJSONFileWith(path string, opts *ValidateOpts) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FileFromJSONWith attempts to return a *File object assuming the input is valid JSON.
//
// It allows custom validation overrides, so the file may not be Nacha compliant
// after parsing. Invalid files may be rejected by Financial Institutions or ACH tools.
//
// Callers should always check for a nil-error before using the returned file.
//
// Date and Time fields in formats: RFC 3339 and ISO 8601 will be parsed and rewritten
// as their YYMMDD (year, month, day) or hhmm (hour, minute) formats.
func FileFromJSONWith(bs []byte, opts *ValidateOpts) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the ValidateOpts first

// read file root level

// Read FileHeader

// Build resulting file

// Overwrite various timestamps with their ACH formatted values

// Read FileControl

// Read ADVFileControl

// MarshalJSON will produce a JSON blob with the ACH file's fields and validation settings.
func (f *File) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON parses a JSON blob with ach.FileFromJSON
func (f *File) UnmarshalJSON(p []byte) error {
	_ = "STUB: not implemented"
	// merge any validate opts with the current file
	return nil
}

// Read the file

func readValidateOpts(p []byte) (*ValidateOpts, error) { _ = "STUB: not implemented"; return nil, nil }

type batchesJSON struct {
	Batches []*Batch `json:"batches"`
}

type iatBatchesJSON struct {
	IATBatches []IATBatch `json:"iatBatches"`
}

func setEntryRecordType(e *EntryDetail) { _ = "STUB: not implemented"; return }

func setADVEntryRecordType(e *ADVEntryDetail) { _ = "STUB: not implemented"; return }

func setIATEntryRecordType(e *IATEntryDetail) { _ = "STUB: not implemented"; return }

// these values need to be inferred from the json field names

// setBatchesFromJson takes bs as JSON and attempts to read out all the Batches within.
//
// We have to break this out as Batcher is an interface (and can't be read by Go's
// json struct tag decoding).
func (f *File) setBatchesFromJSON(bs []byte) error { _ = "STUB: not implemented"; return nil }

// Clear out any nil batches

// Add new batches to file

// these values need to be inferred from the json field names

// A few SEC codes don't follow the standard columns so we have to smush
// them together as the JSON doesn't support ReceivingCompany separate
// from IndividualName.

// Skip batches with no entries after filtering nulls

// Attach a batch with the correct type

// Add new iatBatches to file

// overwriteDateTimeFields will scan through fields in a File for Date / Time
// values which are not in their ACH format (YYMMDD, hhmm). It'll attempt to parse
// various formats and overwrite them to the expected values (YYMMDD, hhmm).
func (f *File) overwriteDateTimeFields() {
	_ = "STUB: not implemented"
	// Sometimes FileCreationTime is empty but FileCreationDate is populated, so set Time to 0000
	return
}

// If both Date and Time are empty use the current wall clock value

// File header

// Batches

// BatchHeader

// TODO(adam): Addenda99 has DateOfDeath which is hard to parse and overwrite with Batcher.GetEntries() copying structs

// IAT Batches

var datetimeformats = []string{
	"2006-01-02T15:04:05.999Z", // Default javascript (new Date).toISOString()
	"2006-01-02T15:04:05Z",     // ISO 8601 without milliseconds
	time.RFC3339,               // Go default
	"01/02/2006",               // DD/MM/YYYY
}

func datetimeParse(v string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Annotate each record with the number of its corresponding line in the
// Nacha file generated by Writer.Write().
// Moov writes records in the following order:
//
// fh
// [bh, [ed + addenda], bc]
// [iatbh, [ed + addenda], iatbc]
// fc
func (f *File) annotateLineNumbers() { _ = "STUB: not implemented"; return }

func annotateBatchLineNumbers(b Batcher, startIndex int) int { _ = "STUB: not implemented"; return 0 }

func annotateIATBatchLineNumbers(b *IATBatch, startIndex int) int {
	_ = "STUB: not implemented"
	return 0
}

func annotateEntryLineNumbers(ed *EntryDetail, startIndex int) int {
	_ = "STUB: not implemented"

	// Moov addenda order: 02, [05], 98, 98refused, 99, 99dishonored, 99contested
	return 0
}

func annotateADVEntryLineNumbers(ed *ADVEntryDetail, startIndex int) int {
	_ = "STUB: not implemented"
	// Moov addenda order: 99
	return 0
}

func annotateIATEntryLineNumbers(ed *IATEntryDetail, startIndex int) int {
	_ = "STUB: not implemented"

	// Moov addenda order: 10, 11, 12, 13, 14, 15, 16, [17], [18], 98, 99
	return 0
}

// Create will modify the File to tabulate and assemble it into a valid state.
// This includes setting any posting dates, sequence numbers, counts, and sums.
//
// Create requires a FileHeader and at least one Batch if validateOpts.AllowZeroBatches is false.
//
// Since each Batch may modify computable fields in the File, any calls to
// Batch.Create should be done before Create.
//
// To check if the File is Nacha compliant, call Validate or ValidateWith.
func (f *File) Create() error { _ = "STUB: not implemented"; return nil }

// Requires a valid FileHeader to build FileControl

// If AllowZeroBatches is false, require at least one Batch in the new file.

// add 2 for FileHeader/control and reset if build was called twice do to error

// create ascending batch numbers unless batch number has been provided

// sum file entry and addenda records. Assume batch.Create batch properly calculated control

// add 2 for Batch header/control + entry added count

// sum hash from batch control. Assume Batch.Build properly calculated field.

// create ascending batch numbers

// sum file entry and addenda records. Assume batch.Create batch properly calculated control

// add 2 for Batch header/control + entry added count

// sum hash from batch control. Assume Batch.Build properly calculated field.

// create FileControl from calculated values

// blocking factor of 10 is static default value in f.Header.blockingFactor.

// If greater than 10 digits, truncate

// AddBatch appends a Batch to the ach.File
func (f *File) AddBatch(batch Batcher) []Batcher { _ = "STUB: not implemented"; return nil }

// RemoveBatch will delete a given Batcher from an ach.File
func (f *File) RemoveBatch(batch Batcher) { _ = "STUB: not implemented"; return }

// AddIATBatch appends a IATBatch to the ach.File
func (f *File) AddIATBatch(iatBatch IATBatch) []IATBatch { _ = "STUB: not implemented"; return nil }

// SetHeader allows for header to be built.
func (f *File) SetHeader(h FileHeader) *File { _ = "STUB: not implemented"; return nil }

// Validate performs checks on each record according to Nacha guidelines.
// Validate will never modify the File.
//
// ValidateOpts may be set to bypass certain rules and will only be applied to the FileHeader.
// The underlying Batches and Entries on this File will use their own ValidateOpts if they are set.
//
// The first error encountered is returned.
func (f *File) Validate() error { _ = "STUB: not implemented"; return nil }

func (f *File) GetValidation() *ValidateOpts { _ = "STUB: not implemented"; return nil }

// SetValidation stores ValidateOpts on the File which are to be used to override
// the default NACHA validation rules.
func (f *File) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// ValidateOpts contains specific overrides from the default set of validations
// performed on a NACHA file, records and various fields within.
type ValidateOpts struct {
	// SkipAll will disable all validation checks of a File. It has no effect when set on records.
	SkipAll bool `json:"skipAll"`

	// RequireABAOrigin can be set to enable routing number validation
	// over the ImmediateOrigin file header field.
	RequireABAOrigin bool `json:"requireABAOrigin"`

	// BypassOriginValidation can be set to skip validation for the
	// ImmediateOrigin file header field.
	//
	// This also allows for custom TraceNumbers which aren't prefixed with
	// a routing number as required by the NACHA specification.
	BypassOriginValidation bool `json:"bypassOriginValidation"`

	// BypassDestinationValidation can be set to skip validation for the
	// ImmediateDestination file header field.
	//
	// This also allows for custom TraceNumbers which aren't prefixed with
	// a routing number as required by the NACHA specification.
	BypassDestinationValidation bool `json:"bypassDestinationValidation"`

	// CheckTransactionCode allows for custom validation of TransactionCode values
	//
	// Note: Functions cannot be serialized into/from JSON, so this check cannot be used from config files.
	CheckTransactionCode func(code int) error `json:"-"`

	// CustomTraceNumbers disables Nacha specified checks of TraceNumbers:
	// - Ascending order of trace numbers within batches
	// - Trace numbers beginning with their ODFI's routing number
	// - AddendaRecordIndicator is set correctly
	CustomTraceNumbers bool `json:"customTraceNumbers"`

	// AllowZeroBatches allows the file to have zero batches
	AllowZeroBatches bool `json:"allowZeroBatches"`

	// AllowMissingFileHeader allows a file to be read without a FileHeader record.
	AllowMissingFileHeader bool `json:"allowMissingFileHeader"`

	// AllowMissingFileControl allows a file to be read without a FileControl record.
	AllowMissingFileControl bool `json:"allowMissingFileControl"`

	// BypassCompanyIdentificationMatch allows batches in which the Company Identification field
	// in the batch header and control do not match.
	BypassCompanyIdentificationMatch bool `json:"bypassCompanyIdentificationMatch"`

	// CustomReturnCodes can be set to skip validation for the Return Code field in an Addenda99
	// This allows for non-standard/deprecated return codes (e.g. R97)
	CustomReturnCodes bool `json:"customReturnCodes"`

	// UnequalServiceClassCode skips equality checks for the ServiceClassCode in each pair of BatchHeader
	// and BatchControl records.
	UnequalServiceClassCode bool `json:"unequalServiceClassCode"`

	// AllowUnorderedBatchNumebrs allows a file to be read with unordered batch numbers.
	AllowUnorderedBatchNumbers bool `json:"allowUnorderedBatchNumbers"`

	// AllowInvalidCheckDigit allows the CheckDigit field in EntryDetail to differ from
	// the expected calculation
	AllowInvalidCheckDigit bool `json:"allowInvalidCheckDigit"`

	// UnequalAddendaCounts skips checking that Addenda Count fields match their expected and computed values.
	UnequalAddendaCounts bool `json:"unequalAddendaCounts"`

	// PreserveSpaces keeps the spacing before and after values that normally have spaces trimmed during parsing.
	PreserveSpaces bool `json:"preserveSpaces"`

	// AllowInvalidAmounts will skip verifying the Amount is valid for the TransactionCode and entry type.
	AllowInvalidAmounts bool `json:"allowInvalidAmounts"`

	// AllowZeroEntryAmount will skip enforcing the entry Amount to be non-zero
	AllowZeroEntryAmount bool `json:"allowZeroEntryAmount"`

	// AllowSpecialCharacters will permit a wider range of UTF-8 characters in alphanumeric fields
	AllowSpecialCharacters bool `json:"allowSpecialCharacters"`

	// AllowEmptyIndividualName will skip verifying IndividualName fields are populated
	// for SEC codes that require the field to be non-blank (and non-zero)
	AllowEmptyIndividualName bool `json:"allowEmptyIndividualName"`

	// BypassBatchValidation will skip validation for batches in a file and only validate file header and control info
	BypassBatchValidation bool `json:"bypassBatchValidation"`

	// SkipFileCreationValidation will skip validation of the FileCreationTime and FileCreationDate fields in a file header
	SkipFileCreationValidation bool `json:"skipFileCreationValidation"`

	// SkipBatchHeaderCompanyValidation will bypass validation of Company fields in a BatchHeader
	SkipBatchHeaderCompanyValidation bool `json:"skipBatchHeaderCompanyValidation"`
}

// merge will combine two ValidateOpts structs and keep any non-zero field values.
func (v *ValidateOpts) merge(other *ValidateOpts) *ValidateOpts {
	_ = "STUB: not implemented"
	// If either ValidateOpts is nil return the other
	return nil
}

// ValidateWith performs checks on each record according to Nacha guidelines.
// ValidateWith will never modify the File.
//
// ValidateOpts may be set to bypass certain rules and will only be applied to the FileHeader.
// opts passed in will override ValidateOpts set by SetValidation.
// The underlying Batches and Entries on this File will use their own ValidateOpts if they are set.
//
// The first error encountered is returned.
func (f *File) ValidateWith(opts *ValidateOpts) error { _ = "STUB: not implemented"; return nil }

// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.

// File contains ADV batches BatchADV

// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.

// ValidateTotals performs checks on: 1.File entry addenda counts 2. File credit/debit totals 3. File entry hash 4. File batch count
// ValidateTotals will also call the ValidateTotals function on all contained batches
// ValidateTotals will never modify the File or contained Batches.
//
// The first error encountered is returned.
func (f *File) ValidateTotals() error { _ = "STUB: not implemented"; return nil }

// isBatchCount validates that the batch count is equal to the number of batches in the file
func (f *File) isBatchCount(IsADV bool) error { _ = "STUB: not implemented"; return nil }

// isEntryAddendaCount is prepared by hashing the RDFI's 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) isEntryAddendaCount(IsADV bool) error {
	_ = "STUB: not implemented"
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types
	return nil
}

// we assume that each batch block has already validated the addenda count is accurate in batch control.

// isFileAmount The Total Debit and Credit Entry Dollar Amounts Fields contain accumulated
// Entry Detail debit and credit totals within the file
func (f *File) isFileAmount(IsADV bool) error {
	_ = "STUB: not implemented"
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types
	return nil
}

// IAT

// isEntryHash validates the hash by recalculating the result
func (f *File) isEntryHash(IsADV bool) error {
	_ = "STUB: not implemented"
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types but not ADV
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each batch.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) calculateEntryHash(IsADV bool) int {
	_ = "STUB: not implemented"
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types but not ADV
	return 0
}

// IAT

// Ensure the entry hash cannot exceed 10 digits
// If greater than 10 digits, truncate

// IsADV determines if the File is a File containing ADV batches
func (f *File) IsADV() bool { _ = "STUB: not implemented"; return false }

func (f *File) createFileADV() error {
	_ = "STUB: not implemented"
	// add 2 for FileHeader/control and reset if build was called twice do to error
	return nil
}

// create ascending batch numbers

// sum file entry and addenda records. Assume batch.Create batch properly calculated control

// add 2 for Batch header/control + entry added count

// sum hash from batch control. Assume Batch.Build properly calculated field.

// blocking factor of 10 is static default value in f.Header.blockingFactor.

// SegmentFile takes a valid ACH File and returns 2 segmented ACH Files, one ACH File containing credit entries
// and one ACH File containing debit entries.  The return is 2 Files a Credit File and Debit File, or an error.
//
// Callers should always check for a nil-error before using the returned file.
//
// The File returned may not be valid and callers should confirm with Validate. Invalid files may be rejected
// by other Financial Institutions or ACH tools.
func (f *File) SegmentFile(_ *SegmentFileConfiguration) (*File, *File, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Additional Sorting to be FI specific

func (f *File) segmentFileBatches(creditFile, debitFile *File) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the Entry to its Batch

// segmentFileIATBatches segments IAT batches debits and credits into debit and credit files
func (f *File) segmentFileIATBatches(creditFile, debitFile *File) {
	_ = "STUB: not implemented"
	return
}

// unset so Batch.build generates a TraceNumber

// createSegmentFileBatchHeader adds BatchHeader data for a debit/credit Segment File
func createSegmentFileBatchHeader(serviceClassCode int, bh *BatchHeader) *BatchHeader {
	_ = "STUB: not implemented"
	return nil
}

// ADV requires this be 0

// createSegmentFileIATBatchHeader adds IATBatchHeader data for a debit/credit Segment File
func createSegmentFileIATBatchHeader(serviceClassCode int, IATBh *IATBatchHeader) *IATBatchHeader {
	_ = "STUB: not implemented"
	return nil
}

// addFileHeaderData adds FileHeader data for a debit/credit Segment File
func (f *File) addFileHeaderData(file *File) *File { _ = "STUB: not implemented"; return nil }

// HHmm

// segmentFileBatchAddEntry adds entries to batches in a segmented file
// Applies to All SEC Codes except ADV (Automated Accounting Advice)
func segmentFileBatchAddEntry(creditBatch, debitBatch Batcher, entry *EntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// segmentFileBatchAddADVEntry adds entries to batches in a segment file for SEC Code ADV (Automated Accounting Advice)
func segmentFileBatchAddADVEntry(creditBatch Batcher, debitBatch Batcher, entry *ADVEntryDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// Validates that the batch numbers are ascending
func (f *File) isSequenceAscending() error { _ = "STUB: not implemented"; return nil }
