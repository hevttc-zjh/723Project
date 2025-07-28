package config

import (
	"log"

	"github.com/spf13/viper"
)

// Init 初始化配置
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置默认值
	setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("警告: 无法读取配置文件，使用默认配置: %v", err)
	}

	// 读取环境变量
	viper.AutomaticEnv()

	return nil
}

// setDefaults 设置默认配置值
func setDefaults() {
	// 服务器配置
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")

	// 数据库配置
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.name", "risk_insight")
	viper.SetDefault("database.sslmode", "disable")

	// JWT配置
	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expire_hours", 24)

	// 日志配置
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", "logs/app.log")
}

// GetString 获取字符串配置
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt 获取整数配置
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool 获取布尔配置
func GetBool(key string) bool {
	return viper.GetBool(key)
}
