package cmd

import (
	"fmt"

	"github.com/calvn/brokr/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd *cobra.Command

func createRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "brokr",
		Short: "brokr brings your trades into the console",
		Long: `brokr let's you place trades via the CLI.
  It currently supports making trades against Tradier.

Made with ♥︎ in Golang.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.Flags().BoolVarP(&config.PrintVersion, "version", "v", false, "print version and exit")
	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd = createRootCommand()
	RootCmd.AddCommand(initCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".brokr") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
