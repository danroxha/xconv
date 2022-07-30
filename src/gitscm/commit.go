package gitscm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dannrocha/xconv/src/cmd"
	"github.com/dannrocha/xconv/src/util"
)

func LoadCommitFromBegin() ([]GitCommit, error) {
	rev := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"rev-list",
			"--max-parents=0",
			"HEAD",
		},
	}

	output, err := rev.Execute()

	if err != nil {
		return []GitCommit{}, err
	}

	
	initialCommit := strings.TrimSpace(string(output))
	hashList, err := loadCommitsBetween(initialCommit, "HEAD")

	if err != nil {
		return []GitCommit{}, err
	}

	gitCommitGroup := []GitCommit{}

	for _, hash := range hashList {
		commit := GitCommit{
			Message: findShortMessageFromCommit(hash),
			Hash: hash,
			Date: findDateFromCommit(hash),
			Author: findAuthorFromCommit(hash),
		}

		gitCommitGroup = append(gitCommitGroup, commit)
	}
	
	return gitCommitGroup, nil
}

func LoadCommitsFrom(beginFromCommit string) ([]GitCommit, error) {
	
	hashList, err := loadCommitsBetween(beginFromCommit, "HEAD")

	if err != nil {
		return []GitCommit{}, err
	}

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

func loadCommitsBetween(start string, end string) ([]string, error) {
	log := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"log",
			"--pretty=%H",
			"--author-date-order",
			"--reverse",
			fmt.Sprintf(`%v..%v`, start, end),
		},
	}
	output, err := log.Execute()

	if err != nil {
		return []string{}, err
	}

	return util.RemoveContains(strings.Split(string(output), "\n"), ""), nil
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

	regx := regexp.MustCompile(`\r\n|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`) 
	return regx.ReplaceAllString(string(output), "")
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

	regx := regexp.MustCompile(`\r\n|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`) 
	return regx.ReplaceAllString(string(output), "")
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

	regx := regexp.MustCompile(`\r\n|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`) 
	return regx.ReplaceAllString(string(output), "")
}

func findShortMessageFromCommit(hash string) string {

	show := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"show",
			"-q",
			"--pretty=%B",
			"--oneline",
			hash,
		},
	}

	output, err := show.Execute()

	if err != nil {
		panic(err)
	}
	
	regx := regexp.MustCompile(`\r\n|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`) 
	return regx.ReplaceAllString(string(output), "")
}