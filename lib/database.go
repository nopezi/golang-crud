package lib

import (
	"database/sql"
	"fmt"
	env "infolelang/lib/env"
)

// Database modal
type Databases struct {
	DB *sql.DB
}

// NewDatabase creates a new database instance
func NewDatabases(env env.Env, zapLogger Logger) Databases {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := sql.Open("mysql", url)

	if err != nil {
		zapLogger.Zap.Info("Url: ", url)
		LogChecklist("Mysql Connection Refused", false)
		zapLogger.Zap.Panic(err)
	}

	// zapLogger.Zap.Info("Database connection established")
	LogChecklist("Mysql Connection Established", true)
	fmt.Println("tes")

	return Databases{
		DB: db,
	}
}
