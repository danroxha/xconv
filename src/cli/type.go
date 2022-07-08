package cmd

const (
	COMMIT    = "commit"
	BUMP      = "bump"
	ROLLBACK  = "rollback"
	EXAMPLE   = "example"
	INIT      = "init"
	SCHEMA    = "schema"
	TAG       = "tag"
	CHANGELOG = "changelog"
	VERSION   = "version"
)

type InternalCommand struct {
	Application string
	Args        []string
}
