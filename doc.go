/*
Package denary provides base-10 numbers.

This type is safe for using with money.


JSON Data Format

Package denary provides support for the JSON data format, and in particular for the built-in "encoding/json" package,
via ‘denary.Nullable’.

So, for example, it could be used as:

	type Purchase struct {
		MerchanName string          `json:"merchant_name"`
		Amount      denary.Nullable `json:"amount"`
	}

In this example, the ‘Purchase.Amount’ field is used to safely store a money amount.

*/
package denary
