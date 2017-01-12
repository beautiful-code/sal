package common

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDBConnection() *gorm.DB {
	var err error

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		AppConfig.DBUser,
		AppConfig.DBPwd,
		AppConfig.DBProtocol,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
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
