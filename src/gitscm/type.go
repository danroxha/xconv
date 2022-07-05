package gitscm

import "github.com/dannrocha/czen/src/cmd"


type Git struct {
	Tags []string
	GitTags []GitTag
	GitCommits []GitCommit
	Command  map[string]cmd.InternalCommand
}

type GitCommit struct {
	Author  string
	Date    string
	Hash    string
	Message string
}

type GitTag struct {
	Annotation string
	Author     string
	Commit     GitCommit
	Date       string
}
