package database

import (
	"innogocartapi/internal/config"
)

type databaseContext interface {
	initConnection() error
}

type DatabaseContextReference *databaseContext

var databaseContextInstance databaseContext

const (
	PostgreDatabaseId = iota
)

func GetDatabaseContext(configuration *config.CartConfiguration) *databaseContext {
	if databaseContextInstance == nil {
		switch configuration.Database.DatabaseId {
		case PostgreDatabaseId:
			databaseContextInstance = newPostgreContext(configuration)
			databaseContextInstance.initConnection()
		}
	}

	return &databaseContextInstance
}
