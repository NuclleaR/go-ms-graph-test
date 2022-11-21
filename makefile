PROTO_DIR=./pkg/proto
PROTO_FILES_DIR=$(PROTO_DIR)/buf
# PLATFORM = $(shell uname -m)

.PHONY: protoc
protoc: # Generate protobuf files
	docker-compose -f $(PROTO_DIR)/docker-compose.yaml run --rm protoc

.PHONY: install-validator
install-validator: # Install protoc-gen-validate
	@if [ -z $(VALIDATOR_VERSION) ]; then\
		echo "VALIDATOR_VERSION is not set";\
		exit 1;\
  fi
	@echo "Installing protoc-gen-validate v$(VALIDATOR_VERSION)..."
	@go install github.com/envoyproxy/protoc-gen-validate
	@echo "Copy files..."
	@mkdir -p ./$(PROTO_FILES_DIR)/validate
	@cp $(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v$(VALIDATOR_VERSION)/validate/validate.proto $(PROTO_FILES_DIR)/validate/validate.proto
	@echo "Done"

.PHONY: run-auth
run-auth: # Run auth service // TODO add docker-compose
	go run ./auth/cmd/main.go -host localhost -port 8081

.PHONY: run-api
run-api: # Run api service // TODO add docker-compose
	AUTH_SERVICE=localhost:8081 go run ./api/cmd/main.go -host localhost -port 8080