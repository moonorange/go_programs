package handler

import (
	connect "github.com/protogo/gen/genconnect"
)

// petStoreServiceServer implements the PetStoreService API.
type TaskHandler struct {
	connect.UnimplementedTaskServiceHandler
}
