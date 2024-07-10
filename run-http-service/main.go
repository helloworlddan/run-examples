package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/helloworlddan/run"
)

func main() {
	http.HandleFunc("/", indexHandler)

	// Store config
	run.PutConfig("some-key", "some-val")

	// Store client with lazy initialization
	var bqClient *bigquery.Client
	lazyInit := func() {
		run.Debug(nil, "lazy init: bigquery")
		var err error
		ctx := context.Background()
		bqClient, err = bigquery.NewClient(ctx, run.ProjectID())
		if err != nil {
			run.Error(nil, err)
		}
	}
	run.AddLazyClient("nibquery", bqClient, lazyInit)

	// Define shutdown behavior
	shutdown := func(ctx context.Context) {
		run.Debug(nil, "shutting down connections...")
		time.Sleep(time.Second * 1) // Pretending to clean up
		run.Debug(nil, "connections closed")
	}
	err := run.ServeHTTP(shutdown, nil)
	if err != nil {
		run.Fatal(nil, err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name: %s\n", run.ServiceName())
	fmt.Fprintf(w, "Revision: %s\n", run.ServiceRevision())
	fmt.Fprintf(w, "ProjectID: %s\n", run.ProjectID())
	// Access config
	cfg, err := run.GetConfig("some-key")
	if err != nil {
		run.Error(r, err)
	}

	// Access client
	clientRef, err := run.GetClient("bigquery")
	if err != nil {
		run.Error(nil, err)
	}
	bqClient := clientRef.(*bigquery.Client)
	// NOTE: use client
	_ = bqClient

	fmt.Fprintf(w, "Config[some-key]: %s\n", cfg)
	run.Debugf(r, "request completed")
}
