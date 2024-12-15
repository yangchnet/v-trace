.DEFAULT_GOAL := all

.PHONY: all
all: tools.verify gen image.build compose.up
	@sleep 50s
	@$(MAKE) migrate.up

# ==============================================================
# Build option

ROOT_PACKAGE=gitee.com/qciip-icp/v-trace
VERSION_PACAKGE=github.com/yangchnet/component-base/pkg/version

# ==============================================================
# Include
include scripts/make-rules/common.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/migrate.mk
include scripts/make-rules/gen.mk
include scripts/make-rules/run.mk
include scripts/make-rules/image.mk

HELP_LIST = tools gen run image golang migrate
## help: Show this help info.
.PHONY: help
help: $(addsuffix .mk,$(addprefix scripts/make-rules/,$(HELP_LIST)))
	@echo -e "\nUsage: make ...\n\nTargets:"
	@sed -n 's/^##//p' $^ | column -t -s ':' | sed -e 's/^/ /'