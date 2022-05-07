package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "lws",
	Short:   "LWS Client Command",
	Long:    "Main Entry Point of LWS client",
	Example: "lws register",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(registerCommand)
}

func Execute() error {
	return rootCmd.Execute()
}
