package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/lucasjones/reggen"
	"github.com/spf13/cobra"
	"github.com/wallaroo/flowgenerator/flow"
)

var stringGenerator *reggen.Generator
var numberGenerator *reggen.Generator

//flags
var rowNumber int

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

var rootCmd = &cobra.Command{
	Use:   "flowgenerator <yaml config file> [destinationfile]",
	Short: "A tool to generate file flows based on flow.yaml descriptor",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if Exists(args[0]) {
			var dest string
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal("Error", err)
			}

			if len(args) > 1 {
				dest = path.Join(wd, args[1])
			} else {
				dest = path.Join(wd, strings.Replace(args[0], ".yaml", ".dat", 1))
			}
			err = flow.Create(args[0], dest, rowNumber)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("src file", args[0], "doesn't exists")
		}
	},
}

// Execute executes commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&rowNumber, "rows", "r", 10, "number of content rows")
}
