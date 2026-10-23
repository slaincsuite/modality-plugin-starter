GO ?= go
UI ?= npm
PLUGIN_NAME ?= $(shell jq -r '.name // empty' plugin/plugin.json)
PLUGIN_BINARY ?= $(shell jq -r '.binary // empty' plugin/plugin.json)
PLUGIN_DIR ?= build/plugin/$(PLUGIN_NAME)
BUILD_DIR ?= build/bin

ifeq ($(PLUGIN_NAME),)
$(error plugin/plugin.json is missing a valid name field)
endif
ifeq ($(PLUGIN_BINARY),)
$(error plugin/plugin.json is missing a valid binary field)
endif

# Detect host OS. Go output uses an .exe suffix on Windows.
ifeq ($(OS),Windows_NT)
    HOST_OS := windows
    BINARY_SUFFIX := .exe
else
    HOST_OS := $(shell uname -s 2>/dev/null | tr '[:upper:]' '[:lower:]')
    ifeq ($(HOST_OS),darwin)
        HOST_OS := macos
    endif
    BINARY_SUFFIX :=
endif

# Build for a specific platform: make package PLATFORM=linux
PLATFORM ?= $(HOST_OS)
.PHONY: test-go test-ui build-ui build-plugin package verify tidy sync-check

test-go:
	$(GO) test $(shell $(GO) list ./src/...)

test-ui:
	cd ui && $(UI) install && $(UI) run test:unit

build-ui:
	cd ui && $(UI) install && $(UI) run build

build-plugin:
	@mkdir -p $(BUILD_DIR)
ifeq ($(PLATFORM),macos)
	$(GO) build -trimpath -o $(BUILD_DIR)/$(PLUGIN_BINARY)$(BINARY_SUFFIX) ./src/apps
else ifeq ($(PLATFORM),windows)
	$(GO) build -trimpath -o $(BUILD_DIR)/$(PLUGIN_BINARY).exe ./src/apps
else ifeq ($(PLATFORM),linux)
	$(GO) build -trimpath -o $(BUILD_DIR)/$(PLUGIN_BINARY) ./src/apps
else
	@echo "unsupported platform: $(PLATFORM) (use macos, linux, or windows)"
	@exit 2
endif

package: build-ui build-plugin
	rm -rf $(PLUGIN_DIR)
	mkdir -p $(PLUGIN_DIR)
	cp $(BUILD_DIR)/$(PLUGIN_BINARY)$(BINARY_SUFFIX) $(PLUGIN_DIR)/$(PLUGIN_BINARY)
	mkdir -p $(PLUGIN_DIR)/ui
	cp plugin/plugin.json $(PLUGIN_DIR)/plugin.json
	cp -R ui/dist/. $(PLUGIN_DIR)/ui/

verify: test-go test-ui
	$(GO) vet $(shell $(GO) list ./src/...)
	@$(MAKE) --no-print-directory sync-check

tidy:
	$(GO) mod tidy

sync-check:
	@test -f AGENTS.md
	@test -f .agents/rules/development-workflow.md
	@test -f .agents/rules/synchronization.md
	@grep -q "upstream repository" .agents/rules/synchronization.md
