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
		{
			ID: "202405270014",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Weapon{}, &models.WeaponCombo{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("weapon"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("weapon_combo")
			},
		},
		{
			ID: "202405270104",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Weapon{}, &models.WeaponCombo{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("weapons", "weapon_combos")
			},
		},
		{
			ID: "202405272325",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Enemy{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("enemy")
			},
		},
		{
			ID: "202405271001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.ComboCard{}, &models.DamageComboCard{}, &models.SpecialEffect{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("combo_cards", "damage_combo_cards", "special_effects")
			},
		},
	})

	return m.Migrate()
}
