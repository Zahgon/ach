// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"

	"github.com/moov-io/ach"

	"github.com/juju/ansiterm"
)

func diffFiles(paths []string, validateOpts *ach.ValidateOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func readTwoFiles(paths []string, validateOpts *ach.ValidateOpts) (*ach.File, *ach.File, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// TODO(adam): we should lookup batches which are in f1 against those in f2 and if a similar
// batch is found (TODO: what, if any fields can we do exact matches on?) print that in the
// order of f1. Otherwise show missing batches (from f1's view) and additional batches (from
// f2's view) at the end.

func printDiffedFileHeader(w *ansiterm.TabWriter, f1, f2 *ach.File) {
	_ = "STUB: not implemented"
	return
}

func printColumn(minusBuf, plusBuf *bytes.Buffer, v1, v2 string) { _ = "STUB: not implemented"; return }
