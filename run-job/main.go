package main

import (
	"context"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/helloworlddan/run"
)

func main() {
	// Store config
	run.PutConfig("my.app.key", "some value")
	cfgVal, err := run.GetConfig("my.app.key")
	if err != nil {
		run.Debugf(nil, "unable to read config: %v", err)
	}
	run.Infof(nil, "loaded config: %s", cfgVal)

	// Store client
	ctx := context.Background()
	bqClient, err := bigquery.NewClient(ctx, run.ProjectID())
	if err != nil {
		run.Error(nil, err)
	}
	run.AddClient("bigquery", bqClient)

	// Later usage
	clientRef, err := run.GetClient("bigquery")
	if err != nil {
		run.Error(nil, err)
	}
	bqClient2 := clientRef.(*bigquery.Client)
	_ = bqClient2

	// Make service account authenticated requests
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		run.Error(nil, err)
	}
	req = run.AddOAuth2Header(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		run.Error(nil, err)
	}
	defer resp.Body.Close()
	// read response
}
