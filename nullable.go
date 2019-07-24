package denary

import (
	"fmt"
	"reflect"
	"unsafe"
)

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

func (receiver Nullable) GoString() string {
	if Nothing().Nullable() == receiver {
		return "denary.Nothing().Nullable()"
	}

	if Null() == receiver {
		return "denary.Null()"
	}

	return fmt.Sprintf("denary.Parse(%q).NullError()", receiver.value)
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

func (receiver *Nullable) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	stringHeader := reflect.StringHeader{Data: sliceHeader.Data, Len: sliceHeader.Len}
	var str string = *(*string)(unsafe.Pointer(&stringHeader))

	switch str {
	case "null":
		*receiver = Null()
		return nil
	default:
		n, err := Parse(str).Return()
		if nil != err {
			return err
		}

		*receiver = someNullable(n.String())
		return nil
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
