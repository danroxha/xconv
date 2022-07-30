package setup

import _ "embed"

//go:embed default.yaml
var XCONVFileContent []byte

var Filename string = ".xconv.yaml"

func NewConfiguration() Configuration {
	return Configuration{
		Rule:   NewRule(),
		Script: NewScript(),
	}
}
