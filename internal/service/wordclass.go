package service

import (
	"context"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/response"
)

func (s *Service) GetAllClasses(ctx context.Context) ([]response.Class, error) {
	// fetch the word classes
	wordClasses, err := s.repo.GetAllClasses(ctx)
	if err != nil {
		return nil, err
	}

	// build response
	var res []response.Class
	for _, wc := range wordClasses {
		r := response.Class{
			Id:          wc.Id,
			Name:        wc.Name,
			Description: wc.Description,
		}

		for _, word := range wc.Words {
			r.Words = append(r.Words, word.Word)
		}

		res = append(res, r)
	}

	return res, nil
}

func (s *Service) AddClass(ctx context.Context, req request.Class) (uint64, error) {
	var classId uint64
	err := s.repo.RunTx(ctx, func(ctx context.Context) error {
		// adding class
		class := &dao.Class{
			Name:        req.Name,
			Description: req.Description,
		}
		err := s.repo.Create(ctx, class)
		if err != nil {
			return err
		}

		classId = class.Id

		// associate words to the class
		return s.repo.AddWordsToClass(ctx, classId, req.Words)
	})

	return classId, err
}

func (s *Service) UpdateWordClass(ctx context.Context, classId uint64, req request.Class) error {
	return s.repo.RunTx(ctx, func(ctx context.Context) error {
		err := s.repo.Update(ctx, &dao.Class{
			Id:          classId,
			Name:        req.Name,
			Description: req.Description,
		})
		if err != nil {
			return err
		}

		return s.repo.UpdateWordsInClass(ctx, classId, req.Words)
	})
}

func (s *Service) DeleteClass(ctx context.Context, classId uint64) error {
	return s.repo.DeleteByID(ctx, &dao.Class{}, classId)
}
