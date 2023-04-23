package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "go-cli-template",
	Short: "A CLI template app",
	Long:  `Fork and modify to quickly bootstrap a CLI app`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initLogging)

	// Define global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/...yaml)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug logging")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".." (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("go-cli-template")
		viper.AddConfigPath("$HOME/.config")
	}

	//viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("debug", rootCmd.Flags().Lookup("debug"))

	// Read in environment variables that match
	viper.SetEnvPrefix("GO_CLI_TEMPLATE")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initLogging() {
	var config zap.Config

	if viper.GetBool("debug") {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()

	if err != nil {
		log.Fatalf("Can't initialise zap logger: %v", err)
	}
	defer logger.Sync()

	// Make this logger global
	zap.ReplaceGlobals(logger)
}
