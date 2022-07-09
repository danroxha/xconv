package setup

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)


func NewRule() Rule {
	rule := Rule{}
	rule.loadRuleFromFile()
	rule.setDefaultValues()

	return rule
}

func (conf *Rule) loadRuleFromFile() error {

	file := struct {
		Rule Rule `yaml:"rule"`
	}{}

	blob, err := ioutil.ReadFile(".czen.yaml")

	if err != nil {
		panic(err)
	}

	parseError := yaml.Unmarshal(blob, &file)

	if parseError != nil {
		panic(parseError)
	}

	*conf = file.Rule

	if conf.Version == "" {
		exitStd := ExitCodeStardard["NoVersionSpecifiedError"]
		return errors.New(exitStd.Description)
	}

	return nil
}

func (rule *Rule) FindCurrentProfileEnable() (Profile, error) {

	for _, profile := range rule.Profiles {
		if profile.Name == rule.ActiveProfile {
			return profile, nil
		}
	}

	return Profile{}, errors.New("Profile setupnot found or disabled")
}

func (rule *Rule) setDefaultValues() {

}
