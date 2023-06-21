package response

import (
	"github.com/spongeling/admin-api/internal/service/model"
)

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type CategoryInfo struct {
	Pattern       *string    `json:"pattern"`
	SubCategories []Category `json:"sub_categories"`
}

type CategoryDetail struct {
	Id             uint64         `json:"id"`
	Phrase         string         `json:"phrase"`
	DefaultPattern *PhrasePattern `json:"default_pattern"`
	Pattern        []model.Token  `json:"pattern"`
	Examples       []string       `json:"examples"`
	Rule           string         `json:"rule"`
	Tip            string         `json:"tip"`
}
