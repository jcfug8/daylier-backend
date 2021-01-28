package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "apictl",
	Short: "apictl",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := viper.GetString("log")
		logLevel, err := log.ParseLevel(level)
		if err != nil {
			panic(fmt.Errorf("failed to parse log level: %+v", err))
		}

		log.SetLevel(logLevel)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.WithError(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().String("log", "info", "set log level (default: info)")

	viper.BindPFlag("log", RootCmd.PersistentFlags().Lookup("log"))
}
