
PROTO_ROOT := "proto/"
RPC_ROOT  := "quiz/"

CRM_ROOT := "crm/"
ORM_ROOT := "package/orm/"

PROTOTOOL_VERSION := v1.9.0
PROTOC_GEN_GO_VERSION := v1.4.2
PROTOC_GEN_TWIRP_VERSION := v5.10.1


API_OUT := "bin/"

.PHONY: clean
clean:
	@echo "\n removing all generated files \n"
	@rm -rf $(API_OUT) $(RPC_ROOT)
	@find . -type f -name 'wire_gen.go' -delete

.PHONY: deps
deps:
	@echo "\n fetching dependencies \n"
	@go get google.golang.org/grpc
	@go get github.com/uber/prototool/cmd/prototool
	@go get github.com/golang/protobuf/protoc-gen-go
	@go get github.com/twitchtv/twirp/protoc-gen-twirp
	@go get github.com/bykof/go-plantuml
	@go get github.com/google/wire/cmd/wire
	@echo  "\n fetching dependencies completed \n"

.PHONY: deps
prebuild: deps

.PHONY: proto-generate
proto-generate:
	@protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. proto/crm/v1/service.proto

.PHONY: wire-generate
wire-generate:
	@cd $(ORM_ROOT) && \
	wire	&& \
	cd ../..	&& \
	cd $(CRM_ROOT) && \
	wire	&& \
	cd ..

.PHONY: allowAll
allowAll:
	sudo chmod -R 777 .

.PHONY: buildAll
buildAll: clean deps proto-generate wire-generate allowAll

.PHONY: build
build: buildAll
	go build -o bin/main main.go

.PHONY: run
run:
	go run main.go

.PHONY: build-quiz
build-quiz:build run