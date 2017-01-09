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
)

func newOrdersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders",
		Short: "Get pending orders for an account",
		Long:  `Get pending orders for an account`,
		RunE:  ordersCmdFunc,
	}

	return cmd
}

func ordersCmdFunc(cmd *cobra.Command, args []string) error {
	output, err := brokrRunner.GetOrders()
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
