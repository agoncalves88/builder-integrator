package service

import (
	"builder-integrator/configuration"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/tidwall/gjson"
)

type BuilderService struct {
	Config configuration.Config
}

func (service BuilderService) GetDynamicServices(param string) map[string]interface{} {
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(service.Config.DATASOURCES))
	inputToBuild := map[string]interface{}{
		"CORRELATION": "",
	}

	for _, datasource := range service.Config.DATASOURCES {
		go getService(datasource, param, &waitGroup, &inputToBuild)
	}

	waitGroup.Wait()
	return inputToBuild
}

func getService(datasource configuration.DataSource, param string, group *sync.WaitGroup, inputBuilded *map[string]interface{}) {
	var ret *http.Response
	var err error
	if datasource.HAS_PARAM {
		ret, err = http.Get(fmt.Sprintf(datasource.URL, param))
		checkApiErr(err)
	} else {
		ret, err = http.Get(datasource.URL)
		checkApiErr(err)
	}

	body, _ := ioutil.ReadAll(ret.Body)
	toBuilded := *inputBuilded
	for _, prop := range datasource.PROPERTIES_TO_GET {
		toBuilded[prop.INPUT_NAME] = gjson.Get(string(body), prop.VALUE_TO_GET).String()
	}
	*inputBuilded = toBuilded

	group.Done()
}
func checkApiErr(err error) {
	if err != nil {
		fmt.Println("erro ao chamar api")
	}
}
