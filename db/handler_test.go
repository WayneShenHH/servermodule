package db

import (
	"testing"

	"github.com/WayneShenHH/servermodule/config"
)

const info = `
{
	"firstName": "John",
	"lastName" : "doe",
	"age"      : 26,
	"address"  : {
	  "streetAddress": "naist street",
	  "city"         : "Nara",
	  "postalCode"   : "630-0192"
	},
	"phoneNumbers": [
	  {
		"type"  : "iPhone",
		"number": "0123-4567-8888"
	  },
	  {
		"type"  : "home",
		"number": "0123-4567-8910"
	  }
	],
	"banks": [
		{
			"name":"IBank",
			"balance":"100"
		},
		{
			"name":"CityBank",
			"balance":"890"
		}
	]
  }`

func Test_NestedUpdate(t *testing.T) {
	db := NewArango(&config.DatabaseConfig{
		Host: "http://localhost:8529",
		Name: "Database",
	})
	db.NestedUpdate("info", "1.01", map[string]string{
		"age":                               "age + 1",
		"banks[?(@.balance > 200)].balance": "balance * 100",
		// "banks[?(banks.balance > 200)].balance": "balance * 100",
		// "banks[*].balance": "balance * 100",
	})
}
