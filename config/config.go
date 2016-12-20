package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the merged configuration from the config file, environment variables, and flags
type Config struct {
	Brokerage   string `yaml:"brokerage"`
	AccessToken string `yaml:"access_token"`
}

// Create a new config from Viper object
func New(v *viper.Viper) *Config {
	config := &Config{
		AccessToken: v.GetString("access_token"),
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
