package cli

import (
	"crypto/md5"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/dannrocha/xconv/src/gitscm"
	"github.com/dannrocha/xconv/src/semver"
	"github.com/dannrocha/xconv/src/setup"
	"github.com/urfave/cli/v2"
)

func Bump(c *cli.Context) error {
	script := setup.NewScript()

	for _, task := range script.Task {
		if task.Bind == BUMP && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
			}
		}
	}

	git, err := gitscm.New()

	if err != nil {
		fmt.Println(err)
	}

	lastestTag, ok := git.LastestTag()

	if !ok {
		return nil
	}

	newVersion := incrementVersion(lastestTag)

	annonation := newVersion.ConvertToSemver().Version

	if annonation == lastestTag.Annotation {
		fmt.Println("there are no relevant changes for a new version")
		return nil
	}

	hash := md5.Sum([]byte(fmt.Sprintf("%v-%v", annonation, time.Now())))

	ok, err = gitscm.CreateTag(annonation, fmt.Sprintf("%x", hash))
	if !ok || err != nil {
		panic(err)
	}

	fmt.Printf("bump: version %s \u2192 %s\n", lastestTag.Annotation, annonation)

	return nil
}

func incrementVersion(tag gitscm.GitTag) semver.Version {
	rule := setup.NewRule()
	profile, errProfile := rule.FindCurrentProfileEnable()

	if errProfile != nil {
		panic(errProfile)
	}

	commits, errCommit := gitscm.LoadCommitsFrom(tag.Commit.Hash)
	commits = commits[1:]

	if errCommit != nil {
		panic(errCommit.Error())
	}

	currentSemever := semver.New(tag.Annotation)
	oldVersion, errVersion := currentSemever.FindVersion()

	if errVersion != nil {
		panic(errVersion.Error())
	}

	newVersion := semver.Version{
		Major: oldVersion.Major,
		Minor: oldVersion.Minor,
		Path:  oldVersion.Path,
	}

	for _, commit := range commits {
		for context, pattern := range profile.Bump.Map {
			if profile.Bump.Pattern != "" {
				regex, err := regexp.Compile(profile.Bump.Pattern)

				if err != nil {
					exception := setup.ExitCodeStardard["BumpRegexInvalid"]
					fmt.Println(exception.Description)
					os.Exit(exception.ExitCode)
				}

				if regex.Match([]byte(commit.Message)) {
					newVersion.IncrementVersion(pattern, profile.Tag.Mode)
				}

				break
			} else if strings.Contains(commit.Message, context) {
				newVersion.IncrementVersion(pattern, profile.Tag.Mode)
				break
			}
		}
	}

	return newVersion
}
