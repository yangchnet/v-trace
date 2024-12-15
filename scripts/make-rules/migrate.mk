MIGRATE := migrate

MIGRATE_DBS := algo goods circ iam trans

## migrate.up: 对所有数据库进行正向迁移
.PHONY: migrate.up
migrate.up: $(addprefix migrate.up.,$(MIGRATE_DBS))

## migrate.up.<db>: 对<db>进行数据库正向迁移, <db>一般和<service>同名
.PHONY: migrate.up.%
migrate.up.%: tools.verify.$(MIGRATE)
	@echo $(DB_NAMES)
	@$(MIGRATE) -source=file://app/$*/internal/data/migrations/ -database \
	 "$(DB_TYPE)://root:$(PASSWORD)@tcp(localhost:$(DB_PORT))/$*" \
	 -verbose up

## migrate.down: 对所有数据库进行反向迁移
.PHONY: migrate.down
migrate.down: $(addprefix migrate.down.,$(MIGRATE_DBS))

## migrate.down.<db>: 对<db>进行数据库反向迁移, <db>一般和<service>同名
.PHONY: migrate.down.%
migrate.down.%: tools.verify.$(MIGRATE)
	@$(MIGRATE) -source=file://app/$*/internal/data/migrations/ -database \
	 "$(DB_TYPE)://root:$(PASSWORD)@tcp(localhost:$(DB_PORT))/$*" \
	 -verbose down

## migrate.goto.<db>.<version>: 将<db>迁移到版本<version>
.PHONY: migrate.goto.%
migrate.goto.%: tools.verify.$(MIGRATE)
	@$(MIGRATE) -source=file://app/$*/internal/data/migrations/ -database \
	 "$(DB_TYPE)://root:$(PASSWORD)@tcp(localhost:$(DB_PORT))/$(word 1,$(subst .,$(SPACE),$*))" \
	 -verbose goto $(word 2,$(subst .,$(SPACE),$*))

## migrate.new.<service>.<name>: 为<service>创建一个migrate
migrate.new.%:
	@$(eval list := $(subst .,$(SPACE),$*))
	@$(eval service := $(word 1,$(list)))
	@$(eval name := $(word 2,$(list)))
	@$(MIGRATE) create -ext sql -dir app/$(service)/internal/data/migrations $(name)


# Deprecated
# migrate.gen.<db>.<name>: 根据ent schema生成名为<name>的迁移文件
.PHONY: migrate.gen
migrate.gen.%:
	@$(eval list := $(subst .,$(SPACE),$*))
	@$(eval db := $(word 1,$(list)))
	@$(eval name := $(word 2,$(list)))
	@cd app/$(db)/internal/data && go run migrate_gen.go $(name)
