package main

import (
	"fmt"
	"os"

	"github.com/calvn/brokr/cmd"
	"github.com/calvn/brokr/config"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// If no commands passed in, display help
	if len(os.Args) == 1 {
		if err := cmd.RootCmd.Help(); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	if config.PrintVersion {
		fmt.Println("brokr:")
		fmt.Printf("  %-9s%s\n", "Version:", Version)
		if GitCommit != "" {
			fmt.Printf("  %-9s%s\n", "Build:", GitCommit)
		}
	}
}
