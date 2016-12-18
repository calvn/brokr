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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information about brokr settings",
	Long:  `Show information about brokr settings`,
	Run: func(cmd *cobra.Command, args []string) {

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Config file location:", viper.ConfigFileUsed())
		}
	},
}
