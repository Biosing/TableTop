package enemies

type CreateRequest struct {
	Name         string       `json:"name" example:"Goblin"`
	Level        int          `json:"level" example:"1"`
	Class        string       `json:"class" example:"Warrior"`
	Description  string       `json:"description" example:"A small but fierce warrior"`
	Hp           int          `json:"hp" example:"100"`
	Experience   int          `json:"experience" example:"10"`
	QuantityDeck int          `json:"quantity_deck" example:"2"`
	Defense      int          `json:"defense" example:"2"`
	EnemyMoves   []EnemyMoves `json:"enemy_moves"`
}

type EnemyMoves struct {
	RangeFrom   int    `json:"range_from" example:"1"`
	RangeTo     int    `json:"range_to" example:"10"`
	Description string `json:"description" example:"Close attack"`
}
