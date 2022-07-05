package setup

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

func (conf *Configuration) FindCurrentProfileEnable() (Profile, error) {

	for _, profile := range conf.Profiles {
		if profile.Name == conf.ActiveProfile {
			return profile, nil
		}
	}

	return Profile{}, errors.New("Profile setupnot found or disabled")
}
