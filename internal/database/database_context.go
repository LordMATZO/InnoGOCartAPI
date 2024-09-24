package database

import (
	"innogocartapi/internal/config"
)

type DatabaseContext interface {
	initConnection() error
}

var databaseContextInstance DatabaseContext

const (
	PostgreDatabaseId = iota
)

func GetDatabaseContext(configuration *config.CartConfiguration) DatabaseContext {
	if databaseContextInstance == nil {
		switch configuration.Database.DatabaseId {
		case PostgreDatabaseId:
			databaseContextInstance = newPostgreContext(configuration)
			databaseContextInstance.initConnection()
		}
	}

	return databaseContextInstance
}
