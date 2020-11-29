
PROTO_ROOT := "proto/"
RPC_ROOT  := "quiz/"

PROTOTOOL_VERSION := v1.9.0
PROTOC_GEN_GO_VERSION := v1.4.2
PROTOC_GEN_TWIRP_VERSION := v5.10.1


API_OUT := "bin/"

.PHONY: clean
clean:
	@echo "\n removing all generated files \n"
	@rm -rf $(API_OUT) $(RPC_ROOT)


.PHONY: deps
deps:
	@echo "\n fetching dependencies 1\n"
	@go get google.golang.org/grpc
	@go get github.com/uber/prototool/cmd/prototool
	@go get github.com/golang/protobuf/protoc-gen-go
	@go get github.com/twitchtv/twirp/protoc-gen-twirp
	@go get github.com/bykof/go-plantuml
	@echo  "\n fetching dependencies completed \n"

.PHONY: deps
prebuild: deps

.PHONY: proto-generate
proto-generate:
	@protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. proto/crm/v1/service.proto

.PHONY: build-toml
build-toml: proto-generate

.PHONY: allowAll
allowAll:
	sudo chmod -R 777 .

.PHONY: buildAll
buildAll: clean deps build-toml allowAll

.PHONY: build
build: buildAll
	go build -o bin/main main.go

.PHONY: run
run:
	go run main.go

.PHONY: build-quiz
build-quiz:build run