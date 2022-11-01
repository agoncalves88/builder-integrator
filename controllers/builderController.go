package controllers

import (
	"builder-integrator/configuration"
	"builder-integrator/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BuilderController struct {
	Config configuration.Config
}

// @BasePath /api/v1

// PingExample godoc
// @Summary builder
// @Schemes
// @Description dynamic builder integration
// @Tags builder
// @Accept json
// @Produce json
// @Param param  path string true "{Param}"
// @Success 200 {string} Helloworld
// @Router /api/v1/get-integrator/{param} [get]
func (config BuilderController) GetBuilderIntegration(c *gin.Context) {

	service := service.BuilderService{Config: config.Config}
	param := c.Param("param")
	builder := service.GetDynamicServices(param)
	c.JSON(http.StatusOK, builder)
}
