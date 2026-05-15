package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "evcc",
	Short: "EV Charge Controller",
	Long: `evcc is a smart EV charge controller that optimizes charging
based on solar production, grid pricing, and vehicle state.`,
	RunE: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default: evcc.yaml)")
	rootCmd.PersistentFlags().String("log", "info", "log level (trace, debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().String("uri", "0.0.0.0:7070", "listen address")

	// Bind flags to viper
	_ = viper.BindPFlag("log", rootCmd.PersistentFlags().Lookup("log"))
	_ = viper.BindPFlag("uri", rootCmd.PersistentFlags().Lookup("uri"))

	// Version command
	rootCmd.AddCommand(versionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in current directory and home directory
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigName("evcc")
		viper.SetConfigType("yaml")
	}

	// Read environment variables with EVCC_ prefix
	viper.SetEnvPrefix("EVCC")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// run is the main entry point for the evcc command
func run(cmd *cobra.Command, args []string) error {
	// TODO: initialize and start the main application
	fmt.Printf("evcc %s (%s) built %s\n", version, commit, date)
	return nil
}

// versionCmd prints version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("evcc %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Built:  %s\n", date)
	},
}
