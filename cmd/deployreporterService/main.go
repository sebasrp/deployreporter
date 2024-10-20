package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sebasrp/deployreporter"
	"gopkg.in/yaml.v2"
)

type Config struct {
	GrafanaKey string `yaml:"grafanaKey"`
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	home, _ := os.UserHomeDir()
	file, err := os.Open(home + "/.deployreporter.yaml") //TODO: we should really make this configurable
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func setupRouter() *gin.Engine {
	config, err := ReadConfig()
	if err != nil {
		fmt.Printf("error reading config file. %v", err)
	}
	fmt.Printf("config: %v", config)

	r := gin.Default()
	r.Use(cors.Default())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Ping test
	r.GET("/deployments", func(c *gin.Context) {
		limitDefault := 400
		limitString := c.DefaultQuery("limit", strconv.Itoa(limitDefault))
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			fmt.Printf("error parsing query string limit: %v.\n Error: %v", limitString, err)
			limit = limitDefault
		}
		c.AsciiJSON(http.StatusOK, deployreporter.GetDeployments("", "", limit, config.GrafanaKey))
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
