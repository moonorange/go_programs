package client

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/protogo/gen/genconnect"
)

func NewTaskServiceClient(address string) genconnect.TaskServiceClient {
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	client := genconnect.NewTaskServiceClient(
		http.DefaultClient,
		address,
		connect.WithGRPC(),
	)

	return client
}
