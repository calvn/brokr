package config

import "fmt"

type Config struct {
	Brokerage   string
	AccessToken string
	FilePath    string
}

func New() *Config {
	config := &Config{}

	return config
}

func (c *Config) checkConfig() error {
	if len(c.AccessToken) == 0 {
		return fmt.Errorf("config error: access token not provided in config file")
	}

	return nil
}
