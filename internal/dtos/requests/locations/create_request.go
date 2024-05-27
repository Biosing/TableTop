package locations

type CreateRequest struct {
	Name          string `json:"name"`
	Level         int    `json:"level"`
	DangerLevel   int    `json:"danger_level"`
	Description   string `json:"description"`
	MonstersCount int    `json:"monsters_count"`
}
