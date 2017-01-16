package common

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// StartUp bootstrapps the application
func StartUp(configFilePath string) {
	// Initialize AppConfig variable
	initConfig(configFilePath)
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))
	// open a DB connection
	DB = NewDBConnection()
}
