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
	"bufio"
	"io"
	"strings"
)

// Writer writes a File to an io.Writer.
// The File is validated against Nacha guidelines unless BypassValidation is enabled.
type Writer struct {
	w *bufio.Writer

	lineNum    int    //current line being written
	LineEnding string // configurable line ending to support different consumer requirements
	// BypassValidation can be set to skip file validation and will allow non-compliant Nacha files to be written.
	BypassValidation bool
}

// WriteOpts defines options for writing a file.
type WriteOpts struct {
	// LineEnding sets a custom line ending character.
	LineEnding string `json:"lineEnding"`
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer { _ = "STUB: not implemented"; return nil }

// NewWriter returns a new Writer that writes to w.
func NewWriterWithOpts(w io.Writer, opts *WriteOpts) *Writer { _ = "STUB: not implemented"; return nil }

var (
	paddingLine = strings.Repeat("9", 94)
)

// Writer writes a single ach.file record to w
func (w *Writer) Write(file *File) error { _ = "STUB: not implemented"; return nil }

// Iterate over all records in the file

// pad the final block

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) writeBatch(file *File, isADV bool) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) writeIATBatch(file *File) error { _ = "STUB: not implemented"; return nil }

// IAT Addenda17

// IAT Addenda18

type writeEntry interface {
	String() string
}

func (w *Writer) writeLine(entry writeEntry) error { _ = "STUB: not implemented"; return nil }

// Avoid allocations by flushing the buffer
