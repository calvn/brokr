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

	"gopkg.in/yaml.v2"

	"github.com/calvn/brokr/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure .brokr.yaml",
	Long:  `Configure .brokr.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New(config.AccessToken)
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

		filePath := filepath.Join(home, config.DefaultConfigName)

		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Configuration written to %s\n", filePath)
	},
}

func init() {
	fmt.Println("init from config.go")
	configCmd.Flags().StringVarP(&config.AccessToken, "token", "t", "", "Access token obtained from Tradier")
	viper.BindPFlag("access_token", configCmd.Flags().Lookup("token"))

	// fmt.Println("from config file: ", viper.GetString("access_token"))
	// fmt.Println("from config struct:", config.AccessToken)

	err := configCmd.MarkFlagRequired("token")
	if err != nil {
		fmt.Println(err)
	}
}

func checkInit() {
	if viper.ConfigFileUsed() == "" {
		fmt.Println("[Warning] No config file found. Use `brokr init` to generate one.")
	}
}
