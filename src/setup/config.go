package setup

import (
	"errors"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v3"
)


func (conf *Configuration) LoadConfiguration() error {
	
	blob, err := ioutil.ReadFile(".czen.yaml")

	if err != nil {
		panic(err)
	}

	parseError := yaml.Unmarshal(blob, &conf)

	if parseError != nil {
		panic(parseError)
	}

	if conf.Role.Version == "" {
		exitStd := ExitCodeStardard["NoVersionSpecifiedError"]
		return errors.New(exitStd.Description)
	}

	return nil
}

func (conf Configuration) IsEmpty() bool {
	return reflect.DeepEqual((Configuration{}), conf)
}