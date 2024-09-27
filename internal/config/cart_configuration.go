package config

import (
	"encoding/json"
	"os"
)

type cartConfigurationType struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Database struct {
		DatabaseId int    `json:"databaseid"`
		Name       string `json:"databasename"`
		User       string `json:"user"`
		Pass       string `json:"pass"`
		SSLMode    string `json:"sslmode"`
		Host       string `json:"host"`
	} `json:"database"`
}

type CartConfigurationPonterType *cartConfigurationType

var cartConfiguration CartConfigurationPonterType

func NewCartConfiguration() CartConfigurationPonterType {
	var result cartConfigurationType

	if cartConfiguration == nil {
		file, error := os.ReadFile("C:\\GODev\\InnoGOCartAPI\\internal\\config\\cart_configuration.json")
		if error == nil {
			error = json.Unmarshal(file, &result)

			if error != nil {
				panic(error)
			}

			cartConfiguration = &result
		} else {
			panic(error)
		}
	} else {
		result = *cartConfiguration
	}

	return &result
}
