package mask

import (
	"github.com/moov-io/ach"
)

func Number(s string) string { _ = "STUB: not implemented"; return "" }

// too short, we can't show anything

// Since we want the right-most digits unmasked start from the end of our string

// If the char to our right is masked then mask this left-aligned space as well.

func Name(s string) string { _ = "STUB: not implemented"; return "" }

type Options struct {
	MaskNames          bool
	MaskAccountNumbers bool
	MaskCorrectedData  bool
	MaskIdentification bool
}

func File(file *ach.File, options Options) *ach.File { _ = "STUB: not implemented"; return nil }

// Mask some addenda records
