package main

import (
	"builder-integrator/configuration"
	"builder-integrator/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
)

var config = configuration.Config{}

func init() {
	gonfig.GetConf("config.yaml", &config)
}

func main() {
	r := gin.Default()
	r.GET("/get-integrator/:param", func(ctx *gin.Context) {
		param := ctx.Param("param")
		builder := service.BuilderService{Config: config}
		jsonBuilded := builder.GetDynamicServices(param)

		ctx.JSON(http.StatusOK, jsonBuilded)
	})

	r.Run(":5555")
}
