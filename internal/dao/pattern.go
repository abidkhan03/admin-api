package dao

import (
	"github.com/lib/pq"
	"github.com/spongeling/admin-api/internal/errors"
)

type Pattern struct {
	Id       uint64        `gorm:"column:id"`
	TokenIds pq.Int64Array `gorm:"column:token_ids; type:integer[]"`
}

func (p *Pattern) GetID() uint64 {
	return p.Id
}

func (*Pattern) TableName() string {
	return "pattern"
}

func (p *Pattern) Validate() error {
	if len(p.TokenIds) < 1 {
		return errors.BadRequest("at least one token is required")
	}
	return nil
}
