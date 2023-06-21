package service

import (
	"github.com/spongeling/admin-api/internal/freeling"
	"github.com/spongeling/admin-api/internal/gpt"
	"github.com/spongeling/admin-api/internal/repo"
)

type Service struct {
	repo     *repo.Repo
	gpt      *gpt.Client
	freeling *freeling.Client
}

func New(r *repo.Repo, g *gpt.Client, f *freeling.Client) *Service {
	return &Service{r, g, f}
}
