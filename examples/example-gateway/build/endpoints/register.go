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

package endpoints

import (
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow"
	tchannelBaz "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz"
	"github.com/uber/zanzibar/examples/example-gateway/middlewares/example"
	"github.com/uber/zanzibar/runtime/middlewares/logger"

	"github.com/uber/zanzibar/runtime"
)

// Endpoints is a struct that holds all the endpoints
type Endpoints struct {
	BarArgNotStructHTTPHandler             *bar.ArgNotStructHandler
	BarArgWithHeadersHTTPHandler           *bar.ArgWithHeadersHandler
	BarArgWithManyQueryParamsHTTPHandler   *bar.ArgWithManyQueryParamsHandler
	BarArgWithNestedQueryParamsHTTPHandler *bar.ArgWithNestedQueryParamsHandler
	BarArgWithQueryHeaderHTTPHandler       *bar.ArgWithQueryHeaderHandler
	BarArgWithQueryParamsHTTPHandler       *bar.ArgWithQueryParamsHandler
	BarMissingArgHTTPHandler               *bar.MissingArgHandler
	BarNoRequestHTTPHandler                *bar.NoRequestHandler
	BarNormalHTTPHandler                   *bar.NormalHandler
	BarTooManyArgsHTTPHandler              *bar.TooManyArgsHandler
	BazCallHTTPHandler                     *baz.CallHandler
	BazCompareHTTPHandler                  *baz.CompareHandler
	BazPingHTTPHandler                     *baz.PingHandler
	BazSillyNoopHTTPHandler                *baz.SillyNoopHandler
	ContactsSaveContactsHTTPHandler        *contacts.SaveContactsHandler
	GooglenowAddCredentialsHTTPHandler     *googlenow.AddCredentialsHandler
	GooglenowCheckCredentialsHTTPHandler   *googlenow.CheckCredentialsHandler
	BazTChannelCallTChannelHandler         zanzibar.TChannelHandler
}

// CreateEndpoints bootstraps the endpoints.
func CreateEndpoints(
	gateway *zanzibar.Gateway,
) interface{} {
	return &Endpoints{
		BarArgNotStructHTTPHandler:             bar.NewArgNotStructEndpoint(gateway),
		BarArgWithHeadersHTTPHandler:           bar.NewArgWithHeadersEndpoint(gateway),
		BarArgWithManyQueryParamsHTTPHandler:   bar.NewArgWithManyQueryParamsEndpoint(gateway),
		BarArgWithNestedQueryParamsHTTPHandler: bar.NewArgWithNestedQueryParamsEndpoint(gateway),
		BarArgWithQueryHeaderHTTPHandler:       bar.NewArgWithQueryHeaderEndpoint(gateway),
		BarArgWithQueryParamsHTTPHandler:       bar.NewArgWithQueryParamsEndpoint(gateway),
		BarMissingArgHTTPHandler:               bar.NewMissingArgEndpoint(gateway),
		BarNoRequestHTTPHandler:                bar.NewNoRequestEndpoint(gateway),
		BarNormalHTTPHandler:                   bar.NewNormalEndpoint(gateway),
		BarTooManyArgsHTTPHandler:              bar.NewTooManyArgsEndpoint(gateway),
		BazCallHTTPHandler:                     baz.NewCallEndpoint(gateway),
		BazCompareHTTPHandler:                  baz.NewCompareEndpoint(gateway),
		BazPingHTTPHandler:                     baz.NewPingEndpoint(gateway),
		BazSillyNoopHTTPHandler:                baz.NewSillyNoopEndpoint(gateway),
		ContactsSaveContactsHTTPHandler:        contacts.NewSaveContactsEndpoint(gateway),
		GooglenowAddCredentialsHTTPHandler:     googlenow.NewAddCredentialsEndpoint(gateway),
		GooglenowCheckCredentialsHTTPHandler:   googlenow.NewCheckCredentialsEndpoint(gateway),
		BazTChannelCallTChannelHandler:         tchannelBaz.NewSimpleServiceCallHandler(gateway),
	}
}

// Register will register all endpoints
func Register(g *zanzibar.Gateway) {
	endpoints := CreateEndpoints(g).(*Endpoints)

	g.HTTPRouter.Register(
		"POST", "/bar/arg-not-struct-path",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argNotStruct",
			endpoints.BarArgNotStructHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/bar/argWithHeaders",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithHeaders",
			endpoints.BarArgWithHeadersHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/bar/argWithManyQueryParams",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithManyQueryParams",
			endpoints.BarArgWithManyQueryParamsHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/bar/argWithNestedQueryParams",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithNestedQueryParams",
			endpoints.BarArgWithNestedQueryParamsHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/bar/argWithQueryHeader",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithQueryHeader",
			endpoints.BarArgWithQueryHeaderHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/bar/argWithQueryParams",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithQueryParams",
			endpoints.BarArgWithQueryParamsHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/bar/missing-arg-path",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"missingArg",
			endpoints.BarMissingArgHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/bar/no-request-path",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"noRequest",
			endpoints.BarNoRequestHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/bar/bar-path",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"normal",
			zanzibar.NewStack([]zanzibar.MiddlewareHandle{
				example.NewMiddleWare(
					g,
					example.Options{
						Foo: "test",
					},
				),
				logger.NewMiddleWare(
					g,
					logger.Options{},
				),
			}, endpoints.BarNormalHTTPHandler.HandleRequest).Handle,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/bar/too-many-args-path",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"tooManyArgs",
			endpoints.BarTooManyArgsHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/baz/call",
		zanzibar.NewRouterEndpoint(
			g,
			"baz",
			"call",
			endpoints.BazCallHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/baz/compare",
		zanzibar.NewRouterEndpoint(
			g,
			"baz",
			"compare",
			endpoints.BazCompareHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/baz/ping",
		zanzibar.NewRouterEndpoint(
			g,
			"baz",
			"ping",
			endpoints.BazPingHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"GET", "/baz/silly-noop",
		zanzibar.NewRouterEndpoint(
			g,
			"baz",
			"sillyNoop",
			endpoints.BazSillyNoopHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/contacts/:userUUID/contacts",
		zanzibar.NewRouterEndpoint(
			g,
			"contacts",
			"saveContacts",
			endpoints.ContactsSaveContactsHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/googlenow/add-credentials",
		zanzibar.NewRouterEndpoint(
			g,
			"googlenow",
			"addCredentials",
			endpoints.GooglenowAddCredentialsHTTPHandler.HandleRequest,
		),
	)
	g.HTTPRouter.Register(
		"POST", "/googlenow/check-credentials",
		zanzibar.NewRouterEndpoint(
			g,
			"googlenow",
			"checkCredentials",
			endpoints.GooglenowCheckCredentialsHTTPHandler.HandleRequest,
		),
	)
	g.TChannelRouter.Register("SimpleService", "Call", endpoints.BazTChannelCallTChannelHandler)
}
