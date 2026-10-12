.PHONY: setup generate test clean verify help format

# Version management
VERSION_FILE := VERSION
VERSION ?= $(shell if [ -f "$(VERSION_FILE)" ]; then cat "$(VERSION_FILE)"; else date +%Y%m%d.0; fi)

# Configuration
SWAGGER_FILE ?= api/swagger.json
CODE_GEN_DIR = scripts/codegen
BUILD_JOBS ?= 4

# Color definitions
ifneq (,$(findstring xterm,${TERM}))
    RED          := $(shell tput -Txterm setaf 1)
    GREEN        := $(shell tput -Txterm setaf 2)
    YELLOW       := $(shell tput -Txterm setaf 3)
    BLUE         := $(shell tput -Txterm setaf 4)
    NC           := $(shell tput -Txterm sgr0) # No Color
else
    RED          := \033[0;31m
    GREEN        := \033[0;32m
    YELLOW       := \033[1;33m
    BLUE         := \033[0;34m
    NC           := \033[0m # No Color
endif

# Default number of parallel build jobs
BUILD_JOBS ?= 4

# Install dependencies
setup:
	@echo "${BLUE}=== Installing required tools ===${NC}"
	@if ! command -v swagger-codegen >/dev/null 2>&1; then \
		echo "${YELLOW}Installing swagger-codegen...${NC}"; \
		if ! brew install swagger-codegen; then \
			echo "${RED}Failed to install swagger-codegen. Please install it manually.${NC}" >&2; \
			exit 1; \
		fi; \
	else \
		echo "${GREEN}✓ swagger-codegen is already installed${NC}"; \
	fi
	@if ! command -v goimports >/dev/null 2>&1; then \
		echo "${YELLOW}Installing goimports...${NC}"; \
		if ! go install golang.org/x/tools/cmd/goimports@latest; then \
			echo "${RED}Failed to install goimports. Please install it manually.${NC}" >&2; \
			exit 1; \
		fi; \
	else \
		echo "${GREEN}✓ goimports is already installed${NC}"; \
	fi

	@mkdir -p $(CODE_GEN_DIR)/generated
	@echo "${GREEN}✓ Setup complete${NC}"

# Generate the client code
generate: setup
	@echo "${BLUE}=== Generating Go client from $(SWAGGER_FILE) ===${NC}"
	@if [ ! -f "$(CODE_GEN_DIR)/update-client.sh" ]; then \
		echo "${RED}Error: Script not found at $(CODE_GEN_DIR)/update-client.sh${NC}" >&2; \
		echo "Current directory: $$(pwd)" >&2; \
		echo "Looking for: $(CODE_GEN_DIR)/update-client.sh" >&2; \
		exit 1; \
	fi
	@chmod +x "$(CODE_GEN_DIR)/update-client.sh"
	@(cd "$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))" && \
	env \
		SWAGGER_FILE="$(SWAGGER_FILE)" \
		VERSION="$(VERSION)" \
		BUILD_JOBS="$(BUILD_JOBS)" \
		AUTO_FIX="$(AUTO_FIX)" \
		$(if $(DRY_RUN),DRY_RUN=true,) \
		$(if $(VERBOSE),VERBOSE=true,) \
		./$(CODE_GEN_DIR)/update-client.sh) || { \
		echo "${RED}❌ Code generation failed. See above for details.${NC}" >&2; \
		exit 1; \
	}
	@if [ "$(DRY_RUN)" = "true" ]; then \
		echo "${YELLOW}✓ Dry run completed. No changes were made.${NC}"; \
	else \
		echo "${GREEN}✓ Code generation completed successfully${NC}"; \
	fi

# Format Go code
format:
	@echo "${BLUE}=== Formatting Go code ===${NC}"
	@if ! command -v goimports >/dev/null 2>&1; then \
		echo "${YELLOW}goimports not found. Installing...${NC}"; \
		go install golang.org/x/tools/cmd/goimports@latest; \
	fi
	@goimports -w -l .
	@echo "${GREEN}✓ Code formatting complete${NC}"

# Dry-run target for previewing changes
dry-run: export DRY_RUN=true
dry-run: generate
	@# This target runs the generate target with DRY_RUN=true

# Run tests
test: generate
	@echo "${BLUE}=== Running tests ===${NC}"
	@echo "${YELLOW}Running tests with $(BUILD_JOBS) parallel jobs...${NC}"
	@if ! go test -v -p $(BUILD_JOBS) ./...; then \
		echo "${RED}❌ Tests failed. See above for details.${NC}" >&2; \
		exit 1; \
	fi
	@echo "${GREEN}✓ All tests passed${NC}"

# Clean generated files
clean:
	@echo "${BLUE}=== Cleaning generated files ===${NC}"
	@rm -rf api/ model/ docs/ client.go configuration.go git_push.sh *.md
	@echo "${GREEN}✓ Clean complete${NC}"

# Remove all generated and downloaded files
distclean: clean
	@echo "${BLUE}=== Removing all generated files ===${NC}"
	@rm -rf $(CODE_GEN_DIR)/generated/
	@echo "${GREEN}✓ All generated files removed${NC}"

# Verify the generated code
verify: generate format test
	@echo "${GREEN}✓ Verification complete${NC}"

# Show help
help:
	@echo "${BLUE}=== Harness Chaos Go SDK Makefile ===${NC}"
	@echo "Available targets:"
	@echo "  ${YELLOW}make setup${NC}     - Install required dependencies and create necessary directories"
	@echo "  ${YELLOW}make generate${NC}  - Generate the Go client from the OpenAPI spec"
	@echo "  ${YELLOW}make test${NC}      - Run tests (runs generate first)"
	@echo "  ${YELLOW}make format${NC}    - Format Go code using goimports"
	@echo "  ${YELLOW}make verify${NC}    - Generate, format, and test the code"
	@echo "  ${YELLOW}make clean${NC}     - Clean generated files"
	@echo "  ${YELLOW}make distclean${NC} - Remove all generated files"
	@echo "  ${YELLOW}make help${NC}      - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  ${YELLOW}VERSION${NC}      - Version number (default: from VERSION file or YYYYMMDD.0)"
	@echo "  ${YELLOW}SWAGGER_FILE${NC} - Path to OpenAPI spec (default: api/swagger.yaml)"
	@echo "  ${YELLOW}BUILD_JOBS${NC}   - Parallel test jobs (default: 4, use -jN to override)"
	@echo "  ${YELLOW}DRY_RUN${NC}      - Set to 'true' for dry run (no changes)"
	@echo "  ${YELLOW}AUTO_FIX${NC}     - Auto-fix formatting issues (default: false)"
	@echo "  ${YELLOW}VERBOSE${NC}      - Enable verbose output (default: false)"

# Default target
all: verify
