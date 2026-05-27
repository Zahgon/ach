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

//go:build examples

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/moov-io/ach"
)

type JsonParseEvent struct {
	Json ach.File `json:"data"`
}

func main() {
	lambda.Start(HandleRequest)
}

// logic to be executed when lambda starts goes here
func HandleRequest(ctx context.Context, event JsonParseEvent) (string, error) {
	_ = "STUB: not implemented"

	// get file from lambda event, it has already been marshaled from json to ach.File by Go
	return "", nil
}

// set file ID

// validate parsed file

// create buffer to contain NACHA text

// write ach.File to buffer

// get NACHA text from buffer
