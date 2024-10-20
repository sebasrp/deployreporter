package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

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
		csvFlag := viper.GetBool("csv")
		from := viper.GetString("from")
		to := viper.GetString("to")
		limit := viper.GetInt("limit")

		deployments := deployreporter.GetDeployments(from, to, limit, grafanaKey)
		if viper.GetBool("verbose") {
			for _, d := range deployments {
				fmt.Printf("%+v\n", d)
			}
		}

		if csvFlag {
			csvfile, err := os.Create("deploychecker.csv")
			if err != nil {
				fmt.Printf("failed creating file: %s", err)
			}
			csvwriter := csv.NewWriter(csvfile)

			_ = csvwriter.Write([]string{"id", "start", "end", "operator", "service", "environment", "country", "source"})
			for _, d := range deployments {
				row := []string{strconv.Itoa(d.ID), time.Unix(d.Start/1000, 0).UTC().String(), time.Unix(d.End/1000, 0).UTC().String(), d.Operator, d.Service, d.Environment, d.Country, d.Source}
				_ = csvwriter.Write(row)
			}
			csvwriter.Flush()
			csvfile.Close()
		}
	},
}
