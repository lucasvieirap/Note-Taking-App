package models

type Note struct {
	ID	 uint    `json:"id"`
	Markdown []byte  `json:"note"`
}

// TODO: Note Validation
func (n Note) Validate() bool {
	if n.Markdown == nil {
		return false
	}
	if n.ID <= 0 {
		return false
	}
	return true
}
