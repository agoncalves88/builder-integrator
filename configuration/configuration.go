package configuration

import "builder-integrator/docs"

type Config struct {
	TESTE       string
	DATASOURCES []struct {
		NAME              string
		URL               string
		HAS_PARAM         bool
		PROPERTIES_TO_GET []struct {
			VALUE_TO_GET string
			INPUT_NAME   string
		}
	}
}

func SetSwaggerInfo() {
	docs.SwaggerInfo.Title = "Builder Integration API"
	docs.SwaggerInfo.Description = "This is a concept api to builder dynamic response"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
