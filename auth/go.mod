module r-booker/auth

go 1.19

require (
	github.com/envoyproxy/protoc-gen-validate v0.9.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	go.opentelemetry.io/otel v1.10.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hellofresh/health-go/v5 v5.0.0
	github.com/joho/godotenv v1.4.0
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20210226172003-ab064af71705 // indirect
	r-booker/common v0.0.0
	r-booker/proto v0.0.0
)

replace r-booker/common v0.0.0 => ../pkg/common

replace r-booker/proto v0.0.0 => ../pkg/proto
