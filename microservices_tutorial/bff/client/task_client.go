package client

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/protogo/gen/genconnect"
)

func NewTaskServiceClient(baseURL string) genconnect.TaskServiceClient {
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	client := genconnect.NewTaskServiceClient(
		http.DefaultClient,
		baseURL,
		connect.WithGRPC(),
	)

	return client
}
