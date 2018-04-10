// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package barendpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/workflow"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarNoRequestHandler is the handler for "/bar/no-request-path"
type BarNoRequestHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarNoRequestHandler creates a handler
func NewBarNoRequestHandler(deps *module.Dependencies) *BarNoRequestHandler {
	handler := &BarNoRequestHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "noRequest",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarNoRequestHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/bar/no-request-path",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/no-request-path".
func (h *BarNoRequestHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	// log endpoint request to downstream services
	if ce := req.Logger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		req.Logger.Debug("endpoint request to downstream", zfields...)
	}

	w := workflow.NewBarNoRequestWorkflow(h.Clients, req.Logger)

	response, cliRespHeaders, err := w.Handle(ctx, req.Header)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			res.SendError(500, "Unexpected server error", err)
			return
		}

	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}
