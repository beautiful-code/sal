package utils

import (
	"encoding/json"
	"log"
	"os"
)

type (
	AppConfig struct {
		Server, DBHost, DBUser, DBPwd, DBName, DBProtocol, DBPort string
		LogLevel                                                  int
		UserService                                               string
		AllowedOrigins                                            []string
	}
)

func LoadAppConfig(configFilePath string, appConfig *AppConfig) {
	file, err := os.Open(configFilePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(appConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
