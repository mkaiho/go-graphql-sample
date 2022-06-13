ENTRY_FILE:=server.go
BIN_DIR:=_deployments/bin
BIN_NAME:=$(ENTRY_FILE:%.go=%)

.PHONY: build
build:
	go build -o $(BIN_DIR)/$(BIN_NAME) $(ENTRY_FILE)

.PHONY: dev-deps
dev-deps:
	go install gotest.tools/gotestsum@v1.8.1
	go install github.com/vektra/mockery/v2@latest
	go install github.com/99designs/gqlgen

.PHONY: deps
deps:
	go mod download

.PHONY: gen-mock
gen-mock:
	make dev-deps
	mockery --all --case underscore --recursive --keeptree

.PHONY: init-gql
init-gql:
	make dev-deps
	go run github.com/99designs/gqlgen init

.PHONY: gen-gql
gen-gql:
	make dev-deps
	go run github.com/99designs/gqlgen generate

.PHONY: clean
clean:
	@rm -rf ./${BIN_DIR}
