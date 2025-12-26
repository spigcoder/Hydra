package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/spf13/viper"
)

// Init 通用加载逻辑
func Init(configPath string, configName string, envPrefix string) error {
	// 1. 设置配置文件路径和名称
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(envPrefix)

	// 2. 加载 .env 文件到系统环境变量 (手动处理，为了解决 key 映射问题)
	_ = godotenv.Load(filepath.Join(configPath, ".env"))

	// 3. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果只是配置文件没找到，但我们可能只想用环境变量，可以容忍
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("读取配置文件失败: %w", err)
		}
	}

	// 4. 开启环境变量自动覆盖
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return nil
}

func GetViper(configPath, configName string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	return v, nil
}
