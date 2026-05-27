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
	// ErrBatchNoEntries is the error given when a batch doesn't have any entries
	ErrBatchNoEntries = errors.New("must have Entry Record(s) to be built")
	// ErrBatchADVCount is the error given when an ADV batch has too many entries
	ErrBatchADVCount = errors.New("there can be a maximum of 9999 ADV Sequence Numbers (ADV Entry Detail Records)")
	// ErrBatchAddendaIndicator is the error given when the addenda indicator is incorrectly set
	ErrBatchAddendaIndicator = errors.New("is 0 but found addenda record(s)")
	// ErrBatchOriginatorDNE is the error given when a non-government agency tries to originate a DNE
	ErrBatchOriginatorDNE = errors.New("only government agencies (originator status code 2) can originate a DNE")
	// ErrBatchInvalidCardTransactionType is the error given when a card transaction type is invalid
	ErrBatchInvalidCardTransactionType = errors.New("invalid card transaction type")
	// ErrBatchDebitOnly is the error given when a batch which can only have debits has a credit
	ErrBatchDebitOnly = errors.New("this batch type does not allow credit transaction codes")
	// ErrBatchCreditOnly is the error given when a batch which can only have credits has a debit
	ErrBatchCreditOnly = errors.New("this batch type does not allow debit transaction codes")
	// ErrBatchCheckSerialNumber is the error given when a batch requires check serial numbers, but it is missing
	ErrBatchCheckSerialNumber = errors.New("this batch type requires entries to have Check Serial Numbers")
	// ErrBatchSECType is the error given when the batch's header has the wrong SEC for its type
	ErrBatchSECType = errors.New("header SEC does not match this batch's type")
	// ErrBatchServiceClassCode is the error given when the batch's header has the wrong SCC for its type
	ErrBatchServiceClassCode = errors.New("header SCC is not valid for this batch's type")
	// ErrBatchTransactionCode is the error given when a batch has an invalid transaction code
	ErrBatchTransactionCode = errors.New("transaction code is not valid for this batch's type")
	// ErrBatchTransactionCodeAddenda is the error given when a batch has an addenda on a transaction code which doesn't allow it
	ErrBatchTransactionCodeAddenda = errors.New("this batch type does not allow an addenda for this transaction code")
	// ErrBatchAmountNonZero is the error given when an entry for a non-zero amount is in a batch that requires zero amount entries
	ErrBatchAmountNonZero = errors.New("this batch type requires that the amount is zero")
	// ErrBatchAmountZero is the error given when an entry for zero amount is in a batch that requires non-zero amount entries
	ErrBatchAmountZero = errors.New("this batch type requires that the amount is non-zero")
	// ErrBatchCompanyEntryDescriptionAutoenroll is the error given when the Company Entry Description is invalid (needs to be 'Autoenroll')
	ErrBatchCompanyEntryDescriptionAutoenroll = errors.New("this batch type requires that the Company Entry Description is AUTOENROLL")
	// ErrBatchCompanyEntryDescriptionREDEPCHECK is the error given when the Company Entry Description is invalid (needs to be 'REDEPCHECK')
	ErrBatchCompanyEntryDescriptionREDEPCHECK = errors.New("this batch type requires that the Company Entry Description is REDEPCHECK")
	// ErrBatchAddendaCategory is the error given when the addenda isn't allowed for the batch's type and category
	ErrBatchAddendaCategory = errors.New("this batch type does not allow this addenda for category")
)

// BatchError is an Error that describes batch validation issues
type BatchError struct {
	BatchNumber int
	BatchType   string
	FieldName   string
	FieldValue  interface{}
	Err         error
}

func (e *BatchError) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap implements the base.UnwrappableError interface for BatchError
func (e *BatchError) Unwrap() error {
	_ = "STUB: not implemented"

	// error returns a new BatchError based on err
	return nil
}

func (b *Batch) Error(field string, err error, values ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// only the first value counts

// error returns a new BatchError based on err
func (iatBatch *IATBatch) Error(field string, err error, values ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// only the first value counts

// ErrBatchHeaderControlEquality is the error given when the control record does not match the calculated value
type ErrBatchHeaderControlEquality struct {
	Message      string
	HeaderValue  interface{}
	ControlValue interface{}
}

// NewErrBatchHeaderControlEquality creates a new error of the ErrBatchHeaderControlEquality type
func NewErrBatchHeaderControlEquality(header, control interface{}) ErrBatchHeaderControlEquality {
	_ = "STUB: not implemented"
	return *new(ErrBatchHeaderControlEquality)
}

func (e ErrBatchHeaderControlEquality) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchCalculatedControlEquality is the error given when the control record does not match the calculated value
	return ""
}

type ErrBatchCalculatedControlEquality struct {
	Message         string
	CalculatedValue interface{}
	ControlValue    interface{}
}

// NewErrBatchCalculatedControlEquality creates a new error of the ErrBatchCalculatedControlEquality type
func NewErrBatchCalculatedControlEquality(calculated, control interface{}) ErrBatchCalculatedControlEquality {
	_ = "STUB: not implemented"
	return *new(ErrBatchCalculatedControlEquality)
}

func (e ErrBatchCalculatedControlEquality) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchAscending is the error given when the trace numbers in a batch are not in ascending order
	return ""
}

type ErrBatchAscending struct {
	Message       string
	PreviousTrace interface{}
	CurrentTrace  interface{}
}

// NewErrBatchAscending creates a new error of the ErrBatchAscending type
func NewErrBatchAscending(previous, current interface{}) ErrBatchAscending {
	_ = "STUB: not implemented"
	return *new(ErrBatchAscending)
}

func (e ErrBatchAscending) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchCategory is the error given when a batch has entires with two different categories
	return ""
}

type ErrBatchCategory struct {
	Message   string
	CategoryA string
	CategoryB string
}

// NewErrBatchCategory creates a new error of the ErrBatchCategory type
func NewErrBatchCategory(categoryA, categoryB string) ErrBatchCategory {
	_ = "STUB: not implemented"
	return *new(ErrBatchCategory)
}

func (e ErrBatchCategory) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchTraceNumberNotODFI is the error given when a batch's ODFI does not match an entry's trace number
	return ""
}

type ErrBatchTraceNumberNotODFI struct {
	Message     string
	ODFI        string
	TraceNumber string
}

// NewErrBatchTraceNumberNotODFI creates a new error of the ErrBatchTraceNumberNotODFI type
func NewErrBatchTraceNumberNotODFI(odfi, trace string) ErrBatchTraceNumberNotODFI {
	_ = "STUB: not implemented"
	return *new(ErrBatchTraceNumberNotODFI)
}

func (e ErrBatchTraceNumberNotODFI) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchAddendaTraceNumber is the error given when the entry detail sequence number doesn't match the trace number
	return ""
}

type ErrBatchAddendaTraceNumber struct {
	Message           string
	EntryDetailNumber string
	TraceNumber       string
}

// NewErrBatchAddendaTraceNumber creates a new error of the ErrBatchAddendaTraceNumber type
func NewErrBatchAddendaTraceNumber(entryDetail, trace string) ErrBatchAddendaTraceNumber {
	_ = "STUB: not implemented"
	return *new(ErrBatchAddendaTraceNumber)
}

func (e ErrBatchAddendaTraceNumber) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchAddendaCount is the error given when there are too many addenda than allowed for the batch type
	return ""
}

type ErrBatchAddendaCount struct {
	Message      string
	FoundCount   int
	AllowedCount int
}

// NewErrBatchAddendaCount creates a new error of the ErrBatchAddendaCount type
func NewErrBatchAddendaCount(found, allowed int) ErrBatchAddendaCount {
	_ = "STUB: not implemented"
	return *new(ErrBatchAddendaCount)
}

func (e ErrBatchAddendaCount) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchRequiredAddendaCount is the error given when the batch type requires a certain number of addenda, which is not met
	return ""
}

type ErrBatchRequiredAddendaCount struct {
	Message       string
	FoundCount    int
	RequiredCount int
}

// NewErrBatchRequiredAddendaCount creates a new error of the ErrBatchRequiredAddendaCount type
func NewErrBatchRequiredAddendaCount(found, required int) ErrBatchRequiredAddendaCount {
	_ = "STUB: not implemented"
	return *new(ErrBatchRequiredAddendaCount)
}

func (e ErrBatchRequiredAddendaCount) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchExpectedAddendaCount is the error given when the batch type has entries with a field
	// for the number of addenda, and a different number of addenda are foound
	return ""
}

type ErrBatchExpectedAddendaCount struct {
	Message       string
	FoundCount    int
	ExpectedCount int
}

// NewErrBatchExpectedAddendaCount creates a new error of the ErrBatchExpectedAddendaCount type
func NewErrBatchExpectedAddendaCount(found, expected int) ErrBatchExpectedAddendaCount {
	_ = "STUB: not implemented"
	return *new(ErrBatchExpectedAddendaCount)
}

func (e ErrBatchExpectedAddendaCount) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchServiceClassTranCode is the error given when the transaction code is not valid for the batch's service class
	return ""
}

type ErrBatchServiceClassTranCode struct {
	Message          string
	ServiceClassCode int
	TransactionCode  int
}

// NewErrBatchServiceClassTranCode creates a new error of the ErrBatchServiceClassTranCode type
func NewErrBatchServiceClassTranCode(serviceClassCode, transactionCode int) ErrBatchServiceClassTranCode {
	_ = "STUB: not implemented"
	return *new(ErrBatchServiceClassTranCode)
}

func (e ErrBatchServiceClassTranCode) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchAmount is the error given when the amount exceeds the batch type's limit
	return ""
}

type ErrBatchAmount struct {
	Message string
	Amount  int
	Limit   int
}

// NewErrBatchAmount creates a new error of the ErrBatchAmount type
func NewErrBatchAmount(amount, limit int) ErrBatchAmount {
	_ = "STUB: not implemented"
	// TODO: pretty format the amounts to make it more readable
	return *new(ErrBatchAmount)
}

func (e ErrBatchAmount) Error() string {
	_ = "STUB: not implemented"

	// ErrBatchIATNOC is the error given when an IAT batch has an NOC, and there are invalid values
	return ""
}

type ErrBatchIATNOC struct {
	Message  string
	Found    interface{}
	Expected interface{}
}

// NewErrBatchIATNOC creates a new error of the ErrBatchIATNOC type
func NewErrBatchIATNOC(found, expected interface{}) ErrBatchIATNOC {
	_ = "STUB: not implemented"
	// TODO: pretty format the amounts to make it more readable
	return *new(ErrBatchIATNOC)
}

func (e ErrBatchIATNOC) Error() string { _ = "STUB: not implemented"; return "" }
