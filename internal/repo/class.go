package repo

import (
	"context"
	"gorm.io/gorm/clause"

	"github.com/spongeling/admin-api/internal/dao"
)

func (r *Repo) GetAllClasses(ctx context.Context) ([]*dao.Class, error) {
	var res []*dao.Class
	err := r.db(ctx).
		Preload(clause.Associations).
		Find(&res).Error
	return res, wrap(err)
}

func (r *Repo) GetClassByName(ctx context.Context, name string) (*dao.Class, error) {
	var res *dao.Class
	err := r.db(ctx).
		Where("name = ?", name).
		Preload("Words").
		Find(&res).Error
	return res, wrap(err)
}

func (r *Repo) AddWordsToClass(ctx context.Context, classId uint64, words []string) error {
	var wc []dao.WordClass
	for _, word := range words {
		wordID, err := r.GetWordId(ctx, word)
		if err != nil {
			return err
		}
		wc = append(wc, dao.WordClass{
			WordId:  wordID,
			ClassId: classId,
		})
	}

	return wrap(r.db(ctx).Create(wc).Error)
}

func (r *Repo) UpdateWordsInClass(ctx context.Context, classId uint64, words []string) error {
	err := r.db(ctx).
		Where("class_id = ?", classId).
		Delete(&dao.WordClass{}).Error
	if err != nil {
		return err
	}

	err = r.AddWordsToClass(ctx, classId, words)

	return wrap(err)
}
