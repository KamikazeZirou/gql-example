package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}

func (t *Todo) IsNode() {}

func (t *Todo) GetID() string {
	return t.ID
}
