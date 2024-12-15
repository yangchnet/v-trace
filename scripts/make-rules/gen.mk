GEN_LIST ?= proto sqlc config wire

## gen: 运行所有代码生成命令
.PHONY: gen
gen: $(addprefix gen., $(GEN_LIST))

## gen.wire: 运行所有wire代码生成命令
WIRE_PACKAGES ?= $(wildcard $(ROOT_DIR)/app/*/cmd/)
.PHONY: gen.wire
gen.wire: tools.verify.wire
	@echo "==========> Generating from wire.go"
	@wire gen $(WIRE_PACKAGES)

SQLC_GEN_LIST ?= algo goods circ trans iam

## gen.sqlc: 运行所有sqlc代码生成命令
.PHONY: gen.sqlc
gen.sqlc: $(addprefix gen.sqlc., $(SQLC_GEN_LIST))

## gen.sqlc.<service>: 为<service>服务运行sqlc代码生成命令
.PHONY: gen.sqlc.%
gen.sqlc.%: tools.verify.sqlc
	@echo "==========> Generating query statement from sql for $*"
	@sqlc generate -f app/$*/internal/data/sqlc.yaml

## gen.proto: 使用protoc根据proto文件生成pb.go
PROTO_GENS := api
.PHONY: gen.proto
gen.proto: $(addprefix gen.proto., $(PROTO_GENS))

## gen.proto.api: 使用protoc根据proto文件生成pb.go
.PHONY: gen.proto.api
gen.proto.api: $(addprefix tools.verify., $(PROTO_TOOLS))
	@echo "==========> Generating pb.go for api"
	@cd api && buf build && buf generate

## gen.proto.update: 使用buf update更新proto依赖
.PHONY: gen.proto.update
gen.proto.update: $(addprefix tools.verify., $(PROTO_TOOLS))
	@cd api && buf mod update
	@cd internal && buf mod update

# gen.config: 生成配置文件
CONFS ?= iam vtrace goods ca trans circ algo trans-job
.PHONY: gen.config
gen.config: $(addprefix gen.config.,$(CONFS))

.PHONY: gen.config.%
gen.config.%:
	@echo "==========> Generating config for $*"
ifeq ($(VTRACE_DEPLOY_ENV),prd)
	@$(ROOT_DIR)/scripts/genconfig.sh $(ROOT_DIR)/scripts/.env $(ROOT_DIR)/app/$*/config/configs.prd.tmpl $(ROOT_DIR)/app/$*/config/configs.yaml
else
	@$(ROOT_DIR)/scripts/genconfig.sh $(ROOT_DIR)/scripts/.env $(ROOT_DIR)/app/$*/config/configs.tmpl $(ROOT_DIR)/app/$*/config/configs.yaml
endif

GOBS ?= iam goods trans circ algo
.PHONY: gen.gob
gen.gob: $(addprefix gen.gob.,$(GOBS))

.PHONY: gen.gob.%
gen.gob.%: tools.verify.gobgen
	@gobgen $(ROOT_DIR)/app/$*/internal/data/db/models.go

# >>>>>>>>>>>>>>>>>deprecated>>>>>>>>>>>>>>>>>>

# ## gen.error: 生成错误码
# .PHONY: gen.error
# gen.error: tools.verify.stringer
# 	@echo "==========> Generating error code"
# 	@cd app/pkg/gerr/code && go generate

# ## gen.conf: 为所有<service>生成配置文件【已弃用】
# .PHONY: gen.conf
# gen.conf: $(addprefix gen.conf.,$(SERVICE_LIST))

# ## gen.conf.<service>: 根据proto文件为<service>生成conf.pb.go
# .PHONY: gen.conf.%
# gen.conf.%: tools.verify.protoc-gen-go
# 	@echo "==========> Generating conf.pb.go for $*"
# 	@protoc --proto_path=. --go_out=. --go_opt=paths=source_relative \
# 	 app/$*/internal/conf/conf.proto


# # .PHONY: gen.doc
# # gen.doc: $(addprefix gen.doc.,$(SERVICE_LIST))

# # .PHONY: gen.doc.%
# # gen.doc.%:
# # 	@echo "==========> Generating doc for $*"
# # 	@protoc -I. -I/usr/local/include --doc_out=docs/ --doc_opt=markdown,$*.md api/$*/v1/$*.proto



# # ent generate

# ENT = ent
# ## gen.ent.new.<service>.<table>: 在<service>中新建一个<table>表
# .PHONY: gen.ent.new.%
# gen.ent.new.%: tools.verify.ent
# 	@$(eval list := $(subst .,$(SPACE),$*))
# 	@$(eval service := $(word 1,$(list)))
# 	@$(eval table := $(word 2,$(list)))
# 	@cd $(ROOT_DIR)/app/$(service)/internal/data && $(ENT) init $(table)

# ENT_GEN_LIST ?= trans echo
# ## gen.ent: 运行所有ent生成
# .PHONY: gen.ent
# gen.ent: $(addprefix gen.ent.,$(ENT_GEN_LIST))

# ## ent.gen.<service>: 对<service>服务运行ent生成
# .PHONY: gen.ent.%
# gen.ent.%: tools.verify.ent
# 	@echo "==========> Generating ent $* file"
# 	@cd $(ROOT_DIR)/app/$*/internal/data && $(GO) generate ./ent

# ## gen.app.<service>: 生成服务框架代码, 例如：make gen.app.Iam
# .PHONY: gen.app.%
# gen.app.%:
# 	@echo "==========> Generating app $* framework"
# 	@$(GO) run pkg/gen/gen.go -service $*
# 	@$(MAKE) gen.ent.new.$*.DeleteMe
# 	@$(MAKE) gen.ent.$*
# 	@$(MAKE) gen.wire
# 	@echo "Successful! Next: 修改配置文件, 填充grpc端口与数据库dsn"

# <<<<<<<<<<<<<<<<<<<<<<deprecated<<<<<<<<<<<<<<