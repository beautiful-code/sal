package common

// StartUp bootstrapps the application
func StartUp(configFilePath string) {
	// Initialize AppConfig variable
	initConfig(configFilePath)
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))
	// Start a DB session
	createDBSession()
}
