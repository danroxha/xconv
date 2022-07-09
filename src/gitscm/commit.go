package gitscm

import (
	"strings"

	"github.com/dannrocha/czen/src/cmd"
	"github.com/dannrocha/czen/src/util"
)


func LoadCommitsFrom(beginFromCommit string) ([]GitCommit, error) {
	log := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"log",
			"--online",
			"--pretty=%H",
			beginFromCommit,
			"HEAD",
		},
	}

	output, err := log.Execute()

	if err != nil {
		return []GitCommit{}, err
	}

	hashList := util.RemoveContains(strings.Split(string(output), "\n"), "")
	
	gitCommitGroup := []GitCommit{}

	for _, hash := range hashList {
		 
	
		commit := GitCommit{
			Message: findMessageFromCommit(hash),
			Hash: hash,
			Date: findDateFromCommit(hash),
			Author: findAuthorFromCommit(hash),
		}

		gitCommitGroup = append(gitCommitGroup, commit)
	}
	
	return gitCommitGroup, nil
}

func findDateFromCommit(hash string) string {
	show := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"show",
			"-q",
			"--pretty=%ad",
			hash,
		},
	}

	output, err := show.Execute()

	if err != nil {
		panic(err)
	}

	return string(output)
}

func findAuthorFromCommit(hash string) string {
	show := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"show",
			"-q",
			`--pretty=format:"%an <%ae>"`,
			hash,
		},
	}

	output, err := show.Execute()

	if err != nil {
		panic(err)
	}

	return string(output)
}

func findMessageFromCommit(hash string) string {

	show := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"show",
			"-q",
			"--pretty=%B",
			hash,
		},
	}

	output, err := show.Execute()

	if err != nil {
		panic(err)
	}

	return string(output)
}