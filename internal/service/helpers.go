package service

import (
	"context"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/service/model"
	"gorm.io/gorm"
)

func (s *Service) patternResponse(ctx context.Context, p *dao.Pattern) ([]model.Token, error) {
	var res []model.Token
	for i, tokenId := range p.TokenIds {
		token := &dao.Token{}
		err := s.repo.GetWithAssociations(ctx, token, uint64(tokenId))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				token.Id = dao.NullTokenId
			} else {
				return nil, err
			}
		}

		t := model.Token{
			SeqId: uint64(i + 1),
		}
		if token.Pos != nil {
			p := token.Pos.ToPOS().String()
			t.Pos = &p
		}
		if token.Word != nil {
			t.Word = &token.Word.Word
		}
		if token.Class != nil {
			t.Class = &token.Class.Name
		}

		res = append(res, t)
	}

	return res, nil
}

func (s *Service) getTokenIds(ctx context.Context, tokens []model.Token) ([]uint64, []uint64, error) {
	var tokenIds []uint64
	var allTokenIds []uint64

	lastBlank := false
	for _, token := range tokens {
		t, err := s.repo.GetDbToken(ctx, token)
		if err != nil {
			return nil, nil, err
		}

		allTokenIds = append(allTokenIds, t.Id)

		// don't add adjacent blanks
		if token.Type() == model.Blank {
			if lastBlank {
				continue
			}
			lastBlank = true
		} else {
			lastBlank = false
		}

		tokenIds = append(tokenIds, t.Id)
	}

	return allTokenIds, tokenIds, nil
}

func (s *Service) getTokensDetail(ctx context.Context, tokens []model.Token) error {
	for i := range tokens {
		err := s.getTokenDetail(ctx, &tokens[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) getTokenDetail(ctx context.Context, token *model.Token) error {
	switch token.Type() {
	case model.Class:
		class, err := s.repo.GetClassByName(ctx, *token.Class)
		if err != nil {
			return err
		}

		token.ClassId = &class.Id

		for _, w := range class.Words {
			token.ClassWords = append(token.ClassWords, w.Word)
		}
	case model.ExactWord:
		wordId, err := s.repo.GetWordId(ctx, *token.Word)
		if err != nil {
			return err
		}

		token.WordId = &wordId
	case model.POS:
		pos, err := dao.POSFromString(*token.Pos)
		if err != nil {
			return err
		}

		posId, err := s.repo.GetPosId(ctx, pos)
		if err != nil {
			return err
		}

		token.PosId = &posId
	}

	return nil
}
