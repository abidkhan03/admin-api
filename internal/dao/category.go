package dao

import "github.com/spongeling/admin-api/internal/errors"

type Category struct {
	Id       uint64           `json:"id"`
	Name     string           `json:"name"`
	ParentId *uint64          `json:"parent_id"`
	Parent   *Category        `json:"parent"`
	Example  *CategoryExample `json:"example"`
}

func (c *Category) GetID() uint64 {
	return c.Id
}

func (_ *Category) TableName() string {
	return "category"
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.BadRequest("category name is required")
	}
	return nil
}
