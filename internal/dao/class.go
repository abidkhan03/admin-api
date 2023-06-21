package dao

import "github.com/spongeling/admin-api/internal/errors"

type Class struct {
	Id          uint64  `gorm:"column:id"`
	Name        string  `gorm:"column:name"`
	WordId      *uint64 `gorm:"column:word_id"`
	Description string  `gorm:"column:description"`
	Words       []Word  `gorm:"many2many:word_class;"`
}

func (c *Class) GetID() uint64 {
	return c.Id
}

func (*Class) TableName() string {
	return "class"
}

func (c *Class) Validate() error {
	if c.Name == "" {
		return errors.BadRequest("name is required")
	}
	return nil
}
