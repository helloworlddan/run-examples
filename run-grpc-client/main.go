package main

import (
	"context"
	"log"

	"github.com/helloworlddan/run-examples/run-grpc-client/runclock"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot connect to server")
	}
	defer conn.Close()

	client := runclock.NewRunClockClient(conn)

	ctx := context.Background()

	for range 10 {
		resp, err := client.GetTime(ctx, &runclock.Empty{})
		if err != nil {
			log.Fatal("cannot call endpoint")
		}

		log.Printf("receivefd time: %s", resp.GetFormatted())
	}
}
