package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"table_top/internal/models"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202405241000",
			Migrate: func(tx *gorm.DB) error {
				// Установите расширение uuid-ossp
				if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
					return err
				}
				return tx.AutoMigrate(&models.Character{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("characters")
			},
		},
		{
			ID: "202405251830",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Location{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("locations")
			},
		},
		{
			ID: "202405261522",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Enemy{}, &models.EnemyMove{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("enemy_moves"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("enemies")
			},
		},
	})

	return m.Migrate()
}
