
# fiber-layout

> 本项目使用 go-fiber 框架为核心搭建的**api**架构脚手架，可以基于本项目快速完成业务开发，开箱📦 即用


###  运行

拉取代码后在项目根目录执行如下命令：

```shell
# 建议开启GO111MODULE
go env -w GO111MODULE=on

# 下载依赖
go mod download

# 运行项目
go run main.go #默认启动开发环境
go run main.go -mode dev #开发环境
go run main.go -mode prod #生产环境

# 项目起来后执行下面命令访问示例路由
curl "http://127.0.0.1:3000/v1/register?username=admin&password=132456"
# {"code": 1000,"data": "✋ ---- admin","msg": "操作成功"}
```

### 部署

```shell
# 1. bee打包项目

# 安装 bee工具（beego框架带的打包工具）
go get -u github.com/beego/bee/v2

# 运行时请配置指定config文件的位置，否则可能会出现找不到配置的情况，修改完配置请重启
bee pack -be GOOS=linux

# 服务器 nohup工具 跑起来 (在直接放到 服务器跑就行啦)
nohup ./fiber-layout

# 2. go打包

# 设置linux打包环境
$ENV:GOOS="linux"

# go build
go build mian.go

# 服务器 nohup工具 跑起来 (在直接放到 服务器跑就行啦)
nohup ./fiber-layout
```

### 目录结构

```
.
|——.gitignore
|——go.mod
|——go.sum
|——main.go    // 项目入口 main 包
|——LICENSE
|——README.md
|——boot    // 项目初始化目录
|  └──boot.go
|——config    // 这里通常维护一些本地调试用的样例配置文件
|  └──autoload    // 配置文件的结构体定义包
|     └──app.go
|     └──logger.go
|     └──mysql.go
|     └──redis.go
|     └──server.go
|  └──config.example.ini    // .ini 配置示例文件
|  └──config.example.yaml    // .yaml 配置示例文件
|  └──config.go    // 配置初始化文件
|——data    // 数据初始化目录
|  └──data.go
|  └──mysql.go
|  └──redis.go
|——internal    // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
|  └──controller    // 控制器代码
|     └──v1
|        └──auth.go    // 完整流程演示代码，包含数据库表的操作
|        └──helloword.go    // 基础演示代码
|     └──base.go
|  └──middleware    // 中间件目录
|     └──cors.go
|     └──logger.go
|     └──recovery.go
|     └──requestCost.go
|  └──model    // 业务数据访问
|     └──admin_users.go
|     └──base.go
|  └──pkg    // 内部使用包
|     └──errors    // 错误定义
|        └──code.go
|        └──en-us.go
|        └──zh-cn.go
|     └──logger    // 日志处理
|        └──logger.go
|     └──response    // 统一响应输出
|        └──response.go
|  └──routers    // 路由定义
|     └──apiRouter.go
|     └──router.go
|  └──service    // 业务逻辑
|     └──auth.go
|  └──validator    // 请求参数验证器
|     └──form    // 表单参数定义
|        └──auth.go
|     └──validator.go
|——pkg    // 可以被外部使用的包
|  └──convert    // 数据类型转换
|     └──convert.go
|  └──utils    // 帮助函数
|     └──utils.go
```
### 其他说明

##### 项目中使用到的包

- 核心：[fiber](https://github.com/gofiber/fiber)
- 配置：[gopkg.in/yaml.v3](https://github.com/go-yaml/yaml)
- 参数验证：[github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
-日志：[go.uber.org/zap](https://github.com/uber-go/zap)、[github.com/natefinch/lumberjack](http://github.com/natefinch/lumberjack)、[github.com/lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
- 数据库：[gorm.io/gorm](https://github.com/go-gorm/gorm)、[go-redis/v8](https://github.com/go-redis/redis)
- 还有其他不一一列举，更多请查看`go.mod`文件

### 代码贡献

不完善的地方，欢迎大家 Fork 并提交 PR！