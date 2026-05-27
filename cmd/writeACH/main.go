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

package main

import (
	"flag"
	"path/filepath"
	"time"
)

var (
	fPath      = flag.String("fPath", "", "File Path")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	// output formats
	flagJson = flag.Bool("json", false, "Output file in json")
)

// main creates an ACH File with 4 batches of SEC Code PPD.
// Each batch contains an EntryAddendaCount of 2500.
func main() {
	flag.Parse()

	filename := time.Now().UTC().Format("200601021504")
	if *flagJson {
		filename += ".json"
	} else {
		filename += ".ach"
	}

	path := filepath.Join(*fPath, filename)
	write(path)
}

func write(path string) { _ = "STUB: not implemented"; return }

// To create a file

// Create 4 Batches of SEC Code PPD

// Create Entry

// Add addenda record for an entry

// Add entries

// Create the batch.

// Add batch to the file

// ensure we have a validated file structure

// Create the file

// Write to a file

// Write in JSON format

// Write in ACH plain text format
