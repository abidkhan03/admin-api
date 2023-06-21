package dao

import "github.com/spongeling/admin-api/internal/errors"

type Word struct {
	ID      uint64  `gorm:"column:id"`
	Word    string  `gorm:"column:word"`
	Classes []Class `gorm:"many2many:word_class;"`
	POS     []POS   `gorm:"many2many:word_pos;"`
}

func (w *Word) GetID() uint64 {
	return w.ID
}

func (*Word) TableName() string {
	return "word"
}

func (w *Word) Validate() error {
	if w.Word == "" {
		return errors.BadRequest("word is required")
	}

	return nil
}
