package repo

import (
	"context"
	"github.com/spongeling/admin-api/internal/dao"
	"gorm.io/gorm"
	"strings"
)

func (r *Repo) GetCategoryExample(ctx context.Context, categoryID uint64) (*dao.CategoryExample, error) {
	phrase := &dao.CategoryExample{}
	err := r.db(ctx).
		Where("category_id = ?", categoryID).
		Preload("Pattern").
		Preload("FullPattern").
		First(phrase).Error
	return phrase, wrap(err)
}

func (r *Repo) GetCategoryPatternString(ctx context.Context, categoryID uint64) (string, error) {
	example := &dao.CategoryExample{}
	err := r.db(ctx).
		Model(&dao.CategoryExample{}).
		Where("category_id = ?", categoryID).
		Preload("Pattern").
		First(example).Error
	if err != nil {
		return "", wrap(err)
	}

	var tokens []string
	lastBlank := false
	for _, tokenId := range example.Pattern.TokenIds {
		token := &dao.Token{}
		err := r.GetWithAssociations(ctx, token, uint64(tokenId))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				token.Id = dao.NullTokenId
			} else {
				return "", wrap(err)
			}
		}

		// don't add adjacent blanks
		if token.IsBlank() {
			if lastBlank {
				continue
			}
			lastBlank = true
		} else {
			lastBlank = false
		}

		tokens = append(tokens, token.String())
	}

	return strings.Join(tokens, "--"), nil
}

func (r *Repo) AddCategoryExample(ctx context.Context, categoryId uint64, m *dao.CategoryExample) error {
	m.CategoryId = categoryId
	err := r.db(ctx).Create(m).Error
	return wrap(err)
}

func (r *Repo) UpdateCategoryExample(ctx context.Context, categoryId uint64, m *dao.CategoryExample) error {
	m.CategoryId = categoryId

	err := r.db(ctx).
		Model(&dao.CategoryExample{}).
		Where("category_id = ?", categoryId).
		Updates(m).Error

	return wrap(err)
}
