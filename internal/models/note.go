package models

type Note struct {
	ID		 int	`json:"id"`
	Markdown []byte `json:"note"`
}
