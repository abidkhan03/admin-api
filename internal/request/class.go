package request

import "github.com/spongeling/admin-api/internal/errors"

type Class struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Words       []string `json:"words"`
}

func (wc *Class) Validate() error {
	if wc.Name == "" {
		return errors.BadRequest("name should not be empty")
	}

	if len(wc.Words) == 0 {
		return errors.BadRequest("words should not be empty")
	}
	return nil
}
