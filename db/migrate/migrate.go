package migrate

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// Migrate Migrate
func Migrate(db *gorm.DB) error {
	migrate := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		CreateVideos(),
	})

	return migrate.Migrate()
}
