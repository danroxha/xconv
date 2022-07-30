package setup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dannrocha/xconv/src/util"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
)


func NewRule() Rule {
	rule := Rule{}
	rule.loadRuleFromFile()
	rule.setDefaultValues()
	rule.mergeExtendsProfiles()

	return rule
}

func (conf *Rule) loadRuleFromFile() error {

	file := struct {
		Rule Rule `yaml:"rule"`
	}{}

	blob, err := ioutil.ReadFile(Filename)

	if err != nil {
		blob = XCONVFileContent
	}

	parseError := yaml.Unmarshal(blob, &file)

	if parseError != nil {
		fmt.Println(parseError.Error())
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

	return Profile{}, errors.New("Profile setup not found or disabled")
}

func (rule *Rule) FindProfileByName(profileName string) (Profile, int, error) {
	for index, profile := range rule.Profiles {
		if profile.Name == profileName {
			return profile, index, nil
		}
	}

	return Profile{}, -1, errors.New("profile setup not found or disabled")
}

func (rule *Rule) ReplaceProfile(p Profile) error {
	for index, profile := range rule.Profiles {
		if profile.Name == p.Name {
			rule.Profiles[index] = p
			return nil
		}
	}

	return errors.New("profile not found")
}

func (rule *Rule) mergeExtendsProfiles() {
	for indexA, profileA := range rule.Profiles {
		for indexB, profileB := range rule.Profiles {
			if profileA.Name == profileB.Name {
				continue
			}

			rule.mergeExtendsRecursive(&profileA, indexA, []string{})

			if profileA.Name == profileB.Extends {
				mergo.Merge(&profileB, profileA)
				rule.Profiles[indexB] = profileB
			}
		}	
	}
}


func (rule *Rule) mergeExtendsRecursive(profile *Profile, index int, stack []string)  {
	super, indexOf, err := rule.FindProfileByName(profile.Extends)

	if err != nil {
		return
	}

	if util.ContainsSlice(stack, super.Name) {
		stack = append(stack, super.Name)
		exception := ExitCodeStardard["InvalidProfile"]
		
		var recursive string = ""

		for index, name := range stack {
			if index < len(stack) - 1 {
				recursive += name + " \u2192 "
			}else {
				recursive += name
			}
		}


		fmt.Printf("%s - %s", exception.Description, recursive)
		os.Exit(exception.ExitCode)
	}

	stack = append(stack, super.Name)

	if super.Extends != "" {
		rule.mergeExtendsRecursive(&super, indexOf, stack)
	}

	mergo.Merge(profile, super)
	rule.Profiles[index] = *profile
}

func (rule *Rule) setDefaultValues() {
	
	defaultConfig := struct {
		Rule Rule `yaml:"rule"`
	}{}

	err := yaml.Unmarshal(XCONVFileContent, &defaultConfig)

	if err != nil {
		panic(err)
	}

	defaultProfile, _,  errDefaultProfile := defaultConfig.Rule.FindProfileByName("xconv_default")
	extendsDefaultProfile, _, errExtendsDefaultProfile := rule.FindProfileByName("xconv_default")

	if errDefaultProfile != nil {
		panic(err)
	}

	if errExtendsDefaultProfile != nil {
		rule.Profiles = append(rule.Profiles, defaultProfile)
		return
	}
	
	mergo.Merge(&extendsDefaultProfile, defaultProfile)

	rule.ReplaceProfile(extendsDefaultProfile)

}
