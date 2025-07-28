#!/bin/bash

# 情指行风险洞察系统构建脚本

echo "开始构建情指行风险洞察系统..."

# 设置环境变量
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# 清理之前的构建
echo "清理之前的构建文件..."
rm -rf build/
mkdir -p build

# 下载依赖
echo "下载依赖..."
go mod tidy

# 运行测试
echo "运行测试..."
go test ./...

# 构建应用
echo "构建应用..."
go build -o build/risk-insight-system main.go

# 检查构建结果
if [ $? -eq 0 ]; then
    echo "构建成功！"
    echo "可执行文件位置: build/risk-insight-system"
    ls -la build/
else
    echo "构建失败！"
    exit 1
fi 