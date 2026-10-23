GO ?= go
UI ?= npm
PLUGIN_NAME ?= $(shell jq -r '.name // empty' plugin/plugin.json)
PLUGIN_BINARY ?= $(shell jq -r '.binary // empty' plugin/plugin.json)
PLUGIN_DIR ?= build/plugin/$(PLUGIN_NAME)

ifeq ($(PLUGIN_NAME),)
$(error plugin/plugin.json is missing a valid name field)
endif
ifeq ($(PLUGIN_BINARY),)
$(error plugin/plugin.json is missing a valid binary field)
endif

# Keep command names portable; plugin output always uses the manifest binary name.
ifeq ($(OS),Windows_NT)
    RM_RF := rmdir /s /q
    COPY := xcopy /e /i /y
    MKDIR := mkdir
else
    RM_RF := rm -rf
    COPY := cp -R
    MKDIR := mkdir -p
endif

.PHONY: test-go test-ui build-ui build-plugin package verify tidy sync-check

test-go:
	$(GO) test $(shell $(GO) list ./src/...)

test-ui:
	cd ui && $(UI) install && $(UI) run test:unit

build-ui:
	cd ui && $(UI) install && $(UI) run build

build-plugin:
	@$(MKDIR) $(PLUGIN_DIR)
	$(GO) build -trimpath -o $(PLUGIN_DIR)/$(PLUGIN_BINARY) ./src/apps

package: build-ui build-plugin
	@$(MKDIR) $(PLUGIN_DIR)
	@$(RM_RF) $(PLUGIN_DIR)/ui
	@$(MKDIR) $(PLUGIN_DIR)/ui
	cp plugin/plugin.json $(PLUGIN_DIR)/plugin.json
	$(COPY) ui/dist/. $(PLUGIN_DIR)/ui/

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
