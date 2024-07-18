package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

type ServerConfig struct {
	Host string `yaml:"host" envconfig:"SERVER_HOST"`
	Port string `yaml:"port" envconfig:"SERVER_PORT"`
}

type LoggerConfig struct {
	Level string `yaml:"level" envconfig:"LOG_LEVEL"`
}

// Options to
type OptionsNewConfig struct {
	FilePath string
}

// NewConfig returns configuration data for the application. The order of precedence for
// configuration is `environment variables > file > defaults`
func NewConfig(o OptionsNewConfig) (Config, error) {
	applyOptionsDefaults(&o)

	cfg := getDefault()
	// config from file will overwrite the default configuration
	err := setFromConfigFile(o.FilePath, &cfg)
	if err != nil {
		return cfg, err
	}
	// config from environment variables will overwrite the default config and file-based config
	err = setFromEnv(&cfg)

	return cfg, err
}

func applyOptionsDefaults(o *OptionsNewConfig) {
	var empty string
	if o.FilePath == empty {
		o.FilePath = "config.yml"
	}
}

func getDefault() Config {
	return Config{
		Server: ServerConfig{
			Host: "localhost",
			Port: "3000",
		},
		Logger: LoggerConfig{
			Level: "debug",
		},
	}
}

func setFromConfigFile(path string, cfg *Config) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, cfg)
	return err
}

func setFromEnv(cfg *Config) error {
	err := envconfig.Process("", cfg)
	return err
}
