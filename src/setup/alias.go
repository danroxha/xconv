package setup

func (alias *Alias) LoadAlias() error {

	_ = struct {
		Alias Alias `yaml:"alias"`
	}{}

	return nil
}