package dao

import (
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/internal/pos"
)

type POS struct {
	ID                 uint64 `gorm:"column:id"`
	Category           rune   `gorm:"column:category"`
	Type               *rune  `gorm:"column:type"`
	Degree             *rune  `gorm:"column:degree"`
	Gender             *rune  `gorm:"column:gen"`
	Number             *rune  `gorm:"column:num"`
	PossessorPerson    *rune  `gorm:"column:possessorpers"`
	PossessorNumber    *rune  `gorm:"column:possessornum"`
	Person             *rune  `gorm:"column:person"`
	NounClass          *rune  `gorm:"column:neclass"`
	NounSubClass       *rune  `gorm:"column:nesubclass"`
	Case               *rune  `gorm:"column:case"`
	Polite             *rune  `gorm:"column:polite"`
	Mood               *rune  `gorm:"column:mood"`
	Tense              *rune  `gorm:"column:tense"`
	PunctuationEnclose *rune  `gorm:"column:punctenclose"`
}

func (p *POS) GetID() uint64 {
	return p.ID
}

func (*POS) TableName() string {
	return "pos"
}

func (p *POS) Validate() error {
	if p.Category == 0 {
		return errors.BadRequest("category is required")
	}

	return nil
}

func (p *POS) ToPOS() *pos.POS {
	cat := pos.GetCategory(p.Category)

	res := &pos.POS{
		Category: cat,
	}

	res.Values = make(map[pos.Field]rune)

	for _, field := range cat.Fields {
		v := p.valueFromField(field)
		if v != nil {
			res.Values[field] = *v
		}
	}

	return res
}

func (p *POS) valueFromField(field pos.Field) *rune {
	switch field {
	case pos.Type:
		return p.Type
	case pos.Degree:
		return p.Degree
	case pos.Gender:
		return p.Gender
	case pos.Nmber:
		return p.Number
	case pos.PossessorNumber:
		return p.PossessorNumber
	case pos.PossessorPerson:
		return p.PossessorPerson
	case pos.Person:
		return p.Person
	case pos.NounClass:
		return p.NounClass
	case pos.NounSubClass:
		return p.NounSubClass
	case pos.Case:
		return p.Case
	case pos.Polite:
		return p.Polite
	case pos.Mood:
		return p.Mood
	case pos.Tense:
		return p.Tense
	case pos.PunctuationEnclose:
		return p.PunctuationEnclose
	}

	return nil
}

func POSFromString(s string) (*POS, error) {
	p, err := pos.FromString(s)
	if err != nil {
		return nil, err
	}

	res := POS{
		Category: p.Category.Identifier,
	}

	for field, value := range p.Values {
		v := value
		switch field {
		case pos.Type:
			res.Type = &v
		case pos.Degree:
			res.Degree = &v
		case pos.Gender:
			res.Gender = &v
		case pos.Nmber:
			res.Number = &v
		case pos.PossessorNumber:
			res.PossessorNumber = &v
		case pos.PossessorPerson:
			res.PossessorPerson = &v
		case pos.Person:
			res.Person = &v
		case pos.NounClass:
			res.NounClass = &v
		case pos.NounSubClass:
			res.NounSubClass = &v
		case pos.Case:
			res.Case = &v
		case pos.Polite:
			res.Polite = &v
		case pos.Mood:
			res.Mood = &v
		case pos.Tense:
			res.Tense = &v
		case pos.PunctuationEnclose:
			res.PunctuationEnclose = &v
		}
	}

	return &res, nil
}
