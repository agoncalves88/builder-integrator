package main

import (
	"builder-integrator/configuration"
	"builder-integrator/controllers"
	"builder-integrator/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tkanos/gonfig"
)

var config = configuration.Config{}

func init() {
	gonfig.GetConf("config.yaml", &config)
}

func main() {

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("get-integrator/group/:group/:param", controllers.BuilderController{Config: config}.GetBuilderIntegration)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5555")
}
