package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type Config struct {
	Postgres DBConfig `json:"postgres"`
	MySQL    DBConfig `json:"mysql"`
}

func LoadDBConfig() Config {
	var config Config
	file, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Fatal("Error reading database.json file:", err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error unmarshalling database.json file:", err)
	}
	return config
}
