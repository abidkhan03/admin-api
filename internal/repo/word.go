package repo

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

func (r *Repo) GetWordId(ctx context.Context, word string) (uint64, error) {
	res := &dao.Word{Word: word}
	err := r.db(ctx).Where("word = ?", word).FirstOrCreate(&res).Error
	return res.ID, wrap(err)
}

func (r *Repo) GetWordClasses(ctx context.Context, word string) ([]dao.Class, error) {
	var res dao.Word
	err := r.db(ctx).
		Preload("Classes").
		Where("word = ?", word).
		First(&res).Error
	return res.Classes, wrap(err)
}
