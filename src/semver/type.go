package semver

var MAJOR, MINOR, PATCH int = 0, 1, 2

var SEMVER_REGEX string = `^(\w+)?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

type Version struct {
	Major int
	Minor int
	Path  int
}

type SemVer struct {
	Version string
}
