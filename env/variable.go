package env

import (
	"os"
)

var (
	// RethinkHostAddress .
	RethinkHostAddress = ":28015"

	// DatabaseName .
	DatabaseName = "objectdb"
)

// LoadVariables gets environment variables: enableProdMode and serverHost.
func LoadVariables() (bool, string) {
	var (
		enableProdMode = true
		serverHost     = ":9000"
	)
	if os.Getenv("GBP_PROMODE") == "false" {
		enableProdMode = false
	}

	host := os.Getenv("GBP_HOST")
	if len(host) > 0 {
		serverHost = host
	}

	rethinkHost := os.Getenv("GBP_RETHINK_HOST")
	if len(rethinkHost) > 0 {
		RethinkHostAddress = rethinkHost
	}

	dbName := os.Getenv("GBP_DB_NAME")
	if len(rethinkHost) > 0 {
		DatabaseName = dbName
	}

	return enableProdMode, serverHost
}
