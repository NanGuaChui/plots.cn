package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Upload   UploadConfig   `yaml:"upload"`
	Log      LogConfig      `yaml:"log"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"` // debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Path string `yaml:"path"`
}

// JWTConfig JWT 配置
type JWTConfig struct {
	Secret string        `yaml:"secret"`
	Expire time.Duration `yaml:"expire"`
}

// UploadConfig 上传配置
type UploadConfig struct {
	MaxSize      int64    `yaml:"max_size"`
	AllowedTypes []string `yaml:"allowed_types"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

// Load 加载配置文件
func Load() (*Config, error) {
	// 默认配置
	cfg := &Config{
		Server: ServerConfig{
			Port: "8080",
			Mode: "release",
		},
		Database: DatabaseConfig{
			Path: "./data/plots.db",
		},
		JWT: JWTConfig{
			Secret: "default-secret-please-change",
			Expire: 72 * time.Hour,
		},
		Upload: UploadConfig{
			MaxSize:      10 * 1024 * 1024, // 10MB
			AllowedTypes: []string{"image/jpeg", "image/png", "image/gif", "image/webp"},
		},
		Log: LogConfig{
			Level: "info",
			Path:  "./logs/app.log",
		},
	}

	// 从环境变量获取配置文件路径
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs/config.yaml"
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		// 配置文件不存在时使用默认配置
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	// 解析配置文件
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	// 从环境变量覆盖配置
	if port := os.Getenv("PORT"); port != "" {
		cfg.Server.Port = port
	}
	if dbPath := os.Getenv("DB_PATH"); dbPath != "" {
		cfg.Database.Path = dbPath
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		cfg.JWT.Secret = jwtSecret
	}

	return cfg, nil
}
