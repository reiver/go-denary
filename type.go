package denary

import (
	"strconv"
)

type Type struct {
	value string
}

func Int8(n int8) Type {
	value := strconv.FormatInt(int64(n), 10)

	return Type{
		value: value,
	}
}

func Int16(n int16) Type {
	value := strconv.FormatInt(int64(n), 10)

	return Type{
		value: value,
	}
}

func Int32(n int32) Type {
	value := strconv.FormatInt(int64(n), 10)

	return Type{
		value: value,
	}
}

func Int64(n int8) Type {
	value := strconv.FormatInt(int64(n), 10)

	return Type{
		value: value,
	}
}

func Uint8(n uint8) Type {
	value := strconv.FormatUint(uint64(n), 10)

	return Type{
		value: value,
	}
}

func Uint16(n uint16) Type {
	value := strconv.FormatUint(uint64(n), 10)

	return Type{
		value: value,
	}
}

func Uint32(n uint32) Type {
	value := strconv.FormatUint(uint64(n), 10)

	return Type{
		value: value,
	}
}

func Uint64(n uint64) Type {
	value := strconv.FormatUint(uint64(n), 10)

	return Type{
		value: value,
	}
}

func (receiver Type) String() string {
	return receiver.value
}

