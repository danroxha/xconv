package setup

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func (role *Role) LoadRole() error {

	file := struct {
		Role Role `yaml:"configuration"`
	}{}

	blob, err := ioutil.ReadFile(".czen.yaml")

	if err != nil {
		panic(err)
	}

	parseError := yaml.Unmarshal(blob, &file)

	if parseError != nil {
		panic(parseError)
	}

	*role = file.Role

	if role.Version == "" {
		exitStd := ExitCodeStardard["NoVersionSpecifiedError"]
		return errors.New(exitStd.Description)
	}

	return nil
}

func (role *Role) FindCurrentProfileEnable() (Profile, error) {

	for _, profile := range role.Profiles {
		if profile.Name == role.ActiveProfile {
			return profile, nil
		}
	}

	return Profile{}, errors.New("Profile setupnot found or disabled")
}
