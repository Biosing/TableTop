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
		// Добавьте новые миграции здесь
	})

	return m.Migrate()
}
