package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/calvn/brokr/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	fmt.Println("init from cli.go")
	cobra.OnInitialize(initConfig, setConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configName := strings.TrimSuffix(config.DefaultConfigName, filepath.Ext(config.DefaultConfigName))

	viper.SetConfigName(configName)               // name of config file (without extension)
	viper.AddConfigPath(config.DefaultConfigPath) // adding home directory as first search path
	viper.AutomaticEnv()                          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// setConfig sets config variables from viper variables
func setConfig() {
	if t := viper.GetString("access_token"); t != "" {
		fmt.Println("not empty")
		config.AccessToken = t
		fmt.Println(config.AccessToken)
	}
}
