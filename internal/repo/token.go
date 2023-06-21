package repo

import (
	"context"
	"fmt"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/service/model"
	"strings"
)

func (r *Repo) GetDbToken(ctx context.Context, token model.Token) (*dao.Token, error) {
	var t dao.Token
	if token.Pos != nil {
		p, err := dao.POSFromString(*token.Pos)
		if err != nil {
			return nil, err
		}

		posId, err := r.GetPosId(ctx, p)
		if err != nil {
			return nil, err
		}

		t.PosId = &posId
	}
	if token.Word != nil {
		wordId, err := r.GetWordId(ctx, *token.Word)
		if err != nil {
			return nil, err
		}

		t.WordId = &wordId
	}
	if token.Class != nil {
		class, err := r.GetClassByName(ctx, *token.Class)
		if err != nil {
			return nil, err
		}

		t.ClassId = &class.Id
	}

	var condition []string
	if t.PosId != nil {
		condition = append(condition, fmt.Sprintf("pos_id=%d", *t.PosId))
	} else {
		condition = append(condition, "pos_id IS NULL")
	}
	if t.WordId != nil {
		condition = append(condition, fmt.Sprintf("word_id=%d", *t.WordId))
	} else {
		condition = append(condition, "word_id IS NULL")
	}
	if t.ClassId != nil {
		condition = append(condition, fmt.Sprintf("class_id=%d", *t.ClassId))
	} else {
		condition = append(condition, "class_id IS NULL")
	}

	err := r.db(ctx).Where(strings.Join(condition, " AND ")).FirstOrCreate(&t).Error
	if err != nil {
		return nil, err
	}

	return &t, nil
}
