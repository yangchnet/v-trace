DOCKER := docker

DOCKER_COMPOSE := docker$(SPACE)compose

BASE_IMAGES := mysql redis

SERVICE_IMAGES := iam vtrace goods ca trans circ algo

CRON_IMAGES := trans-job

DOKCER_HUB_USER := dockerslc

TO_BUILD := base service cron chainmaker gateway

## image.build: 构建运行所需的所有镜像
.PHONY: image.build
image.build: $(addprefix image.build.,$(TO_BUILD))

## image.build.base: 构建运行所需的基础镜像
.PHONY: image.build.base
image.build.base: $(addprefix image.build.base.,$(BASE_IMAGES))

## image.build.service: 构建运行所需的服务镜像
.PHONY: image.build.service
image.build.service: image.builder $(addprefix image.build.service.,$(SERVICE_IMAGES))

## image.build.cron
.PHONY: image.build.cron
image.build.cron: $(addprefix image.build.cron.,$(CRON_IMAGES))

## image.build.chainmaker
.PHONY: image.build.chainmaker
image.build.chainmaker:
	@$(DOCKER) build --file $(ROOT_DIR)/build/docker/chainmaker/Dockerfile -t $(PROJECT_NAME)/chainmaker $(ROOT_DIR)

## image.build.gateway
.PHONY: image.build.gateway
image.build.gateway:
	@$(DOCKER) build --file $(ROOT_DIR)/build/docker/gateway/Dockerfile -t $(PROJECT_NAME)/kong-gateway $(ROOT_DIR)

# 编译二进制文件
.PHONY: image.builder
image.builder:
	@$(DOCKER) build --file $(ROOT_DIR)/build/docker/builder/Dockerfile -t $(PROJECT_NAME)/builder $(ROOT_DIR)

## image.build.base.<image>: 构建<image>镜像（基础设施）
.PHONY: image.build.base.%
image.build.base.%:
	@$(DOCKER) build --file $(ROOT_DIR)/build/docker/$*/Dockerfile -t $(PROJECT_NAME)/$* $(ROOT_DIR)

## image.build.service.<image>: 构建<image>镜像（服务）
.PHONY: image.build.service.%
image.build.service.%:
	$(eval NAME := $(shell echo $* | tr '[a-z]' '[A-Z]'))
	@$(DOCKER) build \
	 --build-arg SERVICE_NAME=$* \
	 --build-arg SERVICE_PORT=$($(NAME)_PORT) \
	 --file $(ROOT_DIR)/build/docker/service/Dockerfile \
	 -t $(PROJECT_NAME)/$* $(ROOT_DIR)

.PHONY: iamge.build.cron.%
image.build.cron.%:
	@$(DOCKER) build \
	 --build-arg CRON_NAME=$* \
	 --file $(ROOT_DIR)/build/docker/cron/Dockerfile \
	 -t $(PROJECT_NAME)/$* $(ROOT_DIR)

## image.up.mysql: 启动MySQL容器
.PHONY: image.up.mysql
image.up.mysql:
	@$(DOCKER) run --name $(PROJECT_NAME)-mysql-runtime -e MYSQL_ROOT_PASSWORD=$(PASSWORD) --restart=always \
		-p $(MYSQL_PORT):13306 -v ~/.$(PROJECT_NAME)-data/mysql:/var/lib/mysql -d $(PROJECT_NAME)/mysql:latest

## image.up.redis: 启动Redis容器
.PHONY: image.up.redis
image.up.redis:
	@$(DOCKER) run --name $(PROJECT_NAME)-redis-runtime --restart=always -p 16379:16379 -d $(PROJECT_NAME)/redis:latest

## image.del.%: 删除正在运行中的容器
.PHONY: image.del.%
image.del.%:
	@$(DOCKER) rm -f $(PROJECT_NAME)-$*-runtime

.PHONY: image.push
image.push: $(addprefix image.push.,$(SERVICE_IMAGES)) $(addprefix image.push.,$(CRON_IMAGES)) $(addprefix image.push.,$(BASE_IMAGES)) image.push.chainmaker image.push.kong-gateway

.PHONY: image.push.%
image.push.%:
	@$(DOCKER) tag $(PROJECT_NAME)/$*:latest $(DOKCER_HUB_USER)/$(PROJECT_NAME)-$*:latest
	@$(DOCKER) push dockerslc/$(PROJECT_NAME)-$*:latest

## compose.up: 启动所有容器
.PHONY: compose.up
compose.up:
	@$(DOCKER_COMPOSE) -f build/docker-compose/vtrace.yml up -d

## compose.down: 关闭所有容器
.PHONY: compose.down
compose.down:
	@$(DOCKER_COMPOSE) -f build/docker-compose/vtrace.yml down

## compose.base.up: 启动基础容器
.PHONY: compose.base.up
compose.base.up:
	@$(DOCKER_COMPOSE) -f build/docker-compose/vtrace_base.yml up -d

## compose.base.down: 启动基础容器
.PHONY: compose.base.down
compose.base.down:
	@$(DOCKER_COMPOSE) -f build/docker-compose/vtrace_base.yml down


