package repo

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

func (r *Repo) GetAllUsers(ctx context.Context) ([]*dao.User, error) {
	var res []*dao.User
	err := r.db(ctx).Find(&res).Error
	return res, wrap(err)
}
