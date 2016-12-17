package config

import "fmt"

// Config holds the merged configuration from the config file, environment variables, and flags
type Config struct {
	AccessToken string `yaml:"access_token"`
}

func New(token string) *Config {
	config := &Config{
		AccessToken: token,
	}

	// NOTE: Should logic be performed in there , or handled upstream?
	// if err := config.checkConfig(); err != nil {
	// 	return nil
	// }

	return config
}

func (c *Config) checkConfig() error {
	if len(c.AccessToken) == 0 {
		return fmt.Errorf("config error: access token not provided in config file")
	}

	return nil
}
