package model

type Token struct {
	SeqId      uint64  `json:"seq_id"`
	Word       *string `json:"word"`
	Pos        *string `json:"pos"`
	Class      *string `json:"class"`
	WordId     *uint64
	PosId      *uint64
	ClassId    *uint64
	ClassWords []string
}

func (t *Token) Type() TokenType {
	if t.Class != nil {
		return Class
	}
	if t.Word != nil {
		return ExactWord
	}
	if t.Pos != nil {
		return POS
	}
	return Blank
}

type TokenType int

const (
	Blank TokenType = iota
	POS
	ExactWord
	Class
)
