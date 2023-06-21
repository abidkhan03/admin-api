package repo

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

func (r *Repo) GetAllTopLevelCategories(ctx context.Context) ([]*dao.Category, error) {
	var res []*dao.Category
	err := r.db(ctx).
		Where("parent_id IS NULL").
		Order("id ASC").
		Find(&res).Error
	return res, wrap(err)
}

func (r *Repo) GetSubCategories(ctx context.Context, categoryID uint64) ([]*dao.Category, error) {
	var res []*dao.Category
	err := r.db(ctx).
		Where("parent_id = ?", categoryID).
		Order("id ASC").
		Find(&res).Error
	return res, wrap(err)
}

func (r *Repo) GetTopCategoryId(ctx context.Context, name string) (uint64, error) {
	cat := &dao.Category{
		Name: name,
	}
	err := r.db(ctx).Where(cat).FirstOrCreate(cat).Error

	return cat.Id, wrap(err)
}

func (r *Repo) GetSubCategoryId(ctx context.Context, parentId uint64, name string) (uint64, error) {
	cat := &dao.Category{
		Name:     name,
		ParentId: &parentId,
	}
	err := r.db(ctx).Where(cat).FirstOrCreate(cat).Error

	return cat.Id, wrap(err)
}
