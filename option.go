package denary

type Option struct {
	value  string
	loaded bool
}

func Nothing() Option {
	return Option{}
}

func something(value string) Option {
	return Option{
		loaded: true,
		value:  value,
	}
}
