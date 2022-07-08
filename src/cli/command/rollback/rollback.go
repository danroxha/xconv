package rollback

import (
	"fmt"

	"github.com/dannrocha/czen/src/cli"
	"github.com/dannrocha/czen/src/gitscm"
	"github.com/dannrocha/czen/src/setup"
	"github.com/manifoldco/promptui"
)

func Execute(args ...string) error {

	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {

		if auto.Bind == cmd.ROLLBACK && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
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

	err = gitscm.BackToCommit(toTag.Commit.Hash)

	if err != nil {
		fmt.Printf("can't rollback to %v with hash [ %v ]\n", toTag.Annotation, toTag.Commit.Hash)
		return err
	}

	fmt.Printf("current version %v", toTag.Annotation)

	return nil
}
