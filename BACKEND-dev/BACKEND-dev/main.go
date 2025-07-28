package main

import (
	"log"
	"os"

	"risk-insight-system/config"
	"risk-insight-system/internal/router"
	"risk-insight-system/internal/server"
)

func main() {
	// 打开日志文件
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer logFile.Close()

	// 设置日志输出到文件
	log.SetOutput(logFile)

	// 设置日志格式，包含时间和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置初始化失敗: %v", err)
	}

	// 初始化資料庫連接
	if err := server.InitDB(); err != nil {
		log.Fatalf("資料庫連接失敗: %v", err)
	}

	// 使用 router 包初始化路由
	r := router.InitRouter()

	// 启动服务器
	port := config.GetString("server.port")
	log.Printf("服务器启动在端口: %s", port)
	log.Printf("健康检查: http://localhost:%s/health", port)
	log.Printf("API文档: http://localhost:%s/api/v1", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
