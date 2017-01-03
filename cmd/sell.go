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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newSellCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sell QUANTITY SYMBOL",
		Aliases: []string{"s"},
		Short:   "Preview or place a sell order",
		Long:    `Preview or place a sell order`,
		Run:     sellCmdFunc,
	}
	cmd.Flags().BoolVarP(&previewFlag, "preview", "p", true, "Preview order, default: true")
	cmd.Flags().StringVarP(&durationFlag, "duration", "d", "day", "Duration of the order, default: day")
	viper.BindPFlag("preview_order", cmd.Flags().Lookup("preview"))

	return cmd
}

func sellCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Invalid buy command")
		return
	}

	// Quantity
	q, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Symbol
	symbol := args[1]
	if len(symbol) == 0 {
		fmt.Println("Cannot provide empty symbol")
	}

	// Type and trigger price
	orderType := "market"
	triggerPrice := 0.0
	if len(args) == 4 {
		switch args[2] {
		case "limit":
			orderType = "limit"
		case "stop":
			orderType = "stop"
		}

		triggerPrice, _ = strconv.ParseFloat(args[3], 64)
	}

	output, err := brokrRunner.PlaceOrder("equity", symbol, durationFlag, "buy", q, orderType, triggerPrice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Order IDs: \n%s\n", output)
}
