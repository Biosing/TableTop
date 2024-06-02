package enemy_moves

type CreateRequest struct {
	RangeFrom   int    `json:"range_from" example:"1"`
	RangeTo     int    `json:"range_to" example:"10"`
	Description string `json:"description" example:"Close attack"`
}
