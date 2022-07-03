package git

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dannrocha/czen/src/cmd"
	"github.com/dannrocha/czen/src/config"
	"github.com/dannrocha/czen/src/semver"
	"github.com/flosch/pongo2"
)

type GitTagGroup []GitTag

var _ interface{}

func (version GitTagGroup) Len() int {
	return len(version)
}

func (version GitTagGroup) Swap(i, j int) {
	version[i], version[j] = version[j], version[i]
}

func (version GitTagGroup) Less(i, j int) bool {
	semVersionA := semver.New(version[i].Annotation)
	semVersionB := semver.New(version[j].Annotation)

	versionA, _ := semVersionA.FindVersion()
	versionB, _ := semVersionB.FindVersion()	

	if versionA.Major < versionB.Major {
		return true
	}

	if versionA.Major <= versionB.Major && versionA.Minor < versionB.Minor {
		return true
	}

	if versionA.Major <= versionB.Major && versionA.Minor <= versionB.Minor && versionA.Path < versionB.Path {
		return true
	}

	return false
}

func LoadGitTags() ([]GitTag, error) {
	tags := []GitTag{}

	command := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"tag",
			"-l",
		},
	}

	stdout, err := command.Execute()

	if err != nil {
		fmt.Println(err.Error())
		return tags, err
	}

	tagList := strings.Split(string(stdout), "\n")

	regexGroup := map[string]func(string) string{
		"author": func(content string) string {
			regex := regexp.MustCompile("Author.\\s+(\\w+\\s+<.+>)")
			return strings.TrimSpace(strings.Split(regex.FindString(content), " ")[1])
		},
		"hash": func(content string) string {
			regex := regexp.MustCompile("commit\\s+([a-f0-9]+)")
			return strings.TrimSpace(strings.Split(regex.FindString(content), " ")[1])
		},
		"tagger": func(content string) string {
			regex := regexp.MustCompile("Tagger.\\s+(\\w+\\s+<.+>)")
			return strings.TrimSpace(strings.Split(regex.FindString(content), ":")[1])
		},
	}

	for _, version := range tagList {
		command = cmd.InternalCommand{
			Application: "git",
			Args: []string{
				"show",
				"-q",
				version,
			},
		}

		stdout, err = command.Execute()

		rawCommit := string(stdout)

		if len(rawCommit) == 0 {
			continue
		}

		tag := make(map[string]string)

		for key, regex := range regexGroup {
			tag[key] = regex(rawCommit)
		}

		tags = append(tags, GitTag{
			Author:     tag["tagger"],
			Annotation: version,
			Commit: GitCommit{
				Author: tag["author"],
				Hash:   tag["hash"],
			},
		})
	}

	sort.Sort(GitTagGroup(tags))

	return tags, nil
}

func FormatCommit(messages map[string]string) (string, error) {

	conf := config.Configuration{}
	conf.LoadConfigurationFile()

	profile, profileErr := conf.FindCurrentProfileEnable()

	if profileErr != nil {
		// exitStd := config.ExitCodeStardard["InvalidConfigurationError"]
		// cli.Exit(exitStd.Description, exitStd.ExitCode)
		os.Exit(1)
	}

	parse, err := pongo2.FromString(profile.MessageTemplate)

	if err != nil {
		panic(err)
	}

	var arguments pongo2.Context = make(pongo2.Context)

	for key, value := range messages {
		boolean, err := strconv.ParseBool(value)

		if err == nil {
			arguments[key] = boolean
			continue
		}

		arguments[key] = value
	}

	commitFormated, parseError := parse.Execute(arguments)

	if parseError != nil {
		panic(parseError)
	}

	return commitFormated, nil
}
