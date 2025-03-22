package commons

import (
	"ecommerce-api/config"
	"gorm.io/gorm"
)

type Options struct {
	Database *gorm.DB
	Config   *config.Config
}
