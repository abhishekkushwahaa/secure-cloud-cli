package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secure-cloud",
	Short: "Secure Cloud CLI",
	Long:  "A CLI tool for managing Secure Cloud resources.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(register)
	rootCmd.AddCommand(login)
	rootCmd.AddCommand(upload)
	rootCmd.AddCommand(download)
}
