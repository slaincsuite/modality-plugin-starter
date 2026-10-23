GO ?= go
UI ?= npm
PLUGIN_NAME ?= modality-plugin-starter
PLUGIN_DIR ?= build/plugin/$(PLUGIN_NAME)
.PHONY: test-go test-ui build-ui build-plugin package verify tidy sync-check

test-go:
	$(GO) test $(shell $(GO) list ./src/...)

test-ui:
	cd ui && $(UI) install && $(UI) run test:unit

build-ui:
	cd ui && $(UI) install && $(UI) run build

build-plugin:
	$(GO) build -trimpath -o $(PLUGIN_DIR)/$(PLUGIN_NAME) ./src/modality-plugin-starter

package: build-ui build-plugin
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
