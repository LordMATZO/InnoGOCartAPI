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

	file, _ := os.ReadFile("configuration.json")
	json.Unmarshal(file, &result)

	return &result
}
