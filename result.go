package denary

type Result struct {
	value  string
	err    error
	loaded bool
}

func (receiver Result) errored() bool {
	return receiver.loaded && nil != receiver.err
}

func NoResult() Result {
	return Result{}
}

func (receiver Result) Unwrap() (Type, bool) {
	if NoResult() == receiver {
		return Type{}, false
	}
	if receiver.errored() {
		return Type{}, false
	}

	return Type{receiver.value}, true
}
