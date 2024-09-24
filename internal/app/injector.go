package injector

import (
	"fmt"
	"innogocartapi/internal/config"
	"innogocartapi/internal/database"

	"go.uber.org/dig"
)

var cartConfiguration config.CartConfiguration
var DigContainer *dig.Container

func initConfiguration() {
	cartConfiguration = *config.NewCartConfiguration()
}

func InitInjector() {
	var error error
	DigContainer = dig.New()

	initConfiguration()

	error = DigContainer.Provide(func(databaseContext database.DatabaseContext) database.DatabaseRepository {
		return database.NewDatabaseRepository(databaseContext, cartConfiguration.Database.DatabaseId)
	})
	if error != nil {
		panic(error)
	}

	error = DigContainer.Provide(func(configuration *config.CartConfiguration) database.DatabaseContext {
		return database.GetDatabaseContext(configuration)
	})
	if error != nil {
		panic(error)
	}

	error = DigContainer.Provide(func() *config.CartConfiguration {
		return &cartConfiguration
	})
	if error != nil {
		panic(error)
	}

	fmt.Println(fmt.Sprintf("Created %s", DigContainer))
}
