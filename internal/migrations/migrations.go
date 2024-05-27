package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	characterModels "table_top/internal/models/characters"
	comboCardModels "table_top/internal/models/combo_cards"
	enemyModels "table_top/internal/models/enemies"
	itemModels "table_top/internal/models/items"
	locationModels "table_top/internal/models/locations"
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
				return tx.AutoMigrate(&characterModels.Character{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("characters")
			},
		},
		{
			ID: "202405251830",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&locationModels.Location{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("locations")
			},
		},
		{
			ID: "202405261522",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&enemyModels.Enemy{}, &enemyModels.EnemyMove{})
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
				return tx.AutoMigrate(&itemModels.Weapon{}, &itemModels.WeaponCombo{})
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
				return tx.AutoMigrate(&itemModels.Weapon{}, &itemModels.WeaponCombo{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("weapons", "weapon_combos")
			},
		},
		{
			ID: "202405272325",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&enemyModels.Enemy{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("enemy")
			},
		},
		{
			ID: "202405271001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&comboCardModels.ComboCard{}, &comboCardModels.DamageComboCard{}, &comboCardModels.SpecialEffect{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("combo_cards", "damage_combo_cards", "special_effects")
			},
		},
		{
			ID: "202405280300",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&comboCardModels.ComboCard{}, &comboCardModels.DamageComboCard{}, &comboCardModels.SpecialEffect{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("combo_cards", "damage_combo_cards", "special_effects")
			},
		},
		{
			ID: "202405280305",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&itemModels.Weapon{}, &itemModels.WeaponCombo{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("weapons", "weapon_combos")
			},
		}, {
			ID: "202405280306",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&comboCardModels.ComboCard{}, &comboCardModels.DamageComboCard{}, &comboCardModels.SpecialEffect{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("combo_cards", "damage_combo_cards", "special_effects")
			},
		},
		{
			ID: "202405280307",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&itemModels.Weapon{}, &itemModels.WeaponCombo{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("weapons", "weapon_combos")
			},
		},
	})

	return m.Migrate()
}
