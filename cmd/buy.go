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

var previewFlag bool

func newBuyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy",
		Short: "Preview  or place a buy order",
		Long:  `Preview  or place a buy order`,
		Run:   buyCmdFunc,
	}
	cmd.Flags().BoolVarP(&previewFlag, "preview", "p", true, "Preview order, default: true")
	viper.BindPFlag("preview_order", cmd.Flags().Lookup("preview"))

	return cmd
}

func buyCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("this is the buy command")
}
