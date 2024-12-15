# V-Trace

## 1. 项目概述

针对高值产品存在以次充好，真假参杂等现象，通过建立全流程、全链条的溯源平台记录高值产品从生产到出售的全过程。同时使用各种仪器进行辅助判断，通过多元评价分析模型对产品的真伪、优劣进行判断。

## 2. 启动

> 依赖： protoc v3.19.4; wire; docker; docker-compose; sqlc; ...

```bash
make all
```

## 3. 项目架构

![项目架构](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/%E9%A1%B9%E7%9B%AE%E6%9E%B6%E6%9E%84.png)

## 4. 代码结构

![代码结构](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/%E4%BB%A3%E7%A0%81%E7%BB%93%E6%9E%84.png)

## 5. 目录结构

```
.
├── api  // 接口IDL
│   ├── buf.gen.yaml // buf 配置文件
│   ├── buf.lock
│   ├── buf.yaml
│   ├── circ // circ服务接口
│   ├── ca // ca服务接口，用于签发及管理证书
│   ├── trans // transaction服务接口，用于管理交易
│   ├── goods // goods服务接口，用于管理商品
│   ├── iam // iam服务接口，用于身份管理及权限校验
│   └── vtrace // vtrace 应用接口
├── app // 服务实现
│   ├── algo // 算法模型服务实现
│   ├── goods // 产品服务实现
│   ├── ca // 证书服务实现
│   ├── echo // echo服务实现
│   ├── iam // iam服务实现
│   ├── trans // 交易服务实现
│   ├── trans-job // 交易轮询任务
│   ├── vtrace // vtrace服务实现
├── build // 部署构建相关
│   ├── chainmaker // 区块链配置
│   ├── docker // Dockerfile
│   ├── docker-compose // docker compose 配置文件
│   └── release // 发布文件
├── docs // 接口文档， markdown, 自动生成
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── pkg // 依赖包
│   ├── app // 应用
│   ├── cache // 缓存
│   ├── chain // 区块链
│   ├── constants // 常量
│   ├── endpoint //
│   ├── grpc // grpc 服务端
│   ├── logger
│   ├── plugin
│   ├── third-party // swagger 相关
│   ├── token // token maker
│   └── ...
├── README.md
├── scripts // 脚本
│   ├── bin // 一些二进制文件
│   ├── buf-install.sh
│   ├── make-rules // makefile
│   └── protoc-install.sh
└── vtrace.config.js // pm2 配置文件
```

## 6. 所属课题
- 课题名称：高值产品多元分析评价模型及质量追溯平台研发
- 课题编号：2021YFF0601204
