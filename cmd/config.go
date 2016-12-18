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
		Use:   "config",
		Short: "Configure .brokr.yaml",
		Long:  `Configure .brokr.yaml`,
		Run:   configCmdFunc,
	}
	cmd.Flags().StringVarP(&accessToken, "token", "t", "", "Access token obtained from Tradier")
	viper.BindPFlag("access_token", cmd.Flags().Lookup("token"))

	// FIXME: Doesn't work?
	err := cmd.MarkFlagRequired("token")
	if err != nil {
		fmt.Println(err)
	}

	return cmd
}

func configCmdFunc(cmd *cobra.Command, args []string) {
	cfg := config.New(accessToken)
	if cfg == nil {
		fmt.Println("Access token not provided")
		return
	}

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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configName := strings.TrimSuffix(defaultConfigName, filepath.Ext(defaultConfigName))

	viper.SetConfigName(configName)        // name of config file (without extension)
	viper.AddConfigPath(defaultConfigPath) // adding home directory as first search path
	viper.ReadInConfig()                   // read in config
	viper.AutomaticEnv()                   // read in environment variables that match
}

// setConfig sets config variables from viper variables
func setConfig() {
	t := viper.GetString("access_token")

	mergedConfig = config.New(t)
	if mergedConfig.AccessToken == "" {
		fmt.Println("[Warning] No access token provided")
	}
}