package denary

import (
	"math/big"
)

// Option is an ‘option type’ for ‘denary.Type’.
//
// It can contain:
//
// • ‘nothing’, or
//
// • ‘something’.
//
//
// Nothing
//
// An uninitialize variable, of type ‘denary.Option’, contains ‘nothing’. I.e.,:...
//
//	var option denary.Option
//
// You can determine if a variable of type ‘denary.Option’ contains ‘nothing’, or not,
// by using the ‘denary.Nothing()’ function.
//
// For example:
//
//	if denary.Nothing() == option {
//		//@TODO
//	}
//
// Or:
//
//	switch option {
//	case denary.Nothing():
//		//@TODO
//	default:
//		//@TODO
//	}
//
// You can also use the ‘denary.Nothing()’ function to give a variable of type
// ‘denary.Option’ the value of nothing ‘nothing’:
//
//	option = denary.Nothing()
//
//
// Something
//
// You can create a ‘denary.Option’ with ‘something’ in it by turning a ‘denary.Type’
// into a ‘denary.Option’.
//
// For example:
//
//	var value denary.Type = denary.Int64(-12345)
//	
//	var option denary.Option = value.Option()
//
// Or, more compactly:
//
//	var option denary.Option = denary.Int64(-12345).Option()
//
// There are a number of other functions you can use in addition to this, to create an ‘denary.Option’
// with ‘something’ in it:
//
// • denary.Parse()
//
// • denary.Int8()
//
// • denary.Int16()
//
// • denary.Int32()
//
// • denary.Int64()
//
// • denary.Uint8()
//
// • denary.Uint16()
//
// • denary.Uint32()
//
// • denary.Uint64()
//
// You can determine if a variable of type ‘denary.Option’ contains a specific ‘something’, or not,
// with code like the following:
//
// For example:
//
//	if denary.Int64(-12345).Option() == option {
//		//@TODO
//	}
//
// Or:
//
//	switch option {
//	case denary.Int64(-12345).Option():
//		//@TODO
//	default:
//		//@TODO
//	}
type Option struct {
	value  string
	loaded bool
}

// Nothing returns an empty denary.Option.
func Nothing() Option {
	return Option{}
}

func something(value string) Option {
	return Option{
		loaded: true,
		value:  value,
	}
}

func (receiver Option) BigRat() (*big.Rat, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return Type{receiver.value}.BigRat()
}

// Else defaults this ‘denary.Option’ to ‘value’ if this ‘denary.Option’ has a value of ‘denary.Nothing()’,
// else it just returns itself as is.
func (receiver Option) Else(value Type) Option {
	if Nothing() == receiver {
		return something(value.String())
	}

	return receiver
}

// Nullable returns the equivalent ‘denary.Nullable’ for this ‘denary.Option’.
func (receiver Option) Nullable() Nullable {
	if Nothing() == receiver {
		var nothing Nullable
		return nothing
	}

	return someNullable(receiver.value)
}

// Nullable returns the equivalent ‘denary.Result’ for this ‘denary.Option’.
func (receiver Option) Result() Result {
	if Nothing() == receiver {
		return Result{}
	}

	return someResult(receiver.value)
}

// Return returns the ‘denary.Type’ inside, if there is one inside.
func (receiver Option) Return() (Type, error) {
	if Nothing() == receiver {
		return Type{}, errNothing
	}

	return Type{receiver.value}, nil
}

// Unwrap returns the ‘denary.Type’ inside, if there is one inside.
func (receiver Option) Unwrap() (Type, bool) {
	if Nothing() == receiver {
		return Type{}, false
	}

	return Type{receiver.value}, true
}
