package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/cmd"
	"github.com/dannrocha/czen/src/git"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func Rollback(c *cli.Context) error {

	tags, err := git.LoadGitTags()

	var items []string = []string{}

	if err != nil {
		panic(err)
	}

	for _, tag := range tags {
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

	toTag := tags[index]

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
			"reset",
			"--hard",
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
