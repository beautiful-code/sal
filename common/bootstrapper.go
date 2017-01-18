package common

// StartUp bootstrapps the application
func StartUp(configFilePath string) {
	// Initialize AppConfig variable
	InitConfig(configFilePath)
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Initialize Logger objects with Log Level
	SetLogLevel(Level(AppConfig.LogLevel))
	// Start a DB session
	CreateDBSession()
}
