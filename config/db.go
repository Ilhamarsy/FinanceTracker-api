package config

import (
	"finance-tracker/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global DB variable to access the database connection throughout the app

type Database struct {
	env *Config
}

// NewDatabase initializes a new Database instance with environment configurations
func NewDatabase(cfg *Config) *Database {
	return &Database{
		env: cfg,
	}
}

// ConnectDB establishes the connection to the database and migrates models
func (db *Database) ConnectDB() {
	// Create the DSN (Data Source Name) for Postgres connection
	dsn := "host=" + db.env.DBHost + " user=" + db.env.DBUser + " password=" + db.env.DBPassword + " dbname=" + db.env.DBName + " port=" + db.env.DBPort + " sslmode=disable"

	// Open the database connection
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Store the connection in the global variable
	DB = connection

	// Automatically migrate the models (tables creation)
	err = DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Income{}, &models.Expense{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database connected and models migrated successfully.")
}
