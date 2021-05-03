
PROTO_ROOT := "proto/"

RPC_ROOT  := "quiz/"

PROTOTOOL_VERSION := v1.9.0
PROTOC_GEN_GO_VERSION := v1.4.2
PROTOC_GEN_TWIRP_VERSION := v5.10.1


MIGRATION_OUT       := "bin/migration"
MIGRATION_MAIN_FILE := "cmd/main.go"

API_OUT := "bin/"

.PHONY: clean
clean:
	@echo "\n removing all generated files \n"
	@rm -rf $(API_OUT) $(RPC_ROOT)
	@find . -type f -name 'wire_gen.go' -delete

.PHONY: deps
deps:
	@echo "\n fetching dependencies \n"
	@export GO111MODULE=on
	@go get github.com/google/wire/cmd/wire
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	@go install google.golang.org/protobuf/cmd/protoc-gen-go

	@go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo  "\n fetching dependencies completed \n"

	@go get -u github.com/twitchtv/twirp/protoc-gen-twirp

.PHONY: run-migration-up
run-migration-up:
	@go run cmd/main.go up

.PHONY: run-migration-down
run-migration-down:
	@go run cmd/main.go down

.PHONY: proto-generate
proto-generate: deps
	@protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. proto/crm/v1/service.proto
	@echo  "\n generating proto completed \n"

.PHONY: wire-generate
wire-generate: deps
	@wire github.com/razorpay/quiz/crm
	@wire github.com/razorpay/quiz/package/orm
	@echo  "\n generating wire dependencies completed \n"

.PHONY: proto-generate-docker
proto-generate-docker:
	@echo "\n generating proto \n"
	@ls $(GOPATH)/bin
	@export PATH=$(PATH):/$(GOPATH)/"bin"
	@protoc --proto_path=$(GOPATH)/src:. --twirp_out=. --go_out=. proto/crm/v1/service.proto

.PHONY: wire-generate-docker
wire-generate-docker:
	@wire ./crm
	@wire ./package/orm
	@echo  "\n generating wire dependencies completed \n"

.PHONY: allowAll
allowAll:
	@sudo chmod -R 777 .

.PHONY: buildAllDockerApi
buildAllDockerApi: deps proto-generate-docker wire-generate-docker

.PHONY: buildAllDockerMigration
buildAllDockerMigration: deps proto-generate-docker wire-generate-docker

.PHONY: buildAll
buildAll: allowAll clean  buildAllDockerApi

.PHONY: build
build: buildAll
	@go build -o bin/main main/main.go

.PHONY: run
run:
	@go run main/main

.PHONY: quiz
quiz:build run