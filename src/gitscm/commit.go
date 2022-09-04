package gitscm

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dannrocha/xconv/src/cmd"
	"github.com/dannrocha/xconv/src/util"
)

var wg sync.WaitGroup
var load sync.WaitGroup
var amount int64

func LoadCommitFromBegin() ([]GitCommit, error) {

	amount = 0

	rev := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"rev-list",
			"--max-parents=0",
			"HEAD",
			"-1",
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
	shared := make(chan string)
	parse := make(chan GitCommit)
	MAX_GO_ROUTINE := 320
	wg.Add(1)
	load.Add(MAX_GO_ROUTINE)
	
	for i := 0; i < MAX_GO_ROUTINE; i++ {
		go loadCommit(shared, parse)
	}

	go func(parse chan GitCommit) {
		defer wg.Done()
		for  {
			commit, ok := <- parse
			atomic.AddInt64(&amount, 1)
			if !ok || atomic.LoadInt64(&amount) == int64(len(hashList)) {	
				break
			}
			gitCommitGroup = append(gitCommitGroup, commit)
		}
	}(parse)

	for _, hash := range hashList {
		shared <- hash
	}

	close(shared)
	
	wg.Wait()

	return gitCommitGroup, nil
}

func loadCommit(channel chan string, parse chan<- GitCommit) {

	for {
		hash, ok := <-channel
		if !ok {
			break
		}

		commit := GitCommit{
			Message: findShortMessageFromCommit(hash),
			Hash:    hash,
			Date:    findDateFromCommit(hash),
			Author:  findAuthorFromCommit(hash),
		}
		parse <- commit
	}
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
			Hash:    hash,
			Date:    findDateFromCommit(hash),
			Author:  findAuthorFromCommit(hash),
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

	commits := util.RemoveContains(strings.Split(string(output), "\n"), "")
	commits = append([]string{start}, commits...)

	return commits, nil
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
