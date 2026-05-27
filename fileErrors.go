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
	"errors"
)

var (
	// ErrFileTooLong is the error given when a file exceeds the maximum possible length
	ErrFileTooLong = errors.New("file exceeds maximum possible number of lines")
	// ErrFileHeader is the error given if there is the wrong number of file headers
	ErrFileHeader = errors.New("none or more than one file headers exists")
	// ErrFileControl is the error given if there is the wrong number of file control records
	ErrFileControl = errors.New("none or more than one file control exists")
	// ErrMisplacedFileHeader is the error given when a file header is the non-first record
	ErrMisplacedFileHeader = errors.New("file header is not the first record")
	// ErrExtraRecordsAfterFileControl is the error given when records are found after a FileControl
	ErrExtraRecordsAfterFileControl = errors.New("records found after file control")
	// ErrFileEntryOutsideBatch is the error given if an entry is outside of a batch
	ErrFileEntryOutsideBatch = errors.New("entry outside of batch")
	// ErrFileAddendaOutsideBatch is the error given if an addenda is outside of a batch
	ErrFileAddendaOutsideBatch = errors.New("addenda outside of batch")
	// ErrFileAddendaOutsideEntry is the error given if an addenda is outside of an entry
	ErrFileAddendaOutsideEntry = errors.New("addenda outside of entry")
	// ErrFileBatchControlOutsideBatch is the error given if a batch control record is outside of a batch
	ErrFileBatchControlOutsideBatch = errors.New("batch control outside of batch")
	// ErrFileConsecutiveBatchHeaders is the error given when multiple batch header records occur in sequence
	ErrFileConsecutiveBatchHeaders = errors.New("consecutive Batch Headers in file")
	// ErrFileADVOnly is the error given if an ADV only file has a non-ADV batch
	ErrFileADVOnly = errors.New("file can only have ADV Batches")
	// ErrFileIATSEC is the error given if an IAT batch uses the normal NewBatch
	ErrFileIATSEC = errors.New("IAT Standard Entry Class Code should use iatBatch")
	// ErrFileNoBatches is the error given if a file has no batches
	ErrFileNoBatches = errors.New("must have []*Batches or []*IATBatches to be built")

	ErrInvalidJSON = errors.New("invalid JSON")
)

// RecordWrongLengthErr is the error given when a record is the wrong length
type RecordWrongLengthErr struct {
	Message string
	Length  int
}

// NewRecordWrongLengthErr creates a new error of the RecordWrongLengthErr type
func NewRecordWrongLengthErr(length int) RecordWrongLengthErr {
	_ = "STUB: not implemented"
	return *new(RecordWrongLengthErr)
}

func (e RecordWrongLengthErr) Error() string {
	_ = "STUB: not implemented"

	// ErrUnknownRecordType is the error given when a record does not have a known type
	return ""
}

type ErrUnknownRecordType struct {
	Message string
	Type    string
}

// NewErrUnknownRecordType creates a new error of the ErrUnknownRecordType type
func NewErrUnknownRecordType(recordType string) ErrUnknownRecordType {
	_ = "STUB: not implemented"
	return *new(ErrUnknownRecordType)
}

func (e ErrUnknownRecordType) Error() string {
	_ = "STUB: not implemented"

	// ErrFileUnknownSEC is the error given when a record does not have a known type
	return ""
}

type ErrFileUnknownSEC struct {
	Message string
	SEC     string
}

// NewErrFileUnknownSEC creates a new error of the ErrFileUnknownSEC type
func NewErrFileUnknownSEC(secType string) ErrFileUnknownSEC {
	_ = "STUB: not implemented"
	return *new(ErrFileUnknownSEC)
}

func (e ErrFileUnknownSEC) Error() string {
	_ = "STUB: not implemented"

	// ErrFileCalculatedControlEquality is the error given when the control record does not match the calculated value
	return ""
}

type ErrFileCalculatedControlEquality struct {
	Message         string
	Field           string
	CalculatedValue int
	ControlValue    int
}

// NewErrFileCalculatedControlEquality creates a new error of the ErrFileCalculatedControlEquality type
func NewErrFileCalculatedControlEquality(field string, calculated, control int) ErrFileCalculatedControlEquality {
	_ = "STUB: not implemented"
	return *new(ErrFileCalculatedControlEquality)
}

func (e ErrFileCalculatedControlEquality) Error() string {
	_ = "STUB: not implemented"

	// ErrFileBatchNumberAscending is the error given when the batch numbers in a file are not in ascending order
	return ""
}

type ErrFileBatchNumberAscending struct {
	Message       string
	PreviousBatch int
	CurrentBatch  int
}

// NewErrFileBatchNumberAscending creates a new error of the ErrFileBatchNumberAscending type
func NewErrFileBatchNumberAscending(previous, current int) ErrFileBatchNumberAscending {
	_ = "STUB: not implemented"
	return *new(ErrFileBatchNumberAscending)
}

func (e ErrFileBatchNumberAscending) Error() string { _ = "STUB: not implemented"; return "" }
