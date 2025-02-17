#!/bin/bash

# 确保日志目录存在
mkdir -p logs

# 开发环境启动
go run cmd/main.go -mode dev 