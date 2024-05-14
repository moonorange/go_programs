package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"golang.org/x/sys/unix"
)

const (
	KAFKA_SERVER = "localhost:9092"
)

var wg = sync.WaitGroup{}

var GlobalCounter = NewCounter()

func main() {
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())

	go Publish(ctx, "test1")
	go Publish(ctx, "test2")

	go Consume(ctx, "test1", &Printer{})
	go Consume(ctx, "test2", &Printer{})

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, unix.SIGTERM, unix.SIGINT)

	<-sig
	fmt.Println()
	cancel()

	wg.Wait()
	fmt.Println("Finished")
}
