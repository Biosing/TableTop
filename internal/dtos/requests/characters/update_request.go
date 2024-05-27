package characters

type UpdateRequest struct {
	Name         string `json:"name"`
	Class        string `json:"class"`
	Race         string `json:"race"`
	Age          int    `json:"age"`
	Sex          string `json:"sex"`
	Description  string `json:"description"`
	BeginningWay string `json:"beginning_way"`
}
