package migrate

import (
	"github.com/jinzhu/gorm"
	"github.com/tett23/mizuaoi-v2/models"
	"gopkg.in/gormigrate.v1"
)

// CreateVideos Create videos table
func CreateVideos() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20190129000851",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Video{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("videos").Error
		},
	}
}

