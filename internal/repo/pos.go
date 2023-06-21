package repo

import (
	"context"
	"github.com/spongeling/admin-api/internal/dao"
)

func (r *Repo) GetPosId(ctx context.Context, pos *dao.POS) (uint64, error) {
	err := r.db(ctx).Where(pos).FirstOrCreate(pos).Error
	return pos.ID, wrap(err)
}
