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
	"strings"
)

// FileHeader is a Record designating physical file characteristics and identify
// the origin (sending point) and destination (receiving point) of the entries
// contained in the file. The file header also includes creation date and time
// fields which can be used to uniquely identify a file.
type FileHeader struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`

	// PriorityCode consists of the numerals 01
	priorityCode string

	// ImmediateDestination contains the Routing Number of the ACH Operator or receiving
	// point to which the file is being sent. The ach file format specifies a 10 character
	// field begins with a blank space in the first position, followed by the four digit
	// Federal Reserve Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC). ImmediateDestinationField will append the blank space to the
	// routing number.
	ImmediateDestination string `json:"immediateDestination"`

	// ImmediateOrigin contains the Routing Number of the ACH Operator or sending
	// point that is sending the file. The ach file format specifies a 10 character
	// field begins with a blank space in the first position, followed by the four digit
	// Federal Reserve Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC). ImmediateOriginField will append the blank space to the
	// routing number.
	ImmediateOrigin string `json:"immediateOrigin"`

	// FileCreationDate is the date on which the file is prepared by an ODFI (ACH input files)
	// or the date (exchange date) on which a file is transmitted from ACH Operator
	// to ACH Operator, or from ACH Operator to RDFIs (ACH output files).
	//
	// The format is: YYMMDD. Y=Year, M=Month, D=Day
	FileCreationDate string `json:"fileCreationDate"`

	// FileCreationTime is the system time when the ACH file was created.
	//
	// The format is: HHmm. H=Hour, m=Minute
	FileCreationTime string `json:"fileCreationTime"`

	// This field should start at zero and increment by 1 (up to 9) and then go to
	// letters starting at A through Z for each subsequent file that is created for
	// a single system date. (34-34) 1 numeric 0-9 or uppercase alpha A-Z.
	// I have yet to see this ID not A
	FileIDModifier string `json:"fileIDModifier,omitempty"`

	// RecordSize indicates the number of characters contained in each
	// record. At this time, the value "094" must be used.
	recordSize string

	// BlockingFactor defines the number of physical records within a block
	// (a block is 940 characters). For all files moving between a DFI and an ACH
	// Operator (either way), the value "10" must be used. If the number of records
	// within the file is not a multiple of ten, the remainder of the block must
	// be nine-filled.
	blockingFactor string

	// FormatCode a code to allow for future format variations. As
	// currently defined, this field will contain a value of "1".
	FormatCode string

	// ImmediateDestinationName us the name of the ACH or receiving point for which that
	// file is destined. Name corresponding to the ImmediateDestination
	ImmediateDestinationName string `json:"immediateDestinationName"`

	// ImmediateOriginName is the name of the ACH operator or sending point that is
	// sending the file. Name corresponding to the ImmediateOrigin
	ImmediateOriginName string `json:"immediateOriginName"`

	// ReferenceCode is reserved for information pertinent to the Originator.
	ReferenceCode string `json:"referenceCode,omitempty"`

	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewFileHeader returns a new FileHeader with default values for none exported fields
func NewFileHeader() FileHeader { _ = "STUB: not implemented"; return *new(FileHeader) }

// Parse takes the input record string and parses the FileHeader values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (fh *FileHeader) Parse(record string) { _ = "STUB: not implemented"; return }

// (character position 1-1) Always "1"
// (2-3) Always "01"

// (4-13) A blank space followed by your ODFI's routing number. For example: " 121140399"

// (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them

// 24-29 Today's date in YYMMDD format
// must be after today's date.

// 30-33 The current time in HHmm format

// 35-37 Always "A"

// 35-37 always "094"

// 38-39 always "10"

// 40 always "1"

// 41-63 The name of the ODFI. example "SILICON VALLEY BANK    "

// 64-86 ACH operator or sending point that is sending the file

// 87-94 Optional field that may be used to describe the ACH file for internal accounting purposes

func trimRoutingNumberLeadingZero(s string) string { _ = "STUB: not implemented"; return "" }

// trim off a leading 0 as ImmediateOriginField or ImmediateDestinationField will pad it back

// String writes the FileHeader struct to a 94 character string.
func (fh *FileHeader) String() string { _ = "STUB: not implemented"; return "" }

// SetValidation stores ValidateOpts on the FileHeader which are to be used to override
// the default NACHA validation rules.
func (fh *FileHeader) SetValidation(opts *ValidateOpts) { _ = "STUB: not implemented"; return }

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (fh *FileHeader) Validate() error { _ = "STUB: not implemented"; return nil }

var (
	zeroRoutingNumber9  = strings.Repeat("0", 9)
	zeroRoutingNumber10 = strings.Repeat("0", 10)
)

// ValidateWith performs NACHA format rule checks on each record according to their specification
// overlayed with any custom flags.
// The first error encountered is returned and stops the parsing.
func (fh *FileHeader) ValidateWith(opts *ValidateOpts) error { _ = "STUB: not implemented"; return nil }

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (fh *FileHeader) fieldInclusion() error { _ = "STUB: not implemented"; return nil }

// ImmediateDestinationField gets the immediate destination number with zero padding
func (fh *FileHeader) ImmediateDestinationField() string { _ = "STUB: not implemented"; return "" }

// ImmediateOriginField gets the immediate origin number with 0 padding
func (fh *FileHeader) ImmediateOriginField() string { _ = "STUB: not implemented"; return "" }

// FileCreationDateField gets the file creation date in YYMMDD (year, month, day) format
// A blank string is returned when an error occurred while parsing the timestamp. ISO 8601
// is the only other format supported.
func (fh *FileHeader) FileCreationDateField() string { _ = "STUB: not implemented"; return "" }

// YYMMDD

// YYMMDD

// FileCreationTimeField gets the file creation time in HHmm (hour, minute) format
// A blank string is returned when an error occurred while parsing the timestamp. ISO 8601
// is the only other format supported.
func (fh *FileHeader) FileCreationTimeField() string { _ = "STUB: not implemented"; return "" }

// HHmm

// HHmm

// ImmediateDestinationNameField gets the ImmediateDestinationName field padded
func (fh *FileHeader) ImmediateDestinationNameField() string { _ = "STUB: not implemented"; return "" }

// ImmediateOriginNameField gets the ImmImmediateOriginName field padded
func (fh *FileHeader) ImmediateOriginNameField() string { _ = "STUB: not implemented"; return "" }

// ReferenceCodeField gets the ReferenceCode field padded
func (fh *FileHeader) ReferenceCodeField() string { _ = "STUB: not implemented"; return "" }
