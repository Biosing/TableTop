package characters

type UpdateRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Race         string `json:"race"`
	Age          int    `json:"age"`
	Sex          string `json:"sex"`
	Description  string `json:"description"`
	BeginningWay string `json:"beginning_way"`
}
