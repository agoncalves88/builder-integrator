package main

import (
	"builder-integrator/configuration"
	"builder-integrator/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tkanos/gonfig"
)

var config = configuration.Config{}

func init() {
	gonfig.GetConf("config.yaml", &config)
}

func main() {

	configuration.SetSwaggerInfo()
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/get-integrator/:param", func(ctx *gin.Context) {
			param := ctx.Param("param")
			builder := service.BuilderService{Config: config}
			jsonBuilded := builder.GetDynamicServices(param)

			ctx.JSON(http.StatusOK, jsonBuilded)
		})
	}

	// url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition

	url := ginSwagger.URL("http://localhost:5555/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":5555")
}
