package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a data file against a flow schema yaml",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
