package main

import (
	"builder-integrator/configuration"
	"builder-integrator/service"
	"fmt"

	"github.com/tkanos/gonfig"
)

var config = configuration.Config{}

func init() {
	gonfig.GetConf("config.yaml", &config)
}

func main() {
	fmt.Println(config)

	builder := service.BuilderService{Config: config}
	builder.GetDynamicServices()

}
