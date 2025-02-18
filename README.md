# NextEraAbyss/FiberForge

> 企业级Go Web开发基础框架 | 基于Fiber的高性能脚手架 | 开箱📦 即用

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
NextEraAbyss/FiberForge
├── api                 # API 相关定义
│   └── v1             # API 版本控制
│       ├── request    # 请求参数结构体
│       └── response   # 响应结构体定义
│
├── cmd                # 应用程序入口
│   └── main.go       # 主程序入口文件
│
├── config            # 配置文件目录
│   ├── config.dev.yaml   # 开发环境配置
│   └── config.prod.yaml  # 生产环境配置
│
├── internal          # 内部应用代码
│   ├── controller   # 控制器层：处理请求响应
│   ├── middleware   # 中间件：认证、日志等
│   ├── model       # 数据模型定义
│   ├── repository  # 数据访问层
│   ├── router      # 路由注册管理
│   └── service     # 业务逻辑层
│
├── pkg              # 公共代码包
│   ├── config      # 配置管理工具
│   ├── database    # 数据库相关
│   │   ├── migration.go  # 数据库迁移
│   │   └── mysql.go      # MySQL连接管理
│   ├── email       # 邮件服务模块
│   ├── errors      # 错误处理模块
│   ├── jwt         # JWT认证工具
│   ├── logger      # 日志管理模块
│   ├── rabbitmq    # 消息队列工具
│   ├── redis       # Redis缓存工具
│   ├── utils       # 通用工具函数
│   └── validator   # 参数验证工具
│
├── logs            # 日志文件目录
├── scripts         # 脚本文件目录
│   ├── start.sh    # 启动脚本
│   └── init.sql    # 数据库初始化脚本
│
├── .gitignore      # Git忽略文件
├── go.mod          # Go模块文件
├── go.sum          # Go依赖版本锁定
└── README.md       # 项目说明文档
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