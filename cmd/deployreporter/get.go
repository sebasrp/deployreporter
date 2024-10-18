package main

import (
	"fmt"

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

		deployments := deployreporter.GetDeployments(grafanaKey)
		if viper.GetBool("verbose") {
			for _, d := range deployments {
				fmt.Printf("%+v\n", d)
			}
		}
	},
}
