package database

import (
	"fmt"
	"innogocartapi/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgreContext struct {
	db            *sqlx.DB
	configuration *config.CartConfiguration
}

func newPostgreContext(configuration *config.CartConfiguration) *postgreContext {
	postgreContext := new(postgreContext)
	postgreContext.configuration = configuration

	return postgreContext
}

func (driver *postgreContext) initConnection() error {
	db, error := sqlx.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s",
		driver.configuration.Database.User,
		driver.configuration.Database.Name,
		driver.configuration.Database.SSLMode,
		driver.configuration.Database.Pass,
		driver.configuration.Database.Host))

	if error == nil {
		driver.db = db
	} else {
		panic(error)
	}

	return error
}
