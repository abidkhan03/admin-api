package service

import (
	"context"
	"fmt"
	"github.com/spongeling/admin-api/internal/freeling"
	"github.com/spongeling/admin-api/internal/response"
	"github.com/spongeling/admin-api/internal/service/model"
	"log"
)

func (s *Service) GetPhrasePattern(ctx context.Context, phrase string) (*response.PhrasePattern, error) {
	//phrase = strings.ToLower(shared.CleanUpText(phrase))

	result, err := s.freeling.Analyze(phrase)
	if err != nil {
		return nil, err
	}

	go func() {
		err = s.repo.SaveFreeLingResponse(ctx, result)
		if err != nil {
			log.Println(err)
		}
	}()

	if len(result.Sentences) < 1 {
		return nil, fmt.Errorf("failed to extract sentence for the phrase: %s", phrase)
	}
	flSentence := result.Sentences[0]
	if len(flSentence.Tokens) < 1 {
		return nil, fmt.Errorf("failed to extract phrase for the phrase: %s", phrase)
	}

	resp := &response.PhrasePattern{}
	for _, flWord := range flSentence.Tokens {
		p := response.TokenInfo{
			SeqId: flWord.ID,
			Word:  flWord.Form,
			Pos:   flWord.Tag,
		}

		c, _ := s.repo.GetWordClasses(ctx, flWord.Form)
		for _, wc := range c {
			p.Classes = append(p.Classes, wc.Name)
		}

		resp.Tokens = append(resp.Tokens, p)
	}

	return resp, err
}

func (s *Service) GetPatternExamples(ctx context.Context, pattern []model.Token) ([]string, error) {
	err := s.getTokensDetail(ctx, pattern)
	if err != nil {
		return nil, err
	}

	examples, err := s.gpt.GetPatternExamples(pattern)
	if err != nil {
		return nil, err
	}

	// filter valid examples
	//var resp []string
	//for _, example := range examples {
	//	if s.IsExampleValid(example, pattern) {
	//		resp = append(resp, example)
	//	}
	//}

	return examples, nil
}

// IsExampleValid checks if the example is valid according to the input pattern
func (s *Service) IsExampleValid(example string, inputPattern []model.Token) bool {
	resp, err := s.freeling.Analyze(example)
	if err != nil {
		log.Panicln(err) // HANDLE ERROR
		return false
	}

	var examplePattern []freeling.Token
	for _, sentence := range resp.Sentences {
		examplePattern = append(examplePattern, sentence.Tokens...)
	}

	dp := make([][]*bool, len(inputPattern))
	for i := range dp {
		dp[i] = make([]*bool, len(examplePattern))
	}

	return s.matchRecursive(dp, inputPattern, examplePattern, 0, 0)
}

// matchRecursive recursively checks if the example pattern matches with the input pattern
func (s *Service) matchRecursive(dp [][]*bool, inputPattern []model.Token, examplePattern []freeling.Token, ii, ei int) bool {
	if ii >= len(inputPattern) && ei >= len(examplePattern) {
		return true
	}
	if ii >= len(inputPattern) || ei >= len(examplePattern) {
		return false
	}

	if dp[ii][ei] != nil {
		return *dp[ii][ei]
	}

	ans := true
	if inputPattern[ii].Class != nil {
		// get words under this class from the database
		class, err := s.repo.GetClassByName(context.TODO(), *inputPattern[ii].Class)
		if err != nil {
			log.Panicln(err) // HANDLE ERROR
			return false
		}

		// check if the example word is one of the class words or their lemma
		ans = false
		for _, w := range class.Words {
			if w.Word == examplePattern[ei].Form || w.Word == examplePattern[ei].Lemma { // TODO: check lemma otherway around
				ans = s.matchRecursive(dp, inputPattern, examplePattern, ii+1, ei+1)
				break
			}
		}
		if !ans {
			var words []string
			for _, w := range class.Words {
				words = append(words, w.Word)
			}
			log.Printf("> `%s` or its lemma `%s` doesn't belong to the class `%s` %s\n",
				examplePattern[ei].Form, examplePattern[ei].Lemma,
				class.Name, words,
			)
		}
	} else if inputPattern[ii].Word != nil {
		// in case of word, match word/lemma and move to next item
		if *inputPattern[ii].Word != examplePattern[ei].Form && *inputPattern[ii].Word != examplePattern[ei].Lemma { // TODO: check lemma otherway around
			log.Printf("> `%s` or its lemma `%s` != `%s` or its lemma `%s`\n",
				examplePattern[ei].Form, examplePattern[ei].Lemma,
				*inputPattern[ii].Word, "", // inputPattern[ii].Lemma,
			)
			ans = false
		} else {
			ans = s.matchRecursive(dp, inputPattern, examplePattern, ii+1, ei+1)
		}
	} else if inputPattern[ii].Pos != nil {
		// for pos, we need to check if it matches, except 0, because it means it can be any option
		iTag := *inputPattern[ii].Pos
		eTag := examplePattern[ei].Tag

		if len(iTag) != len(eTag) {
			ans = false
		} else {
			for i := range iTag {
				if iTag[i] != '0' && eTag[i] != iTag[i] {
					ans = false
					log.Printf("> `%s` (%s) != `%s`\n",
						examplePattern[ei].Form, eTag, iTag,
					)
					break
				}
			}
		}
		if ans {
			ans = s.matchRecursive(dp, inputPattern, examplePattern, ii+1, ei+1)
		}
	} else {
		// either pick one work for blank, or try the next words too
		ans = s.matchRecursive(dp, inputPattern, examplePattern, ii+1, ei+1) || s.matchRecursive(dp, inputPattern, examplePattern, ii, ei+1)
	}

	dp[ii][ei] = &ans

	return ans
}
