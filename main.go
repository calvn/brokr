package main

import (
	"os"

	"github.com/calvn/brokr/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
