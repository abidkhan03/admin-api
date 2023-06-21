package dao

type WordPos struct {
	WordId  uint64 `gorm:"column:word_id"`
	LemmaId uint64 `gorm:"column:lemma_id"`
	PosId   uint64 `gorm:"column:pos_id"`
}

func (*WordPos) GetID() uint64 {
	return 0
}

func (*WordPos) TableName() string {
	return "word_pos"
}

func (wp *WordPos) Validate() error {
	return nil
}
