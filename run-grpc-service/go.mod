module github.com/helloworlddan/run-examples/run-grpc-service

go 1.22.4

require (
	github.com/helloworlddan/run v0.0.0-unpublished
	// github.com/helloworlddan/run v0.4.1
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

replace github.com/helloworlddan/run v0.0.0-unpublished => ../../run/

require (
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)
