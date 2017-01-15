package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	previewFlag  bool
	durationFlag string
)

const (
	limitOrder = "limit"
	stopOrder  = "stop"
)

func newOrderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order [ID]",
		Short: "Get a particular order information for an account",
		Long:  `Get a particular order information for an account`,
		RunE:  orderCmdFunc,
	}

	return cmd
}

func orderCmdFunc(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No order ID provided\n")
	}

	output, err := brokrRunner.GetOrder(args[0])
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
