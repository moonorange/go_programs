package main

import (
	"net/http"

	handler "command_service/handler"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	connect "github.com/protogo/gen/genconnect"
)

func main() {
	const host = "localhost:8081"

	mux := http.NewServeMux()
	path, handler := connect.NewTaskServiceHandler(handler.TaskHandler{})
	mux.Handle(path, handler)
	logrus.Println("... Listening on", host)

	eg := errgroup.Group{}
	// Start the gRPC server
	eg.Go(func() error { return http.ListenAndServe(host, h2c.NewHandler(mux, &http2.Server{})) })
	logrus.Printf("Command service is running on host %s", host)

	err := eg.Wait()
	if err != nil {
		logrus.Fatal("failed to serve: ", err)
	}
}
