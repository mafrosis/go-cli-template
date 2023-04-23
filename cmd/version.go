package cmd

import (
	"fmt"

	"github.com/mafrosis/go-cli-template/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (%s) built on %s\n", version.String, version.Revision, version.Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
