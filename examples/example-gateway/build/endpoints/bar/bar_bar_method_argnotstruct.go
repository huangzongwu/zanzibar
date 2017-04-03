// Code generated by zanzibar
// @generated

package bar

import (
	"context"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
)

// HandleArgNotStructRequest handles "/bar/arg-not-struct-path".
func HandleArgNotStructRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {

	// Handle request body.
	var body ArgNotStructHTTPRequest
	if ok := req.ReadAndUnmarshalBody(&body); !ok {
		return
	}
	clientRequest := convertToArgNotStructClientRequest(&body)
	clientResp, err := clients.Bar.ArgNotStruct(ctx, clientRequest)
	if err != nil {
		req.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		res.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	defer func() {
		if cerr := clientResp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Handle client respnse.
	expectedStatusCode := []int{200}
	if !res.IsOKResponse(clientResp.StatusCode, expectedStatusCode) {
		req.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	res.WriteJSONBytes(200, nil)
}

func convertToArgNotStructClientRequest(body *ArgNotStructHTTPRequest) *barClient.ArgNotStructHTTPRequest {
	clientRequest := &barClient.ArgNotStructHTTPRequest{}

	clientRequest.Request = string(body.Request)

	return clientRequest
}
