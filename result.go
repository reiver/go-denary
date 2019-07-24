package denary

type Result struct {
	value  string
	err    error
	loaded bool
}

func (receiver Result) errored() bool {
	return receiver.loaded && nil != receiver.err
}

