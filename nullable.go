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

func someNullable(value string) Nullable {
	return Nullable{
		loaded: true,
		value:  value,
	}
}

func (receiver Nullable) Return() (Type, error) {
	var nothing Nullable
        if nothing == receiver {
                return Type{}, errNothing
        }
        if Null() == receiver {
                return Type{}, errNull
        }

        return Type{receiver.value}, nil
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
