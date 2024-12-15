# 检查是否安装protoc
$(if $(shell ! which protoc &>/dev/null),$(error please install protoc first))

# 检查是否安装npm
$(if $(shell ! which npm &>/dev/null),$(error please install npm first))

TOOLS ?= $(BLOCKER_TOOLS) $(CRITICAL_TOOLS) $(TRIVIAL_TOOLS)

## tools.install: 安装所需的所有工具
.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

## tools.verify: 检查所需的工具是否安装
.PHONY: tools.verify
tools.verify: $(addprefix tools.verify.,$(TOOLS))

## toosl.install.<tool>: 安装<tool>;
.PHONY: tools.install.%
tools.install.%:
	@echo "==========> Install $*"
	@$(MAKE) install.$*

## tools.verify.<tool>: 检查<tool>是否安装
.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: tools.install.golangci-lint
tools.install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@echo 'source <(golangci-lint completion bash)' >>~/.bashrc
	@$(shell source ~/.bashrc)

.PHONY: tools.install.mockgen
tools.install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@v1.6.0

.PHONY: tools.install.protoc-gen-go
tools.install.protoc-gen-go:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0

.PHONY: tools.install.protoc-gen-doc
tools.install.protoc-gen-doc:
	@$(GO) install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1

.PHONY: tools.install.protoc-gen-grpc-gateway
tools.install.protoc-gen-grpc-gateway:
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

.PHONY: tools.install.protoc-gen-openapiv2
tools.install.protoc-gen-openapiv2:
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

.PHONY: tools.install.protoc-gen-go-grpc
tools.install.protoc-gen-go-grpc:
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: tools.install.protoc-gen-go-errors
tools.install.protoc-gen-go-errors:
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@v2.5.3

.PHONY: tools.install.goimports
tools.install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: tools.install.wire
tools.install.wire:
	@$(GO) install github.com/google/wire/cmd/wire@latest

.PHONY: tools.install.gobgen
tools.install.gobgen:
	@$(GO) build -o $(GOBIN)/gobgen $(ROOT_DIR)/pkg/tools/gobtools/cmd/gobgen.go

.PHONY: tools.install.stringer
tools.install.stringer:
	@$(GO) install golang.org/x/tools/cmd/stringer@latest

# deprecated
.PHONY: tools.install.sqlc
tools.install.sqlc:
	@$(GO) install github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0

.PHONY: tools.install.buf
tools.install.buf:
	@$(ROOT_DIR)/scripts/buf-install.sh

.PHONY: tools.install.protoc
tools.install.protoc:
	@$(ROOT_DIR)/scripts/protoc-install.sh

# deprecated
.PHONY: tools.install.migrate
tools.install.migrate:
	@$(GO) install -tags '$(DB_TYPE)' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: tools.install.ent
tools.install.ent:
	@$(GO) install entgo.io/ent/cmd/ent@latest

.PHONY: tools.install.pm2
tools.install.pm2:
	@npm install -g pm2