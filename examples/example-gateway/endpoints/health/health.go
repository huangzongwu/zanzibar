package health

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/clients"
	"github.com/uber/zanzibar/runtime"
)

// HandleHealthRequest for the health request
func HandleHealthRequest(
	ctx context.Context,
	inc *zanzibar.IncomingMessage,
	g *zanzibar.Gateway,
	clients *clients.Clients,
) {
	resp := &JSONResponse{
		Ok:      true,
		Message: "Healthy, from example-gateway",
	}

	inc.WriteJSON(200, resp)
}
