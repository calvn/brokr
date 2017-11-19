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
		RunE:    sellCmdFunc,
	}
	cmd.Flags().BoolVarP(&previewFlag, "preview", "p", false, "Preview order, overwrites the setting from the config")
	cmd.Flags().StringVarP(&durationFlag, "duration", "d", "day", "Duration of the order")

	return cmd
}

func sellCmdFunc(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Invalid buy command")
	}

	// Quantity
	q, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	// Symbol
	symbol := args[1]
	if len(symbol) == 0 {
		return fmt.Errorf("Cannot provide empty symbol")
	}

	// Type and trigger price
	orderType := "market"
	triggerPrice := 0.0
	if len(args) == 4 {
		switch {
		case args[2] == limitOrder || args[2] == "l":
			orderType = limitOrder
		case args[2] == stopOrder:
			orderType = stopOrder
		}

		triggerPrice, _ = strconv.ParseFloat(args[3], 64)
	}

	isPreview := viper.GetBool("preview_order") || previewFlag

	output, err := brokrRunner.CreateOrder(isPreview, "equity", symbol, durationFlag, "sell", q, orderType, triggerPrice)
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
