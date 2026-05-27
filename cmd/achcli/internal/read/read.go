package read

import (
	"github.com/moov-io/ach"
)

type Format string

var (
	FormatUnknown Format = "unknown"
	FormatNacha   Format = "nacha"
	FormatJSON    Format = "json"
)

func Filepath(path string, validateOptsPath *string, skipAll *bool) (*ach.File, Format, error) {
	_ = "STUB: not implemented"
	return nil, *new(Format), nil
}

func readValidationOpts(path *string, skipAll *bool) (*ach.ValidateOpts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read config file

func readFile(path string, validateOpts *ach.ValidateOpts) (*ach.File, Format, error) {
	_ = "STUB: not implemented"
	return nil, *new(Format), nil
}

func readACHFile(input []byte, validateOpts *ach.ValidateOpts) (*ach.File, Format, error) {
	_ = "STUB: not implemented"
	return nil, *new(Format), nil
}

func readJsonFile(input []byte, validateOpts *ach.ValidateOpts) (*ach.File, Format, error) {
	_ = "STUB: not implemented"
	return nil, *new(Format), nil
}
