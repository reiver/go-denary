package denary

type Nullable struct {
	value  string
	isnull bool
	loaded bool
}

func Null() Nullable {
	return Nullable{
		loaded: true,
		isnull: true,
	}
}

func (receiver Nullable) Unwrap() (Type, bool) {
	var nothing Nullable
	if  nothing == receiver {
		return Type{}, false
	}
	if Null() == receiver {
		return Type{}, false
	}

	return Type{receiver.value}, true
}
