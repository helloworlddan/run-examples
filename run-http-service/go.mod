module github.com/helloworlddan/run-examples/run-http-service

go 1.22.4

// github.com/helloworlddan/run v0.3.2
require github.com/helloworlddan/run v0.0.0-unpublished

replace github.com/helloworlddan/run v0.0.0-unpublished => ../../run/

require (
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
