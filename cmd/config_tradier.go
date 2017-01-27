// Copyright © 2017 Calvin Leung Huang <https://github.com/calvn/brokr>
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
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/calvn/brokr/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	accessTokenFlag string
	accountFlag     string
)

func newConfigTradierCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tradier",
		Short: "Configure tradier settings",
		Long:  `Configure tradier settings`,
		RunE:  configTradierCmdFunc,
	}

	cmd.Flags().StringVarP(&accessTokenFlag, "token", "t", "", "Tradier access token, required if not set")
	viper.BindPFlag("tradier.access_token", cmd.Flags().Lookup("token"))

	cmd.Flags().StringVarP(&accountFlag, "account", "a", "", "Tradier access token, required if not set")
	viper.BindPFlag("tradier.account", cmd.Flags().Lookup("account"))

	return cmd
}

func configTradierCmdFunc(cmd *cobra.Command, args []string) error {
	// Make sure flags are provided if not found in config
	if viper.GetString("tradier.account") == "" || viper.GetString("tradier.access_token") == "" {
		log.Println(viper.Get("tradier.account"))
		log.Println(viper.Get("tradier.access_token"))

		cmd.SilenceUsage = true
		return fmt.Errorf("Tradier configuration missing account ID or access token")
	}

	// Instantiate a new *config.Config from viper config
	// This includes configuration from existing viper data
	cfg := config.New(viper.GetViper())

	// Marshal config into YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(home, defaultConfigName)

	// Write config to file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Configuration written to %s\n", filePath)
	return nil
}
