export GO111MODULE=on
export GOSUMDB=off

BIN_NAME := $(or $(PROJECT_NAME),'api')
PKG_PATH := $(or $(PKG),'.')
PKG_LIST := $(shell go list ${PKG_PATH}/... | grep -v /vendor/)
GOLINT := golangci-lint

MIGRATE=migrate -path migrations -database postgres://postgres:12345@localhost:5432/postgres?sslmode=disable
TEST_MIGRATE=migrate -path migrations  -database postgres://postgres:12345@localhost:5432/postgres-test?sslmode=disable

check-lint:
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.41.1

dep: # Download required dependencies
	go mod tidy
	go mod download
	go mod vendor

cilint: dep check-lint ## Lint the files local env
	$(GOLINT) run -c .golangci.yml --timeout 5m

lint: cilint

build: dep ## Build the binary file
	CGO_ENABLED=1 go build -mod=vendor -o ./bin/${BIN_NAME} -a ./src

clean: ## Remove previous build
	rm -f src/bin/$(BIN_NAME)

check-swagger:
	@which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	swagger generate spec -o src/server/http/static/swagger.yaml  -w ./ --scan-models
	swagger generate spec -o src/server/http/static/swagger.json  -w ./ --scan-models

fmt: ## format source files
	go fmt github.com/Omelman/trucking/api/src/...

migrate-up: ## Run migrations
	$(MIGRATE) up

migrate-down: ## Rollback migrations
	$(MIGRATE) down

migrate-create: ## Create migration file with name
	migrate create -ext sql -dir ./migrations -seq -digits 4 "migrate_name"
