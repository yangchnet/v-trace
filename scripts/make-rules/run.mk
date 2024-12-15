APPS ?= openapi iam-service

## run: 使用pm2启动所有服务
.PHONY: run
run: tools.verify.pm2 gen
	@pm2 start vtrace.config.js

## stop: 停止pm2守护的所有服务
.PHONY: stop
stop: tools.verify.pm2
	@pm2 del all

## run.<service>: 编译并运行service服务
.PHONY: run.%
run.%:
	@echo "==========> Starting $* server"
	@go build -o _output/app/debug-$* app/$*/cmd/$*.go app/$*/cmd/wire_gen.go && chmod +x _output/app/debug-$* && _output/app/debug-$*

## contract.deploy: 根据配置文件部署合约
contrace.deploy:
	@go run app/pkg/contract/deploy.go app/pkg/contract/config/