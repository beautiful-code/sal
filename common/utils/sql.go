package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func NewDBConnection(appConfig *AppConfig) *gorm.DB {
	var err error

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig.DBUser,
		appConfig.DBPwd,
		appConfig.DBProtocol,
		appConfig.DBHost,
		appConfig.DBPort,
		appConfig.DBName,
	)

	conn, err := gorm.Open("mysql", mysqlCredentials)

	if err != nil {
		log.Fatalf("[openDBConnection]: %s\n", err)
	} else {
		log.Println("[openDBConnection]: Connection succesfull.")
	}

	// NewDBConnection returns gorm.DB object to be used for each DB operation.
	return conn
}
