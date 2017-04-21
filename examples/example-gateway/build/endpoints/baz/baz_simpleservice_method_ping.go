// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package baz

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"
)

// PingHandler is the handler for "/baz/ping"
type PingHandler struct {
	Clients *clients.Clients
}

// NewPingEndpoint creates a handler
func NewPingEndpoint(
	gateway *zanzibar.Gateway,
) *PingHandler {
	return &PingHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/baz/ping".
func (handler *PingHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	workflow := PingEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSON(200, cliRespHeaders, response)
}

// PingEndpoint calls thrift client Baz.Ping
type PingEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w PingEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (*endpointsBazBaz.BazResponse, zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Baz.Ping(
		ctx, clientHeaders,
	)

	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		// TODO(sindelar): Consider returning partial headers in error case.
		return nil, nil, err
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertPingClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertPingClientResponse(body *clientsBazBaz.BazResponse) *endpointsBazBaz.BazResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := (*endpointsBazBaz.BazResponse)(body)
	return downstreamResponse
}
