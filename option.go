package denary

type Option struct {
	value  string
	loaded bool
}

func Nothing() Option {
	return Option{}
}
