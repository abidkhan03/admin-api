package service

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

func (s *Service) GetAllUsers(ctx context.Context) ([]*dao.User, error) {
	return s.repo.GetAllUsers(ctx)
}
