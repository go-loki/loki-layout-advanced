# hasaki-layout-advanced — 高级布局


Hasaki是一个基于Golang的应用脚手架，它的名字来自于英雄联盟中的亚索英雄的口头语。Hasaki是由Golang生态中各种非常流行的库整合而成的，它们的组合可以帮助你快速构建一个高效、可靠的应用程序。

[英文介绍](https://github.com/go-hasaki/hasaki-layout-advanced/blob/main/README.md

![Hasaki](https://github.com/go-hasaki/hasaki/blob/main/.github/assets/banner.png)


## 文档
* [使用指南](https://github.com/go-hasaki/hasaki/blob/main/docs/zh/guide.md)
* [分层架构](https://github.com/go-hasaki/hasaki/blob/main/docs/zh/architecture.md)
* [上手教程](https://github.com/go-hasaki/hasaki/blob/main/docs/zh/tutorial.md)

## 功能
- **Hertz**: https://github.com/cloudwego/hertz
- **Gorm**: https://github.com/go-gorm/gorm
- **Wire**: https://github.com/google/wire
- **Viper**: https://github.com/spf13/viper
- **Zap**: https://github.com/uber-go/zap
- **Golang-jwt**: https://github.com/golang-jwt/jwt
- **Go-redis**: https://github.com/go-redis/redis
- **Testify**: https://github.com/stretchr/testify
- **Sonyflake**: https://github.com/sony/sonyflake
- **gocron**:  https://github.com/go-co-op/gocron
- More...
## 特性
* **超低学习成本和定制**：Hasaki封装了Gopher最熟悉的一些流行库。您可以轻松定制应用程序以满足特定需求。
* **高性能和可扩展性**：Hasaki旨在具有高性能和可扩展性。它使用最新的技术和最佳实践，确保您的应用程序可以处理高流量和大量数据。
* **安全可靠**：Hasaki使用了稳定可靠的第三方库，确保您的应用程序安全可靠。
* **模块化和可扩展**：Hasaki旨在具有模块化和可扩展性。您可以通过使用第三方库或编写自己的模块轻松添加新功能和功能。
* **文档完善和测试完备**：Hasaki文档完善，测试完备。它提供了全面的文档和示例，帮助您快速入门。它还包括一套测试套件，确保您的应用程序按预期工作。


这是一个经典的Golang 项目的目录结构，包含以下目录：
- api: 存放应用程序的接口文件。
- assets: 存放应用程序的静态资源。
- build: 存放应用程序的编译产物。
- cmd: 存放应用程序的入口点，包括主函数和依赖注入的代码。
  - root.go: 应用程序的主要入口点，包含主命令和配置加载的代码。
  - server: server类应用程序的主要入口点，包含server主命令代码。
    - http.go: http子命令入口主，用于启动http应用程序。
  - cronjob: cronjob类应用程序的主要入口点，包含cronjob主命令代码。
    - cronjob子命令入口主，用于启动定时脚本子命令。
  - consumer: consumer类应用程序的主要入口点，包含consumer主命令代码。
    - consumer子命令入口主，用于启动消费者子命令。

- config: 存放应用程序的配置文件。
  - debug: 本地环境的配置文件。
    - config.yml: 本地环境的配置文件。

- deploy: 存放应用程序的k8s部署文件。
  - release: 生产环境的部署文件。
  - test: 测试环境的部署文件。

- internal: 存放应用程序的内部代码。
  - app: 各类命令的处理程序。
    - cronjob: 定时脚本的处理程序。
    - consumer: 消费者的处理程序。
    - server: 服务器的处理程序。
  - assembly: 使用Wire库生成的依赖注入代码。
  - middleware: 存放中间件代码。
    - cors.go: 跨域资源共享中间件。
  - model: 存放数据模型代码。
    - form: web程序的入口请求结构体。
  - repository: 存放数据访问代码。
  - server: 存放服务器代码。
    - http.go: HTTP服务器的实现。
  - service: 存放业务逻辑代码。

- pkg: 存放应用程序的公共包。
- storage: 存放应用程序的存储文件。
- web: 存放应用程序的web文件。
- go.mod: Go模块文件。
- go.sum: Go模块的依赖版本文件。

## 要求
要使用Hasaki，您需要在系统上安装以下软件：

* Golang 1.20或更高版本
* Git

### 安装

您可以通过以下命令安装Hasaki：

```bash
go install github.com/go-hasaki/hasaki@latest
```

### 创建新项目

您可以使用以下命令创建一个新的Golang项目：

```bash
hasaki new projectName
```
默认拉取github源
```
// 使用基础模板
hasaki new projectName -r https://github.com/go-hasaki/hasaki-layout-basic.git
// 使用高级模板
hasaki new projectName -r https://github.com/go-hasaki/hasaki-layout-advanced.git
```

此命令将创建一个名为`projectName`的目录，并在其中生成一个优雅的Golang项目结构。

### 启动项目

您可以使用以下命令快速启动项目：

```bash
make http
```

此命令将启动您的Golang项目，并支持文件更新热重启。

### 编译wire.go

您可以使用以下命令快速编译`wire.go`：

```bash
make gen
```

此命令将编译您的`wire.go`文件，并生成所需的依赖项。

## 贡献

如果您发现任何问题或有任何改进意见，请随时提出问题或提交拉取请求。我们非常欢迎您的贡献！

