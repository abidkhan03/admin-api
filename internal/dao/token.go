package dao

import "fmt"

const NullTokenId = 1

type Token struct {
	Id      uint64  `gorm:"column:id"`
	PosId   *uint64 `gorm:"column:pos_id"`
	Pos     *POS
	WordId  *uint64 `gorm:"column:word_id"`
	Word    *Word
	ClassId *uint64 `gorm:"column:class_id"`
	Class   *Class
}

func (t *Token) GetID() uint64 {
	return t.Id
}

func (*Token) TableName() string {
	return "token"
}

func (*Token) Validate() error {
	return nil
}

func (t *Token) IsBlank() bool {
	return t.WordId == nil && t.PosId == nil && t.ClassId == nil
}

func (t *Token) String() string {
	if t.Class != nil {
		return fmt.Sprintf("<%s>", t.Class.Name)
	}
	if t.Word != nil {
		return t.Word.Word
	}
	if t.Pos != nil {
		return t.Pos.ToPOS().String()
	}
	return "BLANK"
}
