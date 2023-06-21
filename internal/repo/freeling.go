package repo

import (
	"context"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/freeling"
)

func (r *Repo) SaveFreeLingResponse(ctx context.Context, response *freeling.Response) error {
	for _, sentence := range response.Sentences {

		var tokenIds []uint64
		for _, word := range sentence.Tokens {
			err := r.AddFreeLingParsedWord(ctx, &word)
			if err != nil {
				return wrap(err)
			}
			pos, err := dao.POSFromString(word.Tag)
			if err != nil {
				return wrap(err)
			}
			posId, err := r.GetPosId(ctx, pos)
			if err != nil {
				return wrap(err)
			}
			tokenIds = append(tokenIds, posId)
		}

		_, err := r.GetPatternId(ctx, tokenIds)
		if err != nil {
			return wrap(err)
		}
	}

	return nil
}

func (r *Repo) AddFreeLingParsedWord(ctx context.Context, word *freeling.Token) error {
	wordId, err := r.GetWordId(ctx, word.Form)
	if err != nil {
		return err
	}
	lemmaId, err := r.GetWordId(ctx, word.Lemma)
	if err != nil {
		return err
	}
	pos, err := dao.POSFromString(word.Tag)
	if err != nil {
		return err
	}
	posId, err := r.GetPosId(ctx, pos)
	if err != nil {
		return err
	}

	return wrap(
		r.db(ctx).FirstOrCreate(
			&dao.WordPos{
				WordId:  wordId,
				LemmaId: lemmaId,
				PosId:   posId,
			},
		).Error,
	)
}
