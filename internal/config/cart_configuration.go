package config

import (
	"encoding/json"
	"os"
)

type CartConfiguration struct {
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

func NewCartConfiguration() *CartConfiguration {
	var result CartConfiguration

	file, error := os.ReadFile("C:\\GODev\\InnoGOCartAPI\\internal\\config\\cart_configuration.json")
	if error == nil {
		error = json.Unmarshal(file, &result)

		if error != nil {
			panic(error)
		}
	} else {
		panic(error)
	}

	return &result
}
