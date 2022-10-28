package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	TESTE       string
	DATASOURCES []struct {
		NAME              string
		URL               string
		PROPERTIES_TO_GET []struct {
			VALUE_TO_GET string
			INPUT_NAME   string
		}
	}
}

var config = Configuration{}

func init() {
	gonfig.GetConf("config.yaml", &config)
}

func main() {
	fmt.Println(config)

	inputToBuild := map[string]interface{}{
		"CORRELATION": "",
	}

	for _, datasource := range config.DATASOURCES {

		ret, err := http.Get(fmt.Sprintf(datasource.URL, "1"))
		CheckApiErr(err)
		body, _ := ioutil.ReadAll(ret.Body)
		fmt.Println(string(body))

		json_map := make(map[string]interface{})
		_ = json.Unmarshal(body, &json_map)

		for _, prop := range datasource.PROPERTIES_TO_GET {
			inputToBuild[prop.INPUT_NAME] = gjson.Get(string(body), prop.VALUE_TO_GET).String()
		}

	}

	jsonBuilded, _ := json.Marshal(inputToBuild)
	fmt.Println(string(jsonBuilded))

}

func CheckApiErr(err error) {
	if err != nil {
		fmt.Println("erro ao chamar api")
	}
}
