package injector

import (
	"encoding/json"
	"innogocartapi/internal/config"
	"innogocartapi/internal/database"
	"os"

	"go.uber.org/dig"
)

var cartConfiguration config.CartConfiguration

func initConfiguration() {
	cartConfigurationFile, error := os.Open("C:\\GODev\\InnoGOCartAPI\\internal\\config\\cart_configuration.json")
	var cartConfigurationFileContent []byte

	if error == nil {
		error = json.Unmarshal(cartConfigurationFileContent, &cartConfiguration)
		if error != nil {
			panic(error)
		}
	} else {
		panic(error)
	}

	defer cartConfigurationFile.Close()
}

func InitInjector() {
	container := dig.New()
	initConfiguration()
	error := container.Provide(func() *config.CartConfiguration {
		return &cartConfiguration
	})
	if error != nil {
		panic(error)
	}

	error = container.Provide(func(configuration *config.CartConfiguration) database.DatabaseContextReference {
		return database.GetDatabaseContext(configuration)
	})
	if error != nil {
		panic(error)
	}

	error = container.Provide(func(databaseContext database.DatabaseContextReference) database.DatabaseRepository {
		return database.NewDatabaseRepository(databaseContext, cartConfiguration.Database.DatabaseId)
	})
	if error != nil {
		panic(error)
	}
}
