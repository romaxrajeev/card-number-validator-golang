package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gitlab.eb-tools.com/engineering/cem-public-safety/public-warning/qa/eb-gitlab-integration/branches"
	"gitlab.eb-tools.com/engineering/cem-public-safety/public-warning/qa/eb-gitlab-integration/pipeline"
)

func main() {
	// Loading env variables
	error := godotenv.Load()
	if error != nil {
		panic(error)
	}
	// Setup Viper
	arg := "." //os.Args[1]
	fmt.Println("Command line config file: " + arg)
	if arg != "" {
		viper.AddConfigPath(arg)
	} else {
		fmt.Println("Getting config file from . location")
		viper.AddConfigPath(".")
	}
	viper.SetConfigName("app.config.json")
	viper.SetConfigType("json")
	viper.WatchConfig()
	errConf := viper.ReadInConfig()
	if errConf != nil { // Handle errors reading the config file
		panic(errConf.Error())
	}

	// Router setup
	router := gin.Default()

	// Setting up CORS
	config := cors.DefaultConfig()
	if viper.GetString("env") == "development" {
		config.AllowOrigins = []string{"*"}
	} else {
		config.AllowOrigins = []string{"http://localhost:3000", "http://3.77.243.153:3000"}
	}
	config.AllowCredentials = true
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	router.Use(cors.New(config))

	// Registering routes
	router.POST("/run-pipeline", pipeline.RunPipeline)
	router.GET("/pipeline-runs", pipeline.GetPipelineRuns)
	router.GET("/list-branches", branches.ListBranches)

	// Run the server
	router.Run(viper.GetString("ip") + ":" + viper.GetString("port"))
}
