package gitscm

type Git struct {
	Tags       []string
	GitTags    []GitTag
	GitCommits []GitCommit
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
