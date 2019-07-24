package denary

type Result struct {
	value  string
	err    error
	loaded bool
}

func (receiver Result) errored() bool {
	return receiver.loaded && nil != receiver.err
}

func someResult(value string) Result {
	return Result{
		loaded: true,
		value:  value,
	}
}

func (receiver Result) Return() (Type, error) {
	if Nothing().Result() == receiver {
		return Type{}, errNoResult
	}
	if receiver.errored() {
		return Type{}, receiver.err
	}

	return Type{receiver.value}, nil
}

func (receiver Result) Validate() error {
	_, err := receiver.Return()
	return err
}

func (receiver Result) Unwrap() (Type, bool) {
	if Nothing().Result() == receiver {
		return Type{}, false
	}
	if receiver.errored() {
		return Type{}, false
	}

	return Type{receiver.value}, true
}
