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
