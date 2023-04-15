package conf

import "sync"

// config 生命全局配置
var config *Config

func getConfig() *Config {
	return config
}

// Config 配置
type Config struct {
	App   *app   `toml:"app"`
	MySQL *MySQL `toml:"mysql"`
}

// app config
type app struct {
	Name      string `toml:"name" env:"APP_NAME"`
	Host      string `toml:"host" env:"APP_HOST"`
	Port      string `toml:"port" env:"APP_PORT"`
	Key       string `toml:"key" env:"APP_KEY"`
	EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

// MySQL mysql config
type MySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}

// NewDefaultConfig new default config
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		MySQL: NewDefaultMySQL(),
	}
}

// NewDefaultApp new default app config
func NewDefaultApp() *app {
	return &app{
		Name: "rest-api-demo",
		Host: "127.0.0.1",
		Port: "8888",
	}
}

// NewDefaultMySQL new default mysql config
func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Database:    "go_rest_api_demo",
		MaxOpenConn: 200,
		MaxIdleConn: 50,
		MaxLifeTime: 180,
		MaxIdleTime: 180,
	}
}
