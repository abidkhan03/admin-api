package service

import (
	"context"
	"fmt"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/response"
	"net/http"
)

func (s *Service) GetAllTopLevelCategories(ctx context.Context) ([]*dao.Category, error) {
	return s.repo.GetAllTopLevelCategories(ctx)
}

func (s *Service) GetSubCategories(ctx context.Context, categoryId uint64) ([]*dao.Category, error) {
	return s.repo.GetSubCategories(ctx, categoryId)
}

func (s *Service) GetCategoryInfo(ctx context.Context, categoryId uint64) (*response.CategoryInfo, error) {
	subCategories, err := s.repo.GetSubCategories(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	res := &response.CategoryInfo{}
	for _, subCategory := range subCategories {
		res.SubCategories = append(res.SubCategories, response.Category{
			Id:   subCategory.Id,
			Name: subCategory.Name,
		})
	}

	pattern, err := s.repo.GetCategoryPatternString(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	res.Pattern = &pattern

	return res, nil
}

func (s *Service) GetCategory(ctx context.Context, categoryId uint64) (*response.CategoryDetail, error) {
	example, err := s.repo.GetCategoryExample(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	// build response
	res := response.CategoryDetail{
		Id:     example.CategoryId,
		Phrase: example.Phrase,
	}

	if example.Rule != nil {
		res.Rule = *example.Rule
	}

	if example.Tip != nil {
		res.Tip = *example.Tip
	}

	res.DefaultPattern, err = s.GetPhrasePattern(ctx, example.Phrase)
	if err != nil {
		return nil, err
	}

	res.Pattern, err = s.patternResponse(ctx, example.FullPattern)
	if err != nil {
		return nil, err
	}

	pattern, err := s.patternResponse(ctx, example.Pattern)
	if err != nil {
		return nil, err
	}

	res.Examples, err = s.GetPatternExamples(ctx, pattern)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (s *Service) AddCategory(ctx context.Context, req request.CategoryDetails) (uint64, error) {
	var categoryId uint64
	err := s.repo.RunTx(ctx, func(ctx context.Context) error {
		var err error
		categoryId, err = s.repo.GetTopCategoryId(ctx, req.Category.Name)
		if err != nil {
			return err
		}
		for req.Category.SubCategory != nil {
			req.Category = *req.Category.SubCategory
			categoryId, err = s.repo.GetSubCategoryId(ctx, categoryId, req.Category.Name)
			if err != nil {
				return err
			}
		}

		// add new example to the database
		example := &dao.CategoryExample{
			CategoryId: categoryId,
			Phrase:     req.Phrase,
			Rule:       req.Rule,
			Tip:        req.Tip,
		}

		allTokenIds, tokenIds, err := s.getTokenIds(ctx, req.Pattern)
		if err != nil {
			return err
		}

		example.PatternId, err = s.repo.GetPatternId(ctx, tokenIds)
		if err != nil {
			return err
		}

		assoc, err := s.repo.GetPatternAssociatedExample(ctx, example.PatternId)
		if err == nil {
			return errors.New(http.StatusConflict, fmt.Sprintf("The given pattern is associated to the category `%s`", assoc.Category.Name))
		}

		example.FullPatternId, err = s.repo.GetPatternId(ctx, allTokenIds)
		if err != nil {
			return err
		}

		err = s.repo.AddCategoryExample(ctx, categoryId, example)

		return err
	})

	return categoryId, err
}

func (s *Service) UpdateCategory(ctx context.Context, categoryId uint64, req request.UpdatePatternPhrase) error {
	return s.repo.RunTx(ctx, func(ctx context.Context) error {
		example, err := s.repo.GetCategoryExample(ctx, categoryId)
		if err != nil {
			return err
		}

		allTokenIds, tokenIds, err := s.getTokenIds(ctx, req.Pattern)
		if err != nil {
			return err
		}

		originalPatternId := example.PatternId
		example.PatternId, err = s.repo.ChangePattern(ctx, originalPatternId, tokenIds)
		if err != nil {
			return err
		}

		assoc, err := s.repo.GetPatternAssociatedExample(ctx, example.PatternId)
		if err == nil && assoc.CategoryId != categoryId {
			return errors.New(http.StatusConflict, fmt.Sprintf("The given pattern is associated to the category `%s`", assoc.Category.Name))
		}

		originalFullPatternId := example.FullPatternId
		if example.Phrase != req.Phrase {
			if example.FullPatternId == example.PatternId {
				example.FullPatternId, err = s.repo.GetPatternId(ctx, allTokenIds)
				if err != nil {
					return err
				}
			} else {
				example.FullPatternId, err = s.repo.ChangePattern(ctx, example.FullPatternId, allTokenIds)
				if err != nil {
					return err
				}
			}

			example.Phrase = req.Phrase
		}

		example.Rule = req.Rule
		example.Tip = req.Tip

		err = s.repo.UpdateCategoryExample(ctx, categoryId, example)
		if err != nil {
			return err
		}

		if example.PatternId != originalPatternId {
			_ = s.repo.DeleteByID(ctx, &dao.Pattern{}, originalPatternId)
		}
		if example.FullPatternId != originalFullPatternId {
			_ = s.repo.DeleteByID(ctx, &dao.Pattern{}, originalFullPatternId)
		}

		return nil
	})
}

func (s *Service) DeleteCategory(ctx context.Context, categoryId uint64) error {
	return s.repo.DeleteByID(ctx, &dao.Category{}, categoryId)
}
