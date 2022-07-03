package cli

import (
	"fmt"
	"strings"

	"github.com/dannrocha/czen/src/cmd"
	"github.com/dannrocha/czen/src/config"
	"github.com/dannrocha/czen/src/git"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func Commit(c *cli.Context) error {

	messages := make(map[string]string)
	conf := config.Configuration{}
	conf.LoadConfigurationFile()

	profile, profileErr := conf.FindCurrentProfileEnable()

	if profileErr != nil {
		return nil
	}

	var questionGroup config.Question = config.Question{}

	for _, question := range profile.Questions {
		if question.Type == "list" {
			questionGroup = question
			break
		}
	}

	options := questionGroup.Choices

	prompt := promptui.Select{
		Label: "Select the type of change you are committing",
		Items: optionDescription(options),
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	messages[questionGroup.Name] = options[index].Value

	sampleInputGroup := []config.Question{}

	for _, question := range profile.Questions {
		if question.Type == "list" {
			continue
		}

		sampleInputGroup = append(sampleInputGroup, question)
	}

	parse := map[string]func(string) string{

		"input": func(content string) string {
			return strings.TrimSpace(content)
		},

		"confirm": func(content string) string {
			trueValues := []string{
				"y",
				"yes",
			}

			var value bool = false
			content = strings.ToLower(strings.TrimSpace(content))

			for _, confirm := range trueValues {
				if confirm == content {
					value = true
					break
				}
			}

			return fmt.Sprintf("%t", value)
		},
	}

	for _, input := range sampleInputGroup {
		clientInput := ""

		if input.Type == "confirm" {
			confirm := promptui.Prompt{
				Label:     input.Message,
				Default:   input.Default,
				IsConfirm: true,
			}

			action, err := confirm.Run()

			if err != nil {
				action = fmt.Sprint(false)
				continue
			}

			clientInput = action

		} else {
			clientInput = parse[input.Type](cmd.ReadInput(input.Message))
		}

		if input.Default != "" && strings.TrimSpace(clientInput) == "" {
			messages[input.Name] = parse[input.Type](input.Default)
			continue
		}

		messages[input.Name] = clientInput
	}

	commitMessage, _ := git.FormatCommit(messages)

	command := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"commit",
			"-m",
			commitMessage,
		},
	}

	stdout, err := command.Execute()

	if err != nil {
		fmt.Println(string(stdout))
	}

	return nil
}

func optionDescription(optionGroup []config.Option) []string {

	names := []string{}

	for _, option := range optionGroup {
		names = append(names, option.Name)
	}

	return names
}