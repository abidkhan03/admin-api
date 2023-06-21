package request

import (
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/internal/service/model"
)

type Category struct {
	Name        string    `json:"category"`
	SubCategory *Category `json:"sub_category"`
}

type CategoryDetails struct {
	Category Category      `json:"category"`
	Phrase   string        `json:"phrase"`
	Pattern  []model.Token `json:"pattern"`
	Examples []string      `json:"examples"`
	Rule     *string       `json:"rule"`
	Tip      *string       `json:"tip"`
}

func (pp *CategoryDetails) Validate() error {
	if pp.Category.Name == "" {
		return errors.BadRequest("category should not be empty")
	}
	if pp.Phrase == "" {
		return errors.BadRequest("phrase should not be empty")
	}
	if len(pp.Pattern) == 0 {
		return errors.BadRequest("pattern should not be empty")
	}
	if len(pp.Examples) == 0 {
		return errors.BadRequest("no example was provided")
	}
	return nil
}

type UpdatePatternPhrase struct {
	Phrase   string        `json:"phrase"`
	Pattern  []model.Token `json:"pattern"`
	Examples []string      `json:"examples"`
	Rule     *string       `json:"rule"`
	Tip      *string       `json:"tip"`
}

func (pp *UpdatePatternPhrase) Validate() error {
	if pp.Phrase == "" {
		return errors.BadRequest("phrase should not be empty")
	}
	if len(pp.Pattern) == 0 {
		return errors.BadRequest("pattern should not be empty")
	}
	return nil
}

type Pattern struct {
	Tokens []model.Token `json:"tokens"`
}

func (t *Pattern) Validate() error {
	if len(t.Tokens) == 0 {
		return errors.BadRequest("there should be at least one token")
	}
	return nil
}

type Phrase struct {
	Phrase string `json:"phrase"`
}

func (p *Phrase) Validate() error {
	if p.Phrase == "" {
		return errors.BadRequest("phrase should not be empty")
	}
	return nil
}
