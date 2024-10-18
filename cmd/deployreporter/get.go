package main

import (
	deployreporter "github.com/sebasrp/deployreporter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(get)
}

var get = &cobra.Command{
	Use:   "get",
	Short: "Retrieves deployments",
	Long:  `Retrieves deployments`,
	Run: func(cmd *cobra.Command, args []string) {
		grafanaKey := viper.GetString("grafanaKey")

		deployreporter.GetDeployments(grafanaKey)
	},
}
