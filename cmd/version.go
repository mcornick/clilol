package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	goVersion "go.hein.dev/go-version"
)

var (
	shortened  = false
	output     = "json"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show clilol's version",
		Long:  "Prints the version and build information for clilol.",
		Run: func(_ *cobra.Command, _ []string) {
			resp := goVersion.FuncWithOutput(shortened, version, commit, date, output)
			fmt.Print(resp)
			return
		},
	}
)

func init() {
	versionCmd.Flags().BoolVarP(&shortened, "short", "s", false, "Print just the version number.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "json", "Output format. One of 'yaml' or 'json'.")
	rootCmd.AddCommand(versionCmd)
}
