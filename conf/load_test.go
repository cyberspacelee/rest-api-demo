package conf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	assertions := assert.New(t)
	err := LoadConfigFromToml("../etc/config.toml")
	if assertions.NoError(err) {
		cfg := getConfig()
		fmt.Printf("app config: %v\n", cfg.App)
		fmt.Printf("database config: %v\n", cfg.MySQL)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	assertions := assert.New(t)
	os.Setenv("APP_NAME", "rest-api-demo-test")
	err := LoadConfigFromEnv()
	if assertions.NoError(err) {
		cfg := getConfig()
		fmt.Printf("app config: %v\n", cfg.App)
		fmt.Printf("database config: %v\n", cfg.MySQL)
	}
}
