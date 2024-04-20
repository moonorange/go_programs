package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	gen "github.com/protogo/gen"
)

func main() {
	// Create a listener on TCP port
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatal("failed to listen: ", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Listen for OS signals to stop the server when receiving a cancel signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the gRPC server
	go func() {
		if err := s.Serve(listener); err != nil {
			logrus.Fatal("failed to serve: ", err)
		}
		logrus.Printf("Command service is running on port %d", port)
	}()

	// Wait for a signal to stop the server
	sig := <-sigChan
	logrus.Printf("Command service is shutting down: %s", sig)

	// Stop the gRPC server gracefully
	s.GracefulStop()
}
