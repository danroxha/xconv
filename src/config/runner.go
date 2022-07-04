package config

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func (conf *Configuration) LoadConfigurationFile() error {

	file := struct {
		Configuration Configuration `yaml:"configuration"`
	}{}

	blob, err := ioutil.ReadFile(".czen.yaml")

	if err != nil {
		panic(err)
	}

	parseError := yaml.Unmarshal(blob, &file)

	if parseError != nil {
		panic(parseError)
	}

	*conf = file.Configuration

	if conf.Version == "" {
		exitStd := ExitCodeStardard["NoVersionSpecifiedError"]
		return errors.New(exitStd.Description)
	}

	return nil
}

func (script *Script) LoadScript() error {
	file := struct {
		Script Script `yaml:"script"`
	}{}

	blob, err := ioutil.ReadFile(".czen.yaml")

	if err != nil {
		panic(err)
	}

	parseError := yaml.Unmarshal(blob, &file)

	if parseError != nil {
		panic(parseError)
	}

	*script = file.Script
	
	return nil
}

func (sc *Script) FindAllFilters() map[string]Filter {
	filterGroup := make(map[string]Filter)

	for _, filter := range sc.Filter {
		filterGroup[filter.Name] = filter
	}

	return filterGroup
}

func (sc *Script) FindAllMiddlewares() map[string]Middleware {
	middlewareGroup := make(map[string]Middleware)

	for _, middleware := range sc.Middleware {
		middlewareGroup[middleware.Name] = middleware
	}

	return middlewareGroup
}

func (conf *Configuration) FindCurrentProfileEnable() (Profile, error) {

	for _, profile := range conf.Profiles {
		if profile.Name == conf.ActiveProfile {
			return profile, nil
		}
	}

	return Profile{}, errors.New("Profile config not found or disabled")
}
