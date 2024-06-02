package game_sessions

type ListRequest struct {
	Name  string `form:"name" json:"name"`
	Page  int    `form:"page" json:"page"`
	Limit int    `form:"limit" json:"limit"`
}
