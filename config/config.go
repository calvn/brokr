package config

import "fmt"

var (
	DefaultConfigPath = "$HOME"
	DefaultConfigName = ".brokr.yaml"
)

type Config struct {
	AccessToken string `yaml:"access_token"`
}

func New(token string) *Config {
	config := &Config{
		AccessToken: AccessToken,
	}

	if err := config.checkConfig(); err != nil {
		return nil
	}

	return config
}

func (c *Config) checkConfig() error {
	if len(c.AccessToken) == 0 {
		return fmt.Errorf("config error: access token not provided in config file")
	}

	return nil
}
