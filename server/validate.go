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

package server

import (
	"io"
	"net/http"

	"github.com/moov-io/ach"
)

const (
	skipAll                          = "skipAll"
	requireABAOrigin                 = "requireABAOrigin"
	bypassOrigin                     = "bypassOrigin"
	bypassOriginValidation           = "bypassOriginValidation"
	bypassDestination                = "bypassDestination"
	bypassDestinationValidation      = "bypassDestinationValidation"
	customTraceNumbers               = "customTraceNumbers"
	allowZeroBatches                 = "allowZeroBatches"
	allowMissingFileHeader           = "allowMissingFileHeader"
	allowMissingFileControl          = "allowMissingFileControl"
	bypassCompanyIdentificationMatch = "bypassCompanyIdentificationMatch"
	customReturnCodes                = "customReturnCodes"
	unequalServiceClassCode          = "unequalServiceClassCode"
	unorderedBatchNumbers            = "unorderedBatchNumbers"
	allowUnorderedBatchNumbers       = "allowUnorderedBatchNumbers"
	allowInvalidCheckDigit           = "allowInvalidCheckDigit"
	unequalAddendaCounts             = "unequalAddendaCounts"
	preserveSpaces                   = "preserveSpaces"
	allowInvalidAmounts              = "allowInvalidAmounts"
	allowZeroEntryAmount             = "allowZeroEntryAmount"
	allowSpecialCharacters           = "allowSpecialCharacters"
	allowEmptyIndividualName         = "allowEmptyIndividualName"
	bypassBatchValidation            = "bypassBatchValidation"
	skipFileCreationValidation       = "skipFileCreationValidation"
	skipBatchHeaderCompanyValidation = "skipBatchHeaderCompanyValidation"
)

// readValidateOpts parses ValidateOpts from the URL query parameters and from the request body.
// A copy of the request body is returned. Callers are responsible for closing the body.
//
// Query parameters override the JSON body
func readValidateOpts(request *http.Request) (io.ReadCloser, *ach.ValidateOpts, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil, nil
}
