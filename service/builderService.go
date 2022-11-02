package service

import (
	"builder-integrator/configuration"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/tidwall/gjson"
)

type BuilderService struct {
	Config configuration.Config
}

func (service BuilderService) GetDynamicServices(param string, group string) map[string]interface{} {
	var waitGroup sync.WaitGroup

	inputToBuild := map[string]interface{}{
		"CORRELATION": "",
	}
	totalDataForGroups := 0
	for _, datasource := range service.Config.DATASOURCES {
		if strings.EqualFold(datasource.GROUP, group) {
			totalDataForGroups++
		}
	}
	waitGroup.Add(totalDataForGroups)

	for _, datasource := range service.Config.DATASOURCES {
		if strings.EqualFold(datasource.GROUP, group) {

			go getService(datasource, param, &waitGroup, &inputToBuild)
		}

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
		value := gjson.Get(string(body), prop.VALUE_TO_GET).String()
		if value == "" {
			if toBuilded[prop.INPUT_NAME] == nil {
				toBuilded[prop.INPUT_NAME] = ""
			}
		} else {
			toBuilded[prop.INPUT_NAME] = value
		}

	}
	*inputBuilded = toBuilded

	group.Done()
}
func checkApiErr(err error) {
	if err != nil {
		fmt.Println("erro ao chamar api")
	}
}
