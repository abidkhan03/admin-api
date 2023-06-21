package dao

import "github.com/spongeling/admin-api/internal/errors"

type WordClass struct {
	WordId  uint64 `gorm:"column:word_id"`
	ClassId uint64 `gorm:"column:class_id"`
}

func (*WordClass) GetID() uint64 {
	return 0
}

func (*WordClass) TableName() string {
	return "word_class"
}

func (wc *WordClass) Validate() error {
	if wc.WordId == 0 {
		return errors.BadRequest("word_id is required")
	} else if wc.ClassId == 0 {
		return errors.BadRequest("class_id is required")
	}
	return nil
}
