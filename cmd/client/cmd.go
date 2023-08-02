package main

import (
	"github.com/juicity/juicity/config"
	"github.com/spf13/cobra"

	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:     "juicity-client [flags] [command [argument ...]]",
		Short:   "juicity-client is a quic-based proxy client.",
		Long:    "juicity-client is a quic-based proxy client.",
		Version: config.Version,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: false,
		},
	}
)

// Execute executes the root command.
func Execute() error {
	os.Setenv("QUIC_GO_ENABLE_GSO", "true")
	return rootCmd.Execute()
}
