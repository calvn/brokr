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
	"os"
	"path/filepath"
	"strings"

	"github.com/calvn/brokr/brokr"
	"github.com/calvn/brokr/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The git commit that will be used to describe the version
var (
	GitCommit string

	// Version of the program
	Version = "0.0.1"
)

var RootCmd *cobra.Command

var printVersion bool

var brokrRunner *brokr.Runner
var mergedConfig *config.Config

func newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "brokr",
		Short: "brokr - bringing your trades into the console.",
		Long: `brokr - bringing your trades into the console.
  It currently supports making trades against Tradier.

Made with ♥︎ in Golang.`,
		Run: rootCmdRunFunc,
	}

	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "print version and exit")
	return rootCmd
}

func rootCmdRunFunc(cmd *cobra.Command, args []string) {
	// If no commands passed in, display help
	if len(args) == 0 || !cmd.HasFlags() {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	if printVersion {
		fmt.Println("brokr:")
		fmt.Printf("  %-9s%s\n", "Version:", Version)
		if GitCommit != "" {
			fmt.Printf("  %-9s%s\n", "Build:", GitCommit)
		}
	}
}

func init() {
	// This gets run after all init()'s, but before any commands'
	// NOTE: config should come before runner
	cobra.OnInitialize(initConfig, setConfig)

	RootCmd = newRootCommand()
	RootCmd.AddCommand(newConfigCmd())
	RootCmd.AddCommand(quoteCmd)
	RootCmd.AddCommand(infoCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configName := strings.TrimSuffix(config.DefaultConfigName, filepath.Ext(config.DefaultConfigName))

	viper.SetConfigName(configName)               // name of config file (without extension)
	viper.AddConfigPath(config.DefaultConfigPath) // adding home directory as first search path
	viper.AutomaticEnv()                          // read in environment variables that match
}

// setConfig sets config variables from viper variables
func setConfig() {
	if t := viper.GetString("access_token"); t != "" {
		config.AccessTokenFlag = t
	}
}

// initClient instantiates a new brokr client
func initRunner() {
	brokrRunner = brokr.NewRunner(mergedConfig)
}
