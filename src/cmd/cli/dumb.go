package cli

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	"github.com/dannrocha/czen/src/gitscm"
	"github.com/dannrocha/czen/src/semver"
	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Bump(c *cli.Context) error {

	script := setup.Script{}
	script.LoadScript()

	for _, auto := range script.Automation {
		if auto.Bind == BUMP && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	git, err := gitscm.New()

	if err != nil {
		fmt.Println(err)
	}

	lastestTag, ok := git.LastestTag()

	newVersion := incrementVersion(lastestTag);

	
	fmt.Println()

	annonation := newVersion.ConvertToSemver().Version
	hash := fmt.Sprintf("%v-%v", annonation, time.Now())

	message := fmt.Sprintf("%x", hash)

	ok, err = gitscm.CreateTag(annonation, message)

	if !ok || err != nil {
		return nil
	}

	return nil
}

func incrementVersion(tag gitscm.GitTag) semver.Version {

	config := setup.Configuration{}

	profile, errProfile := config.FindCurrentProfileEnable()

	if errProfile != nil {
		panic(errProfile)
	}

	commits, errCommit := gitscm.LoadCommitsFrom(tag.Commit.Hash)

	if errCommit != nil {
		panic(errCommit)
	}

	currentSemever := semver.New(tag.Annotation)
	oldVersion, errVersion := currentSemever.FindVersion()

	if errVersion != nil {
		panic(errVersion)
	}

	newVersion := semver.Version{
		Major: oldVersion.Major,
		Minor: oldVersion.Minor,
		Path:  oldVersion.Path,
	}

	for context, pattern := range profile.BumpMap {
		for _, commit := range commits {
			if strings.Contains(commit.Message, context) {
				newVersion.IncrementVersion(pattern)
			}
		}
	}

	return newVersion
}
