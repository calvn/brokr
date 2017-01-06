package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the merged configuration from the config file, environment variables, and flags
type Config struct {
	Brokerage    string `yaml:"brokerage"`
	PreviewOrder bool   `yaml:"preview_order"`

	Tradier *TradierConfig `yaml:"tradier,omitempty"`
}

// TradierConfig holds the configuration for Tradier brokerage
type TradierConfig struct {
	AccountID   string `yaml:"account_id,omitempty"`
	AccessToken string `yaml:"access_token"`
}

// New creates a new config from Viper object
func New(v *viper.Viper) *Config {
	config := &Config{
		Brokerage:    v.GetString("brokerage"),
		PreviewOrder: v.GetBool("preview_order"),
	}

	// Handle tradier config if it exists
	if v.GetStringMap("tradier") != nil {
		tc := &TradierConfig{
			AccessToken: v.GetString("tradier.access_token"),
		}

		config.Tradier = tc
	}

	return config
}

// CheckConfig checks for the validity of the cofiguration
func (c *Config) CheckConfig() error {
	if len(c.Brokerage) == 0 {
		return fmt.Errorf("config error: brokerage not provided")
	}

	if len(c.Tradier.AccessToken) == 0 {
		return fmt.Errorf("config error: access token not provided")
	}

	return nil
}
