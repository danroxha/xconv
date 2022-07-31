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
	rule.handleProfilePropertyInheritance()

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

func (rule *Rule) FindProfileByName(profileName string) (Profile, error) {
	for _, profile := range rule.Profiles {
		if profile.Name == profileName {
			return profile, nil
		}
	}

	return Profile{}, errors.New("profile setup not found or disabled")
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

func (rule *Rule) handleProfilePropertyInheritance() {
	for _, profile := range rule.Profiles {
		rule.resolveInheritedProperties(profile, []string{})
	}
}

func (rule *Rule) resolveInheritedProperties(profile Profile, stackTrace []string)  {
	parent, err := rule.FindProfileByName(profile.Extends)

	if err != nil {
		profile.processed = true
		rule.ReplaceProfile(profile)
		return
	}

	if util.ContainsSlice(stackTrace, parent.Name) {
		stackTrace = append(stackTrace, parent.Name)
		exception := ExitCodeStardard["InvalidProfile"]
		
		dependencyString := buildDependencyStringArrow(stackTrace)

		fmt.Printf("%s - %s", exception.Description, dependencyString)
		os.Exit(exception.ExitCode)
	}

	stackTrace = append(stackTrace, parent.Name)

	if parent.Extends != "" || !parent.processed {
		rule.resolveInheritedProperties(parent, stackTrace)
	}
	
	profile.processed = true

	mergo.Merge(&profile, parent)
	rule.ReplaceProfile(profile)
}

func buildDependencyStringArrow(stack []string) string {
	s := ""
	for index, name := range stack {
		if index < len(stack)-1 {
			s += name + " \u2192 "
		} else {
			s += name
		}
	}
	return s
}

func (rule *Rule) setDefaultValues() {
	defaultConfig := struct {
		Rule Rule `yaml:"rule"`
	}{}

	err := yaml.Unmarshal(XCONVFileContent, &defaultConfig)

	if err != nil {
		panic(err)
	}

	defaultProfile, errDefaultProfile := defaultConfig.Rule.FindProfileByName("xconv_default")
	extendsDefaultProfile, errExtendsDefaultProfile := rule.FindProfileByName("xconv_default")

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
