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

For this example, because ‘Purchase.Amount’ is a ‘denary.Nullable’, it can handle JSON values such as:

	{
		"merchant_name": "Food Is Us",
		"amount": null,
	}

In that example we assigned the JSON ‘null’ to ‘Purchase.Amount’.

And:

	{
		"merchant_name": "Food Is Us",
		"amount": 14,
	}

In that example we assigned the JSON number ‘14’ to ‘Purchase.Amount’.

And:

	{
		"merchant_name": "Food Is Us",
		"amount": 2.31,
	}

In that example we assigned the JSON number ‘2.31’ to ‘Purchase.Amount’.

And:

	{
		"merchant_name": "Food Is Us",
		"amount": "14",
	}

In that example we assigned the JSON string ‘"14"’ to ‘Purchase.Amount’.

And:

	{
		"merchant_name": "Food Is Us",
		"amount": 2.31,
	}

In that example we assigned the JSON string ‘"2.31"’ to ‘Purchase.Amount’.

If you wanted to create a ‘Purchase.Amount’ literal, you could use ‘denary.Parse().NullError()’
to do that:

	purchase := Purchase{
		MerchanName: "Super Mega Store",
		Amount:      denary.Parse("13.21").NullError(),
	}
*/
package denary
