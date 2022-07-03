package git

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
}
