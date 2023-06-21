package dao

import (
	"github.com/spongeling/admin-api/internal/errors"
)

type CategoryExample struct {
	CategoryId    uint64 `gorm:"column:category_id;uniqueIndex"`
	Category      *Category
	Rule          *string `gorm:"column:rule"`
	Tip           *string `gorm:"column:tip"`
	Phrase        string  `gorm:"column:phrase"`
	PatternId     uint64  `gorm:"column:pattern_id;uniqueIndex"`
	Pattern       *Pattern
	FullPatternId uint64 `gorm:"column:full_pattern_id"`
	FullPattern   *Pattern
}

func (p *CategoryExample) GetID() uint64 {
	return p.CategoryId
}

func (*CategoryExample) TableName() string {
	return "category_example"
}

func (p *CategoryExample) Validate() error {
	if p.Phrase == "" {
		return errors.BadRequest("phrase is required")
	}
	return nil
}
