package main

import (
	"context"
	"time"

	"github.com/helloworlddan/run"
	"github.com/helloworlddan/run-examples/run-grpc-service/runclock"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	runclock.RegisterRunClockServer(server, clockServer{})

	err := run.ServeGRPC(func(ctx context.Context) {
		run.Debug(nil, "shutting down connections...")
		time.Sleep(time.Second * 1) // Pretending to clean up
		run.Debug(nil, "connections closed")
	}, server)
	if err != nil {
		run.Fatal(nil, err)
	}
}

type clockServer struct {
	runclock.UnimplementedRunClockServer
}

func (srv clockServer) GetTime(ctx context.Context, in *runclock.Empty) (*runclock.Time, error) {
	now := time.Now()
	run.Debug(nil, "received request")
	return &runclock.Time{
		Formatted: now.GoString(),
	}, nil
}
