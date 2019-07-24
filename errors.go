package denary

import (
	"errors"
)

var (
	errNilReceiver = errors.New("denary: Nil Receiver")
	errNoResult    = errors.New("denary: No Result")
	errNothing     = errors.New("denary: Nothing")
	errNull        = errors.New("denary: Null")
)
