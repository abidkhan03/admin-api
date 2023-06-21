package repo

import (
	"context"
	"github.com/spongeling/admin-api/internal/dao"
	"gorm.io/gorm"
)

func (r *Repo) GetPatternId(ctx context.Context, tokenIds []uint64) (uint64, error) {
	pattern := &dao.Pattern{}
	for _, tokenId := range tokenIds {
		pattern.TokenIds = append(pattern.TokenIds, int64(tokenId))
	}
	err := r.db(ctx).Where(pattern).FirstOrCreate(pattern).Error
	return pattern.Id, wrap(err)
}

func (r *Repo) ChangePattern(ctx context.Context, patternId uint64, tokenIds []uint64) (uint64, error) {
	pattern := &dao.Pattern{}

	lastBlank := false
	for _, tokenId := range tokenIds {
		// don't add adjacent blanks
		if tokenId == dao.NullTokenId {
			if lastBlank {
				continue
			}
			lastBlank = true
		} else {
			lastBlank = false
		}

		pattern.TokenIds = append(pattern.TokenIds, int64(tokenId))
	}

	err := r.db(ctx).Where(pattern).First(pattern).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { // not found
			err = r.UpdatePattern(ctx, patternId, tokenIds)
		}
		return patternId, wrap(err)
	} else { // found
		return pattern.Id, nil
	}
}

func (r *Repo) UpdatePattern(ctx context.Context, patternId uint64, tokenIds []uint64) error {
	pattern := &dao.Pattern{
		Id: patternId,
	}
	for _, tokenId := range tokenIds {
		pattern.TokenIds = append(pattern.TokenIds, int64(tokenId))
	}

	err := r.db(ctx).
		Model(&dao.Pattern{}).
		Where("id=?", patternId).
		Updates(pattern).Error

	return wrap(err)
}

func (r *Repo) GetPatternAssociatedExample(ctx context.Context, patternId uint64) (*dao.CategoryExample, error) {
	example := &dao.CategoryExample{}
	err := r.db(ctx).
		Where("pattern_id=?", patternId).
		Preload("Category").
		First(example).Error
	return example, wrap(err)
}
