package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	from       string
	to         string
	cfgFile    string
	console    bool
	csvFlag    bool
	verbose    bool
	grafanaKey string

	rootCmd = &cobra.Command{
		Use:   "deployreporter",
		Short: "Retrieve and view your deployments",
		Long: `Tired of gathering deployment information from difference sources?
		deployreporter collects and aggregate deployment data to quickly find what has changed in your service`,
	}
)

func main() {
	Execute()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&from, "from", "", "epoch datetime in milliseconds")
	rootCmd.PersistentFlags().StringVar(&to, "to", "", "epoch datetime in milliseconds")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deployreporter.yaml)")
	rootCmd.PersistentFlags().BoolVar(&console, "console", false, "output results to console")
	rootCmd.PersistentFlags().BoolVar(&csvFlag, "csv", false, "output results to a csv file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enables verbose output")
	rootCmd.PersistentFlags().StringVar(&grafanaKey, "grafanaKey", "", "API key used to retrieve data from grafana")

	err := viper.BindPFlag("from", rootCmd.PersistentFlags().Lookup("from"))
	if err != nil {
		fmt.Printf("error binding 'from' flag. %v", err)
	}
	err = viper.BindPFlag("to", rootCmd.PersistentFlags().Lookup("to"))
	if err != nil {
		fmt.Printf("error binding 'to' flag. %v", err)
	}
	err = viper.BindPFlag("console", rootCmd.PersistentFlags().Lookup("console"))
	if err != nil {
		fmt.Printf("error binding 'console' flag. %v", err)
	}
	err = viper.BindPFlag("csv", rootCmd.PersistentFlags().Lookup("csv"))
	if err != nil {
		fmt.Printf("error binding 'csv' flag. %v", err)
	}
	err = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	if err != nil {
		fmt.Printf("error binding 'verbose' flag. %v", err)
	}
	err = viper.BindPFlag("grafanaKey", rootCmd.PersistentFlags().Lookup("grafanaKey"))
	if err != nil {
		fmt.Printf("error binding 'grafanaKey' flag. %v", err)
	}

	viper.SetDefault("console", false)
	viper.SetDefault("csv", false)
	viper.SetDefault("verbose", false)
	viper.SetDefault("grafanaKey", "")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".deploy-reporter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".deployreporter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		verbose := viper.GetBool("verbose")
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
			fmt.Println("Configuration retrieved:")
			for key, value := range viper.AllSettings() {
				fmt.Printf("%v: %v\n", key, value)
			}
		}
	}
}
