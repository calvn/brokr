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

	"github.com/calvn/brokr/brokr"
	"github.com/calvn/brokr/buildtime"
	"github.com/calvn/brokr/config"
	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command

var printVersion bool

var brokrRunner *brokr.Runner
var mergedConfig *config.Config

func newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "brokr",
		Short: "brokr - bringing your trades into the console",
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
	if !printVersion && len(args) == 0 {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	if printVersion {
		fmt.Println("brokr:")
		fmt.Printf("  %-9s%s\n", "Version:", buildtime.Version)
		if buildtime.GitCommit != "" {
			fmt.Printf("  %-9s%s\n", "Build:", buildtime.GitCommit)
		}
	}
}

func init() {
	// This gets run after all init()'s, but before any commands'
	cobra.OnInitialize(initConfig, setConfig, initRunner)

	RootCmd = newRootCommand()
	RootCmd.AddCommand(newConfigCmd())
	RootCmd.AddCommand(newQuoteCmd())
	RootCmd.AddCommand(newBuyCmd())
	RootCmd.AddCommand(newInfoCmd())
	RootCmd.AddCommand(newVersionCmd())
}

// initClient instantiates a new brokr client
func initRunner() {
	brokrRunner = brokr.NewRunner(mergedConfig)
}
