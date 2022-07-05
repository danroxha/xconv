package gitscm

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dannrocha/czen/src/cmd"
	"github.com/dannrocha/czen/src/semver"
	"github.com/dannrocha/czen/src/setup"
	"github.com/dannrocha/czen/src/util"
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

func (git *Git) LoadGitTags() error {
	if git.isUndefinedCommand() {

		git.Command = make(map[string]cmd.InternalCommand)

		git.Command["tag"] = cmd.InternalCommand{
			Application: "git",
			Args: []string{
				"tag",
				"-l",
			},
		}
	}

	stdout, err := git.Command["tag"].Execute()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	git.Tags = util.RemoveContains(strings.Split(string(stdout), "\n"), "")
	
	var command string

	for _, version := range git.Tags {

		if git.isUndefinedCommand() {
			git.Command["show"] = cmd.InternalCommand{
				Application: "git",
				Args: []string{
					"show",
					"-q",
					version,
				},
			}
		}else {
			// tests
			MOCK_INDEX_ARG := 1
			if command == "" {
				command = git.Command["show"].Args[MOCK_INDEX_ARG]
			}

			git.Command["show"].Args[MOCK_INDEX_ARG] = fmt.Sprintf(command, version)
		}


		stdout, err = git.Command["show"].Execute()

		if err != nil {
			fmt.Println("git.tag.go", err)
		}

		rawCommit := string(stdout)

		git.GitTags = append(git.GitTags, ParseTag(rawCommit))
	}

	sort.Sort(GitTagGroup(git.GitTags))
	
	return nil
}

func FormatCommit(messages map[string]string) (string, error) {

	conf := setup.Configuration{}
	conf.LoadConfigurationFile()

	profile, profileErr := conf.FindCurrentProfileEnable()

	if profileErr != nil {
		// exitStd := setupExitCodeStardard["InvalidConfigurationError"]
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

// https://play.golang.com/p/11oot1NWTPd
func ParseTag(commit string) GitTag {

	tag := make(map[string]string)

	regexGroup := map[string]func(string) string{
		"author": func(content string) string {
			regex := regexp.MustCompile(`Author:\s+(.+)`)
			return strings.TrimSpace(regex.FindStringSubmatch(content)[1])
		},
		"hash": func(content string) string {
			regex := regexp.MustCompile(`commit\s+([a-f0-9]+)`)
			return strings.TrimSpace(regex.FindStringSubmatch(content)[1])
		},
		"tagger": func(content string) string {
			regex := regexp.MustCompile(`Tagger:\s+(.+)`)
			return strings.TrimSpace(regex.FindStringSubmatch(content)[1])
		},

		"datetag": func(content string) string {
			regex := regexp.MustCompile(`(\w+:\s+[[:word:]]+\s+[[:word:]]+\s+[1-9]{1,2}\s+[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}\s+[0-9]{1,4}\s+.\d+)`)
			match := regex.FindAllStringSubmatch(content, -1)
			firstDate := match[0][1]
			return strings.TrimSpace(strings.Split(firstDate, ": ")[1])
		},
		"datecommit": func(content string) string {
			regex := regexp.MustCompile(`(\w+:\s+[[:word:]]+\s+[[:word:]]+\s+[1-9]{1,2}\s+[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}\s+[0-9]{1,4}\s+.\d+)`)
			match := regex.FindAllStringSubmatch(content, -1)
			secondDate := match[1][1]
			return strings.TrimSpace(strings.Split(secondDate, ": ")[1])
		},
		"annonation": func(content string) string {
			regex := regexp.MustCompile(`tag\s+(.+)`)
			return strings.TrimSpace(regex.FindStringSubmatch(content)[1])
		},
	}

	for key, regex := range regexGroup {
		tag[key] = regex(commit)
	}

	return GitTag{
		Author:     tag["tagger"],
		Annotation: tag["annonation"],
		Date:       tag["datetag"],
		Commit: GitCommit{
			Author: tag["author"],
			Hash:   tag["hash"],
			Date:   tag["datecommit"],
		},
	}
}


func (git Git) isUndefinedCommand() bool {
	_, undefinedShow := git.Command["show"]
	_, undefinedTag := git.Command["tag"]

	return !undefinedShow || !undefinedTag
}

func (git Git) IsTagsEmpty() bool {
	return len(git.Tags) == 0
}
