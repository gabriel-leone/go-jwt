package initializers

import (
	"github.com/gabriel-leone/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}
