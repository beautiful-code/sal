package common

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var session *gorm.DB

func getSession() *gorm.DB {
	if session == nil {
		CreateDBSession()
	}

	return session
}

func CreateDBSession() {
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

	session, err = gorm.Open("mysql", mysqlCredentials)

	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

// DataStore for MySQL
type DataStore struct {
	Session *gorm.DB
}

// Close closes a gorm.DB value.
// Used to add defer statements for closing the copied session.
func (ds *DataStore) Close() {
	ds.Session.Close()
	session = nil
}

// NewDataStore creates a new DataStore object to be used for each HTTP request.
func NewDataStore() *DataStore {
	session := getSession()

	dataStore := &DataStore{
		Session: session,
	}

	return dataStore
}
