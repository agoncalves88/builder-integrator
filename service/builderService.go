package service

import (
	"builder-integrator/configuration"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

type BuilderService struct {
	Config configuration.Config
}

func (service BuilderService) GetDynamicServices() {
	inputToBuild := map[string]interface{}{
		"CORRELATION": "",
	}

	for _, datasource := range service.Config.DATASOURCES {
		var ret *http.Response
		var err error
		if datasource.HAS_PARAM {
			ret, err = http.Get(fmt.Sprintf(datasource.URL, "1"))
			checkApiErr(err)
		} else {
			ret, err = http.Get(datasource.URL)
			checkApiErr(err)
		}

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
func checkApiErr(err error) {
	if err != nil {
		fmt.Println("erro ao chamar api")
	}
}
