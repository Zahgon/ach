package fix

import (
	"github.com/moov-io/ach"
)

func Perform(path string, validateOptsPath *string, skipAll *bool, conf Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Build up our fixers

// Fix the file

// Batch headers

// Write file

type batchHeaderFixer func(bh *ach.BatchHeader) error
