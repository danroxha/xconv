package cli

import (
	"crypto/md5"
	"fmt"
	"os"
	"time"

	"github.com/dannrocha/xconv/src/gitscm"
	"github.com/dannrocha/xconv/src/semver"
	"github.com/dannrocha/xconv/src/setup"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func Init(c *cli.Context) error {
	script := setup.NewScript()

	for _, task := range script.Task {
		if task.Bind == INIT && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
			}
		}
	}

	checkSetupExist()
	createInitialVersion()
	createSetupFile()

	return nil
}

func formatAnnotation(annotation string) string {
	newSemVer := semver.SemVer{
		Version: annotation,
	}

	newVersion, err := newSemVer.FindVersion()

	if err != nil {
		panic(err)
	}

	return newVersion.ConvertToSemver().Version
}

func checkSetupExist() {
	file, err := os.Open(setup.Filename)
	
	if err == nil {
		throw := setup.ExitCodeStardard["FileSetupExist"]
		fmt.Println(throw.Description)
		os.Exit(throw.ExitCode)
	}

	defer file.Close()
}

func createSetupFile() {
	file, err := os.Create(setup.Filename)

	if err != nil {
		throw := setup.ExitCodeStardard["NoPermissionOnDir"]
		fmt.Println(throw.Description)
		os.Exit(throw.ExitCode)
	}

	defer file.Close()

	file.Write([]byte(setup.XCONVInitialtContent))
}

func createInitialVersion() {
	var items []string = []string{}
	commits, err := gitscm.LoadCommitFromBegin()

	if err != nil {
		exception := setup.ExitCodeStardard["NoCommitsFoundError"]
		fmt.Println("there are no commits for versioning")
		os.Exit(exception.ExitCode)
	}

	for _, commit := range commits {
		items = append(items, fmt.Sprintf("%v - %v", commit.Date, commit.Message))
	}

	selectCommitPrompt := promptui.Select{
		Label: "Select a commit for the initial version",
		Items: items,
	}

	index, _, err := selectCommitPrompt.Run()

	if err != nil {
		return
	}

	selected := commits[index]

	insertInicalVersionPrompt := promptui.Prompt{
		Label:    "Initial version (1.0.0): ",
		Validate: semver.IsVersionValid,
		Default: "1.0.0",
	}

	annotation, err := insertInicalVersionPrompt.Run()

	if err != nil {
		return
	}
	
	annotation = formatAnnotation(annotation)

	hash := md5.Sum([]byte(fmt.Sprintf("%v-%v", annotation, time.Now())))
	gitscm.CreateTagFrom(annotation, selected.Hash, fmt.Sprintf("%x", hash))
}