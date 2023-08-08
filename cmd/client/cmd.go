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
			DisableDefaultCmd: true,
		},
	}
)

func MyExecute(defCmd string) error {
	var cmdFound bool
	cmd := rootCmd.Commands()

	for _, a := range cmd {
		for _, b := range os.Args[1:] {
			if a.Name() == b {
				cmdFound = true
				break
			}
		}
	}
	if cmdFound == false {
		args := append([]string{defCmd}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
	return rootCmd.Execute()
}

// Execute executes the root command.
func Execute() error {
	os.Setenv("QUIC_GO_ENABLE_GSO", "true")
	return MyExecute("run")
	// return rootCmd.Execute("run")
}
