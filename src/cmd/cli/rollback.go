package cli

import (
	"fmt"

	"github.com/dannrocha/xconv/src/cmd"
	"github.com/dannrocha/xconv/src/gitscm"
	"github.com/dannrocha/xconv/src/setup"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func Rollback(c *cli.Context) error {
	script := setup.NewScript()

	for _, task := range script.Task {

		if task.Bind == ROLLBACK && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
			}
		}
	}

	git, err := gitscm.New()

	var items []string = []string{}

	if err != nil {
		panic(err)
	}

	for _, tag := range git.GitTags {
		items = append(items, tag.Annotation)
	}

	prompt := promptui.Select{
		Label: "Select a tag version to rollback",
		Items: items,
	}

	index, _, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	toTag := git.GitTags[index]

	fmt.Printf(`Selected:
    Version: %v
    Author: %v
    Reference commit: %v\n `, toTag.Annotation, toTag.Author, toTag.Commit.Hash)

	confirm := promptui.Prompt{
		Label:     "Do you really want to rollback to version above. Current branch will be changed",
		Default:   "N",
		IsConfirm: true,
	}

	_, err = confirm.Run()

	confirmation := false

	if err != nil {
		confirmation = false
	} else {
		confirmation = true
	}

	if !confirmation {
		return nil
	}

	command := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"revert",
			"--no-edit",
			"--no-commit",
			toTag.Commit.Hash,
		},
	}

	_, err = command.Execute()

	if err != nil {
		fmt.Printf("Can't rollback to %v with hash [ %v ]\n", toTag.Annotation, toTag.Commit.Hash)
		return err
	}

	fmt.Printf("Current version %v", toTag.Annotation)

	return nil
}
