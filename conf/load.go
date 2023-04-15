package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadConfigFromToml(filePath string) error {
	cfg := NewDefaultConfig()
	// 读取 toml 格式配置
	_, err := toml.DecodeFile(filePath, cfg)
	if err != nil {
		return fmt.Errorf("load config from toml file error, path: %s, error: %s", filePath, err)
	}
	config = cfg
	return nil
}

func LoadConfigFromEnv() error {
	cfg := NewDefaultConfig()
	// 从环境变量中读取配置
	err := env.Parse(cfg)
	if err != nil {
		return fmt.Errorf("load config from env error: %s", err)
	}
	config = cfg
	return nil
}
