package setup

func NewConfiguration() Configuration {
	return Configuration{
		Rule: NewRule(),
		Script: NewScript(),
	}
}
