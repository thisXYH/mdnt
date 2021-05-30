package cmd

import "github.com/spf13/cobra"

const VERSION = "v1.1.0"

var rootCmd = &cobra.Command{
	Use:     "nt",
	Version: VERSION,
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
