package models

import (
	"strconv"
)

type Note struct {
	ID	 uint    `json:"id"`
	Markdown []byte  `json:"note"`
}

func (n Note) MarshalJSON() ([]byte, error) {
	id := `"` + strconv.Itoa(int(n.ID)) + `"`
	markdown := `"` + string(n.Markdown) + `"`

	return []byte(`{` + `"id":` + id + `,` + `"note":` + markdown + `}`), nil
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
