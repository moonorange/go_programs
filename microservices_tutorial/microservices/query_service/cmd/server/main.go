package main

import (
	"context"
	"net/http"

	connect "connectrpc.com/connect"
	"github.com/protogo/gen"
	"github.com/protogo/gen/genconnect"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

func main() {
	const host = "localhost:8082"

	mux := http.NewServeMux()
	path, handler := genconnect.NewTaskServiceHandler(&taskServer{})
	mux.Handle(path, handler)
	logrus.Println("... Listening on", host)

	eg := errgroup.Group{}
	// Start the gRPC server
	eg.Go(func() error { return http.ListenAndServe(host, h2c.NewHandler(mux, &http2.Server{})) })
	logrus.Printf("Query service is running on host %s", host)

	err := eg.Wait()
	if err != nil {
		logrus.Fatal("failed to serve: ", err)
	}
}

type taskServer struct {
	genconnect.UnimplementedTaskServiceHandler
}

func (t *taskServer) ListTasksByTag(ctx context.Context, req *connect.Request[gen.ListTasksByTagRequest]) (*connect.Response[gen.ListTasksByTagResponse], error) {
	return &connect.Response[gen.ListTasksByTagResponse]{}, nil
}
