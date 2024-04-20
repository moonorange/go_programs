package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

// Task represents a task item
type Task struct {
	ID   int32    `json:"id"`
	Text string   `json:"text"`
	Tags []string `json:"tags"`
}

func (t *taskServer) ListTasksByTag(ctx context.Context, req *connect.Request[gen.ListTasksByTagRequest]) (*connect.Response[gen.ListTasksByTagResponse], error) {
	// Define the path to the JSON file containing tasks
	// Use a json file for simplicity, but in a real-world scenario this would come from a database
	filePath := "../../../file_storage/task.json"

	// Read the content of the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Check if the file is empty
	if len(data) == 0 {
		return connect.NewResponse(&gen.ListTasksByTagResponse{Tasks: nil}), nil
	}

	// Unmarshal the JSON data into a slice of Task structs
	var tasks []*gen.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Filter tasks by the specified tag
	var filteredTasks []*gen.Task
	for _, task := range tasks {
		for _, t := range task.Tags {
			if t == req.Msg.TagName {
				filteredTasks = append(filteredTasks, task)
				break
			}
		}
	}
	return connect.NewResponse(&gen.ListTasksByTagResponse{Tasks: filteredTasks}), nil
}
