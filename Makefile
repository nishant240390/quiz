
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
		@go mod download
	@echo  "\n fetching dependencies completed \n"


.PHONY: run-migration-up
run-migration-up:
	@go run cmd/main.go up

.PHONY: run-migration-down
run-migration-down:
	@go run cmd/main.go down

.PHONY: proto-generate
proto-generate:
	@protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. proto/crm/v1/service.proto
	@echo  "\n generating proto completed \n"

.PHONY: wire-generate
wire-generate:
	@wire github.com/razorpay/quiz/crm
	@wire github.com/razorpay/quiz/package/orm
	@echo  "\n generating wire dependencies completed \n"

.PHONY: allowAll
allowAll:
	@sudo chmod -R 777 .

.PHONY: buildAll
buildAll: allowAll clean deps allowAll proto-generate wire-generate

.PHONY: build
build: buildAll
	@go build -o bin/main main.go

.PHONY: run
run:
	@go run main.go

.PHONY: quiz
quiz:build run