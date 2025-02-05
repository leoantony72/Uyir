package database

import (
	"log"

	"github.com/leoantony72/Uyir/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartPostgres() *gorm.DB {
	// Hardcoded PostgreSQL connection details
	dbUser := "postgres"         // Username for PostgreSQL
	dbPassword := "yourpassword" // Password for PostgreSQL
	dbName := "uyir"             // Database name
	dbHost := "localhost"        // Host name (service name in Docker Compose)
	dbPort := "5432"             // Default PostgreSQL port

	// Create the PostgreSQL connection string
	dbUrl := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	// Perform migrations
	model.Migrate(db)
	model.MigrateSession(db)
	model.MigrateReport(db)
	return db
}
