// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package describe

import (
	"io"
	"text/tabwriter"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe/mask"
)

type Opts struct {
	mask.Options

	PrettyAmounts bool
}

func File(ww io.Writer, file *ach.File, opts *Opts) { _ = "STUB: not implemented"; return }

// Mask the file

// FileHeader

// Batches

// IATBatches

// FileControl

// formatAmount can optionally convert an integer into a human readable amount
func formatAmount(prettyAmounts bool, amt int) string { _ = "STUB: not implemented"; return "" }

func dumpAddenda02(w *tabwriter.Writer, a *ach.Addenda02) { _ = "STUB: not implemented"; return }

func dumpAddenda99Dishonored(w *tabwriter.Writer, a *ach.Addenda99Dishonored) {
	_ = "STUB: not implemented"
	return
}

func dumpAddenda99Contested(w *tabwriter.Writer, a *ach.Addenda99Contested) {
	_ = "STUB: not implemented"
	return
}

func dumpAddenda05(w *tabwriter.Writer, batch ach.Batcher, a *ach.Addenda05, opts *Opts) {
	_ = "STUB: not implemented"
	return
}

func dumpAddenda98(w *tabwriter.Writer, opts *Opts, a *ach.Addenda98) {
	_ = "STUB: not implemented"
	return
}

func dumpAddenda99(w *tabwriter.Writer, a *ach.Addenda99) { _ = "STUB: not implemented"; return }

func dumpAddenda10(w *tabwriter.Writer, a *ach.Addenda10) { _ = "STUB: not implemented"; return }

func dumpAddenda11(w *tabwriter.Writer, a *ach.Addenda11) { _ = "STUB: not implemented"; return }

func dumpAddenda12(w *tabwriter.Writer, a *ach.Addenda12) { _ = "STUB: not implemented"; return }

func dumpAddenda13(w *tabwriter.Writer, a *ach.Addenda13) { _ = "STUB: not implemented"; return }

func dumpAddenda14(w *tabwriter.Writer, a *ach.Addenda14) { _ = "STUB: not implemented"; return }

func dumpAddenda15(w *tabwriter.Writer, a *ach.Addenda15) { _ = "STUB: not implemented"; return }

func dumpAddenda16(w *tabwriter.Writer, a *ach.Addenda16) { _ = "STUB: not implemented"; return }

func dumpAddenda17(w *tabwriter.Writer, a *ach.Addenda17) { _ = "STUB: not implemented"; return }

func dumpAddenda18(w *tabwriter.Writer, a *ach.Addenda18) { _ = "STUB: not implemented"; return }
