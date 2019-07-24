package denary

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
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

func (receiver Nullable) MarshalJSON() ([]byte, error) {
	if Nothing().Nullable() == receiver {
		return nil, errNothing
	}
	if Null() == receiver {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver Nullable) Return() (Type, error) {
        if Nothing().Nullable() == receiver {
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
	if nil == data {
		return fmt.Errorf("denary: %#v is invalid JSON", data)
	}

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	stringHeader := reflect.StringHeader{Data: sliceHeader.Data, Len: sliceHeader.Len}
	var str string = *(*string)(unsafe.Pointer(&stringHeader))

	if 0 >= len(str) || `"` == str {
		return fmt.Errorf("denary: %q is invalid JSON", data)
	}

	switch {
	case "null" == str:
		*receiver = Null()
		return nil
	case strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`):
		return receiver.UnmarshalJSON(data[1:len(data)-1])
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
	if Nothing().Nullable() == receiver {
		return Type{}, false
	}
	if Null() == receiver {
		return Type{}, false
	}

	return Type{receiver.value}, true
}
