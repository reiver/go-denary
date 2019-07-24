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

func (receiver Option) Result() Result {
	if Nothing() == receiver {
		return NoResult()
	}

	return someResult(receiver.value)
}

func (receiver Option) Unwrap() (Type, bool) {
	if Nothing() == receiver {
		return Type{}, false
	}

	return Type{receiver.value}, true
}
