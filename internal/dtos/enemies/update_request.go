package enemies

type UpdateRequest struct {
	Name         string `json:"name"`
	Level        int    `json:"level"`
	Class        string `json:"class"`
	Description  string `json:"description"`
	Hp           int    `json:"hp"`
	Experience   int    `json:"experience"`
	QuantityDeck int    `json:"quantity_deck"`
	Defense      int    `json:"defense"`
}
