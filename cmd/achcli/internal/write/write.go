package write

import (
	"io"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/internal/read"
)

func File(w io.Writer, file *ach.File, format read.Format) error {
	_ = "STUB: not implemented"
	return nil
}
