package config

import (
	"github.com/spf13/viper"
)

// Build information -ldflags .
var (
	buildVersion string = "dev"
	buildCommit  string = "-"
	buildDate    string = "-"
)

var cfg *Config

// Config - contains all configuration parameters in config package.
type Config struct {
	Project Project `yaml:"project"`
}

// Project - contains all parameters project information.
type Project struct {
	Name  string `yaml:"name"`
	Build Build  `yaml:"build"`
}

// Build - contains all build information
type Build struct {
	Version string `yaml:"version"`
	Commit  string `yaml:"commit"`
	Date    string `yaml:"date"`
}

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	cfg.Project.Build = Build{
		Version: buildVersion,
		Commit:  buildCommit,
		Date:    buildDate,
	}

	return nil
}
