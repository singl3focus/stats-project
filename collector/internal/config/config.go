package config

import (
	"flag"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	DefaultConfigPath = "./config.yaml"
)

type Config struct {
	App        AppConfig    `yaml:"app"`
	GRPCServer GRPCConfig   `yaml:"grpc"`
	Logger     LoggerConfig `yaml:"logger"`
	Database   DBConfig     `yaml:"database"`
}

type AppConfig struct {
	Version string `yaml:"version" env-requried:"true"`
}

type GRPCConfig struct {
	Port              int           `yaml:"port" env-requried:"true"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env-requried:"true"`
	ReadTimeout       time.Duration `yaml:"read_timeout" env-requried:"true"`
	WriteTimeout      time.Duration `yaml:"write_timeout" env-requried:"true"`
	IdleTimeout       time.Duration `yaml:"idle_timeout" env-requried:"true"`
}

type DBConfig struct {
	Migration struct {
		Dir string `yaml:"dir" env-requried:"true"`
	} `yaml:"migration" env-requried:"true"`
	Postgres struct {
		URL string `yaml:"url" env-requried:"true"`
	} `yaml:"postgres" env-requried:"true"`
}

type LoggerConfig struct {
	Enable bool   `yaml:"enable" env-default:"true"`
	Level  string `yaml:"level" env-default:"INFO"` // Avaliable: DEBUG, INFO, WARN, ERROR
}

// MustLoadCfg.
// fetchConfigPath fetches config path from command line flag or set default variable.
// Priority: flag > default.
// Default value is local config path.
func MustLoadCfg() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		configPath = DefaultConfigPath
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var p string
	flag.StringVar(&p, "cfg", "", "path to config")

	flag.Parse()

	return p
}
