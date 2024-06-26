package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/helloworlddan/run"
)

func main() {
	http.HandleFunc("/", indexHandler)

	run.PutConfig("some-key", "some-val")

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
	cfg, err := run.GetConfig("some-key")
	if err != nil {
		run.Error(r, err)
	}
	fmt.Fprintf(w, "Config[some-key]: %s\n", cfg)
	run.Debugf(r, "request completed")
}
