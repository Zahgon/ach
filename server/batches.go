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
	"context"
	"net/http"

	"github.com/moov-io/ach"
	"github.com/moov-io/base/log"

	"github.com/go-kit/kit/endpoint"
)

type createBatchRequest struct {
	FileID string
	Batch  ach.Batcher

	requestID string
}

type createBatchResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func (r createBatchResponse) error() error { _ = "STUB: not implemented"; return nil }

func createBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

func decodeCreateBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In order to use FileFromJSON we need a populated JSON structure that can be parsed.
// We're going to copy the body into this shim to parse the Batch, otherwise we'd have
// to copy/export the logic of reading batches from their JSON representation.

type getBatchesRequest struct {
	fileID string

	requestID string
}

type getBatchesResponse struct {
	// TODO(adam): change this to JSON encode without wrapper {"batches": [..]}
	// We don't wrap json objects in other responses, so why here?
	Batches []ach.Batcher `json:"batches"`
	Err     error         `json:"error"`
}

func (r getBatchesResponse) count() int { _ = "STUB: not implemented"; return 0 }

func (r getBatchesResponse) error() error { _ = "STUB: not implemented"; return nil }

func decodeGetBatchesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBatchesEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

type getBatchRequest struct {
	fileID  string
	batchID string

	requestID string
}

type getBatchResponse struct {
	Batch ach.Batcher `json:"batch"`
	Err   error       `json:"error"`
}

func (r getBatchResponse) error() error { _ = "STUB: not implemented"; return nil }

func decodeGetBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

type deleteBatchRequest struct {
	fileID  string
	batchID string

	requestID string
}

type deleteBatchResponse struct {
	Err error `json:"error"`
}

func (r deleteBatchResponse) error() error { _ = "STUB: not implemented"; return nil }

func decodeDeleteBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}
