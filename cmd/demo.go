package cmd

import (
	"github.com/mafrosis/go-cli-template/pkg/demo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "A demo command",
	Run: func(cmd *cobra.Command, args []string) {
		demo.Demo(viper.GetString("example"))
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)
}
