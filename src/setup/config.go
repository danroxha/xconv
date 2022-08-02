package setup

import _ "embed"

//go:embed default.yaml
var XCONVFileContent []byte

//go:embed init.yaml
var XCONVInitialtContent []byte

var Filename string = ".xconv.yaml"

func NewConfiguration() Configuration {
	return Configuration{
		Rule:   NewRule(),
		Script: NewScript(),
	}
}
