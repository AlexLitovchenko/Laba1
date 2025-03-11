package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db string `json:"db"`
}

var config Config

func InitConfig() {

	config = Config{}

	workDir := os.Getenv("PWD")
	if workDir == "" {
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Open(workDir + "/go-rest-api/config/settings.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return config
}
