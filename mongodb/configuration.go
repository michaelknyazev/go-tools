package database

import "os"

type DatabaseConfig struct {
	Name string
	URI  string
}

var configuration DatabaseConfig

func LoadConfiguration() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	var mongoUri string = "mongodb://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName

	configuration = DatabaseConfig{dbName, mongoUri}
}
