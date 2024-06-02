package enemy_moves

type UpdateRequest struct {
	RangeFrom   int    `json:"range_from"`
	RangeTo     int    `json:"range_to"`
	Description string `json:"description"`
}
