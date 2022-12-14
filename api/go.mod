module r-booker/api

go 1.19

require google.golang.org/grpc v1.51.0

require github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	r-booker/common v0.0.0
	r-booker/proto v0.0.0
)

replace r-booker/common v0.0.0 => ../pkg/common

replace r-booker/proto v0.0.0 => ../pkg/proto
