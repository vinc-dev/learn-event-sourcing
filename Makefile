-include .env

# Go parameters
GOCMD := go
GOBUILD :=$(GOCMD) build
GOMOD := $(GOCMD) mod
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
MOCKGEN := mockgen
BIN_NAME=grosenia-api

# Project variables
PROJECT_ROOT := $(shell pwd)
PROJECT_PKG := code.nbs.co.id/bromo-b2b/core-svc
PROJECT_MAIN_PKG := cmd/bromo-svc
PROJECT_ERROR_CODES := error-codes.yml
PROJECT_DEBUG_ENV_DIR := .api-env
PROJECT_DEBUG_ENV_FILES := $(addprefix $(PROJECT_DEBUG_ENV_DIR)/,config.yml datasources.yml assets-cred.json)
PROJECT_DEBUG_OUTPUT := $(PROJECT_ROOT)/bin/debug
PROJECT_RELEASE_OUTPUT := $(PROJECT_ROOT)/bin/release
PROJECT_NAME := "Grosenia API"

# Mocks
MOCK_DIR := $(PROJECT_ROOT)/internal/bromo-svc/mocks/
MOCK_PKG := mocks
MOCK_IN_COMPONENTS := $(PROJECT_ROOT)/internal/bromo-svc/contracts/components.go
MOCK_IN_DATASOURCES := $(PROJECT_ROOT)/internal/bromo-svc/contracts/datasources.go
MOCK_OUT_COMPONENTS := $(MOCK_DIR)/mock_components.go
MOCK_OUT_DATASOURCES := $(MOCK_DIR)/mock_datasources.go

# Test
TEST_OUT_DIR ?= $(PROJECT_ROOT)/test/out
TEST_COVERAGE_OUT ?= $(TEST_OUT_DIR)/coverage.out
TEST_COVERAGE_HTML ?= $(TEST_OUT_DIR)/coverage.html

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## install: Download dependencies
.PHONY: install
install: go.mod
	@-echo "  > Downloading dependencies..."
	@$(GOMOD) download

vendor: go.mod
	@-echo "  > Vendoring..."
	@$(GOMOD) vendor

## compile: Compile binary.
.PHONY: compile
compile: mocks
	@-echo "  > Compiling..."
	@$(GOBUILD) -o $(PROJECT_DEBUG_OUTPUT)/$(BIN_NAME) $(PROJECT_ROOT)/$(PROJECT_MAIN_PKG)
	@-echo "  > Output: $(PROJECT_DEBUG_OUTPUT)/$(BIN_NAME)"

## copy-error-codes: Copy error codes to binary
.PHONY: copy-error-codes
copy-error-codes: $(PROJECT_ERROR_CODES)
	@-echo "  > Copying error codes..."
	@-cp $(PROJECT_ERROR_CODES) $(PROJECT_DEBUG_OUTPUT)/

## copy-debug-env: Copy required environment files for
.PHONY: copy-debug-env
copy-debug-env: copy-error-codes $(PROJECT_DEBUG_ENV_FILES)
	@-echo "  > Copying env files..."
	@-cp -R $(PROJECT_DEBUG_ENV_DIR) $(PROJECT_DEBUG_OUTPUT)/

## exec: Run server
.PHONY: run
run: compile copy-debug-env
	@-echo "  > Starting Server...\n"
	@$(PROJECT_DEBUG_OUTPUT)/$(BIN_NAME)

## release: Compile binary for deployment.
.PHONY: release
release: mocks vendor
	@-echo "  > Compiling for release..."
	@-echo ""
	@CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -v -mod=vendor \
		-ldflags "-X main.AppVersion=$(CI_COMMIT_TAG) -X main.CommitHash=$(CI_COMMIT_SHA)" \
		-o $(PROJECT_RELEASE_OUTPUT)/$(BIN_NAME) $(PROJECT_ROOT)/$(PROJECT_MAIN_PKG)
	@-echo ""
	@-echo "  > Copying error codes..."
	@-cp $(PROJECT_ERROR_CODES) $(PROJECT_RELEASE_OUTPUT)/
	@-echo "  > Output: $(PROJECT_RELEASE_OUTPUT)"

## mocks: Generate mock files
.PHONY: mocks
mocks: $(MOCK_DIR)
$(MOCK_DIR): $(MOCK_OUT_COMPONENTS) $(MOCK_OUT_DATASOURCES)
$(MOCK_OUT_COMPONENTS): $(MOCK_IN_COMPONENTS)
	@-echo "  > Generating mock: components..."
	@$(MOCKGEN) \
		-source $(MOCK_IN_COMPONENTS) \
	 	-destination $(MOCK_OUT_COMPONENTS) \
	 	-package $(MOCK_PKG)
$(MOCK_OUT_DATASOURCES): $(MOCK_IN_DATASOURCES)
	@-echo "  > Generating mock: datasources..."
	@$(MOCKGEN) \
		-source $(MOCK_IN_DATASOURCES) \
		-destination $(MOCK_OUT_DATASOURCES) \
		-package $(MOCK_PKG)

$(TEST_OUT_DIR):
	@mkdir -p $(TEST_OUT_DIR)

## test: Run test in verbose mode
.PHONY: test
test: $(TEST_OUT_DIR) $(wildcard cmd/*.go) $(wildcard pkg/*.go) $(wildcard internal/*.go)
	@-echo "  > Running test..."
	@WORKING_DIR=$(PROJECT_ROOT) $(GOCMD) test -v ./...

## coverage: Run test and create coverage output
.PHONY: coverage
coverage:
	@-echo "  > Running test..."
	@WORKING_DIR=$(PROJECT_ROOT) $(GOCMD) test ./... -coverprofile=$(TEST_COVERAGE_OUT)
	@-echo "  > Test Coverage output: $(TEST_COVERAGE_OUT)"

.PHONY: html-coverage
html-coverage: $(TEST_COVERAGE_HTML)
$(TEST_COVERAGE_HTML): coverage
	@-echo "  > Generating coverage result in html..."
	@$(GOCMD) tool cover -html=$(TEST_COVERAGE_OUT) -o=$(TEST_COVERAGE_HTML)
	@-echo "  > HTML Report output: $(TEST_COVERAGE_HTML)"

.PHONY: clean
clean:
	@-echo "  > Cleaning generated files..."
	@-rm -rf ./bin
	@-rm -rf $(TEST_OUT_DIR)