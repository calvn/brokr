package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the merged configuration from the config file, environment variables, and flags
type Config struct {
	// General configuration
	PreviewOrder bool `yaml:"preview_order"`

	// Account configuration
	Brokerage string         `yaml:"brokerage"`
	Tradier   *TradierConfig `yaml:"tradier,omitempty"`
}

// TradierConfig holds the configuration for Tradier brokerage
type TradierConfig struct {
	AccountID   string `yaml:"account,omitempty"`
	AccessToken string `yaml:"access_token,omitempty"`
}

// New creates a new config from Viper object
func New(v *viper.Viper) *Config {
	brokerage := v.GetString("brokerage")

	config := &Config{
		Brokerage:    brokerage,
		PreviewOrder: v.GetBool("preview_order"),
	}

	switch brokerage {
	case "tradier":
		config.Tradier = &TradierConfig{
			AccountID:   v.GetString("tradier.account"),
			AccessToken: v.GetString("tradier.access_token"),
		}
	}

	// Handle tradier config if it exists
	// if len(v.GetStringMap("tradier")) != 0 {
	// 	tc := &TradierConfig{
	// 		AccountID:   v.GetString("tradier.account_id"),
	// 		AccessToken: v.GetString("tradier.access_token"),
	// 	}
	//
	// 	config.Tradier = tc
	// }

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

// Copy returns a copy of the configuration
func (c *Config) Copy() *Config {
	if c == nil {
		return nil
	}
	cp := *c
	return &cp
}
