package database

import (
	"log"

	"github.com/rajpatelbot/icollab/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := utils.GenerateDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("⚠️ Failed to connect to database: %v", err)

	}

	return db
}
