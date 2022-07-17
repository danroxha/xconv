package gitscm

import (
	"errors"
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

func New() (Git, error) {
	git := Git{}
	err := git.loadGitTags()

	if err != nil {
		return Git{}, err
	}

	return git, nil
}

func (git *Git) loadGitTags() error {

	tags, ok := findTagList()

	if !ok {
		return errors.New("fail in load tags")
	}

	git.Tags = tags

	for _, annotation := range git.Tags {
		git.GitTags = append(git.GitTags, gitTagDetails(annotation))
	}

	sort.Sort(GitTagGroup(git.GitTags))
	git.GitTags = util.ReverseSlice(git.GitTags)

	return nil
}

func (git *Git) LastestTag() (GitTag, bool) {

	if git.IsTagsEmpty() {
		return GitTag{}, false
	}

	return git.GitTags[0], true

}

func FormatCommit(messages map[string]string) (string, error) {

	rule := setup.NewRule()
	profile, profileErr := rule.FindCurrentProfileEnable()

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

func gitTagDetails(annotation string) GitTag {
	show := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"show",
			"-q",
			annotation,
		},
	}

	stdout, err := show.Execute()

	if err != nil {
		panic(err)
	}

	tagDetails := string(stdout)

	return parseTag(tagDetails)
}

func (git Git) IsTagsEmpty() bool {
	return len(git.Tags) == 0
}

func CreateTag(annonation string, message string) (bool, error) {
	tag := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"tag",
			"-a",
			annonation,
			"-m",
			message,
		},
	}

	_, err := tag.Execute()

	if err != nil {
		return false, err
	}

	return true, nil
}

func CreateTagFrom(annonation, commit, message string) (bool, error) {
	tag := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"tag",
			"-a",
			annonation,
			commit, 
			"-m",
			message,
		},
	}

	_, err := tag.Execute()

	if err != nil {
		return false, err
	}

	return true, nil
}

func findTagList() ([]string, bool) {
	tag := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"tag",
			"-l",
		},
	}

	stdout, err := tag.Execute()

	if err != nil {
		fmt.Println(err.Error())
		return []string{}, false
	}

	return util.RemoveContains(strings.Split(string(stdout), "\n"), ""), true
}

// https://play.golang.com/p/11oot1NWTPd
func parseTag(commit string) GitTag {

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
