package config

import (
	"github.com/gobjserver/gobjserver/core/gateway"
	"github.com/gobjserver/gobjserver/db"
)

// CreateObjectGateway .
func CreateObjectGateway() gateway.ObjectGateway {
	var database db.CommonDatabase
	database = CreateRethinkDB()
	rethinkObjectGateway := db.ObjectGatewayImpl{
		Datatbase: database,
	}
	return rethinkObjectGateway
}

// CreateRethinkDB .
func CreateRethinkDB() *db.RethinkDB {
	var database db.RethinkDB
	database.Session = database.Connect()
	return &database
}
