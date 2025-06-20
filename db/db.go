package db

import (
	"log"
	"os"

	"github.com/sidharth-chauhan/imdb-cli/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // global DB connection

// InitDB initializes the SQLite database and performs auto-migration
func InitDB() {
	var err error

	// Ensure the data directory exists
	//
	if err = os.MkdirAll("data", 0755); err != nil {
		log.Fatal("❌ Failed to create data directory:", err)
	}
	//
	// Open SQLite connection (will create file if not exists)
	DB, err = gorm.Open(sqlite.Open("data/favorites.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto-migrate the Favorite model
	err = DB.AutoMigrate(&models.Favorite{})
	if err != nil {
		log.Fatal("❌ Failed to auto-migrate database:", err)
	}

	log.Println("✅ Connected to SQLite database")
}

// GetDB returns the current database connection
func GetDB() *gorm.DB {
	return DB
}
