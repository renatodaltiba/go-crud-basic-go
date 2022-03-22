package migrations

import (
	"github.com/renatodaltiba/rest-api-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
