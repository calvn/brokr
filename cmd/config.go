// Copyright © 2016 Calvin Leung Huang <https://github.com/calvn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/calvn/brokr/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var accessToken string

const (
	defaultConfigPath = "$HOME"
	defaultConfigName = ".brokr.yaml"
)

// configCmd represents the config command
func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "config",
		Short:         "Configure .brokr.yaml",
		Long:          `Configure .brokr.yaml`,
		Run:           configCmdFunc,
		PreRunE:       configPreRunEFunc,
		SilenceErrors: true,
	}
	cmd.Flags().StringVarP(&accessToken, "token", "t", "", "Tradier access token, required")
	viper.BindPFlag("access_token", cmd.Flags().Lookup("token"))

	return cmd
}

// TODO: If config file exist, merge with it
func configCmdFunc(cmd *cobra.Command, args []string) {
	t := viper.GetString("access_token")

	cfg := config.New(viper.GetViper())

	data, err := yaml.Marshal(cfg)
	if err != nil {
		fmt.Println(err)
	}

	home, err := homedir.Dir()
	if err != nil {
		return
	}

	filePath := filepath.Join(home, defaultConfigName)

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Configuration written to %s\n", filePath)
}

// Check for required flags, reads from viper
func configPreRunEFunc(cmd *cobra.Command, args []string) error {
	t := viper.GetString("access_token")
	mergedConfig.AccessToken = t

	if mergedConfig.AccessToken == "" {
		return fmt.Errorf("No access token provided.")
	}

	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configName := strings.TrimSuffix(defaultConfigName, filepath.Ext(defaultConfigName))

	viper.SetConfigName(configName)        // name of config file (without extension)
	viper.AddConfigPath(defaultConfigPath) // adding home directory as first search path
	viper.ReadInConfig()                   // read in config
	viper.SetEnvPrefix("brokr")            // set env prefix
	viper.AutomaticEnv()                   // read in environment variables that match
}

// setConfig reads config from viper and instantiates mergedConfig, used for proceeding commands
func setConfig() {
	t := viper.GetString("access_token")

	mergedConfig = config.New(t)
}
