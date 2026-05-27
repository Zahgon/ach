package main

import (
	"syscall/js"

	"github.com/moov-io/ach"
)

func parseACH(input string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseReadable(file *ach.File) (string, error) { _ = "STUB: not implemented"; return "", nil }

func prettyJson(input string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func prettyPrintJSON() js.Func { _ = "STUB: not implemented"; return *new(js.Func) }

func printACH() js.Func { _ = "STUB: not implemented"; return *new(js.Func) }

func printReadable() js.Func { _ = "STUB: not implemented"; return *new(js.Func) }

// Parses input, either JSON or Nacha format to an ach.File
func parseFile(input string) (*ach.File, error) { _ = "STUB: not implemented"; return nil, nil }

func reverseFile() js.Func { _ = "STUB: not implemented"; return *new(js.Func) }

func writeVersion() { _ = "STUB: not implemented"; return }

func main() {
	js.Global().Set("parseACH", prettyPrintJSON())
	js.Global().Set("parseJSON", printACH())
	js.Global().Set("parseReadable", printReadable())
	js.Global().Set("reverseFile", reverseFile())

	writeVersion()

	<-make(chan bool)
}
