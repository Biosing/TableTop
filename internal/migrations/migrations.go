package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"table_top/internal/models/characters"
	"table_top/internal/models/combo_cards"
	"table_top/internal/models/enemies"
	"table_top/internal/models/games"
	"table_top/internal/models/items"
	"table_top/internal/models/locations"
	"table_top/internal/models/users"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202406021200",
			Migrate: func(tx *gorm.DB) error {
				// Install the uuid-ossp extension
				if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
					return err
				}

				// Auto migrate the models
				return tx.AutoMigrate(
					&users.User{},
					&games.Backpack{},
					&characters.Character{},
					&items.Weapon{},
					&combo_cards.ComboCard{},
					&combo_cards.DamageComboCard{},
					&games.Deck{},
					&enemies.Enemy{},
					&enemies.EnemyMove{},
					&games.GameSession{},
					&locations.Location{},
					&games.Player{},
					&combo_cards.SpecialEffect{},
					&items.WeaponCombo{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					"users",
					"backpacks",
					"characters",
					"weapons",
					"combo_cards",
					"damage_combo_cards",
					"decks",
					"enemies",
					"enemy_moves",
					"game_sessions",
					"locations",
					"players",
					"special_effects",
					"weapon_combos",
				)
			},
		},
	})

	return m.Migrate()
}
