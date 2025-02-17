# fiber-layout

> 本项目使用 go-fiber 框架为核心搭建的脚手架，可以基于本项目快速完成业务开发，开箱📦 即用

### 运行

拉取代码后在项目根目录执行如下命令：

```shell
# 开启GO111MODULE
go env -w GO111MODULE=on

# 设置代理 
go env -w GOPROXY=https://goproxy.cn,direct

# 下载依赖
go mod download

# 运行项目
go run cmd/main.go #默认启动开发环境
go run cmd/main.go -mode dev #开发环境
go run cmd/main.go -mode prod #生产环境

# 项目起来后执行下面命令访问示例路由
curl "http://127.0.0.1:3000"
# {"msg":"success","code":200,"data":"wat.ink"}
curl "http://127.0.0.1:3000/api/v1/users?page=1&page_size=1" # 入参不符合校验要求
```

### 部署

```shell
# 打包项目
go build -o app cmd/main.go

# 优化打包
go build -ldflags "-s -w" -o app cmd/main.go

# 设置linux打包环境
$ENV:GOARCH="amd64" # x86架构      $ENV:GOARCH="arm64" #arm架构
$ENV:GOOS="linux"
go build -o app cmd/main.go

# 服务器运行
nohup ./app -mode prod > app.log 2>&1 &

# nginx 反向代理配置示例
server {
    listen 80;
    server_name api.example.com;
    location / {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:3000;
    }
}
```

### 目录结构

```
.
├── api
│   └── v1
│       ├── request      // 请求结构体
│       └── response     // 响应结构体
├── cmd
│   └── main.go         // 程序入口
├── config              // 配置文件
│   ├── config.dev.yaml
│   └── config.prod.yaml
├── internal
│   ├── controller      // 控制器
│   ├── middleware      // 中间件
│   ├── model          // 数据模型
│   ├── repository     // 数据访问
│   ├── router         // 路由定义
│   └── service        // 业务逻辑
├── pkg
│   ├── config         // 配置管理
│   ├── database       // 数据库连接
│   ├── email          // 邮件服务
│   ├── errors         // 错误定义
│   ├── jwt           // JWT 认证
│   ├── logger        // 日志管理
│   ├── rabbitmq      // 消息队列
│   ├── redis         // Redis 缓存
│   ├── utils         // 工具函数
│   └── validator     // 参数验证
├── logs              // 日志文件
├── go.mod
├── go.sum
└── README.md
```

### 主要特性

- 完整的项目结构和最佳实践
- 统一的错误处理和响应
- JWT 认证中间件
- 参数验证和绑定
- 日志记录和管理
- 数据库迁移和初始化
- 消息队列集成
- 邮件服务支持
- Redis 缓存支持
- 配置热重载
- 优雅关机

### 使用的包

- Web 框架：[fiber](https://github.com/gofiber/fiber)
- 配置管理：[viper](https://github.com/spf13/viper)
- 参数验证：[validator](https://github.com/go-playground/validator)
- 日志：[zap](https://github.com/uber-go/zap)
- 数据库：[gorm](https://github.com/go-gorm/gorm)
- 缓存：[go-redis](https://github.com/go-redis/redis)
- 消息队列：[amqp091-go](https://github.com/rabbitmq/amqp091-go)

### 开发流程

1. 定义 API 请求和响应结构体
2. 实现数据模型和仓储层
3. 编写业务逻辑服务
4. 实现控制器处理请求
5. 注册路由
6. 编写中间件（可选）
7. 添加单元测试

### 贡献

欢迎提交 Issue 和 Pull Request！