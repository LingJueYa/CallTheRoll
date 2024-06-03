package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config 配置文件的内容
type Config struct {
	Log      Log      `yaml:"log"`
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

// Server 服务器配置
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Addr 返回格式化的服务器地址
func (s Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Log 日志配置
type Log struct {
	Output    []string `yaml:"output"`
	ErrOutput []string `yaml:"errOutput"`
	Level     string   `yaml:"level"`
}

// Database 数据库配置
type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

// AppConfig 定义全局变量，存储配置
var AppConfig Config

// G 获取全局配置实例，修正函数返回类型
func G() *Config {
	return &AppConfig
}

// LoadConfig 加载和解析配置
func LoadConfig() {
	viper.SetConfigName("config")   // 设置配置文件名为config.yaml
	viper.SetConfigType("yaml")     // 设置配置文件类型为YAML
	viper.AddConfigPath("./config") // 添加配置文件搜索路径

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err)
	}

	// 解析到AppConfig结构体
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}
