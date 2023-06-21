package gpt

import (
	"context"
	"fmt"
	"github.com/homelight/json"
	"github.com/sashabaranov/go-openai"
	"github.com/spongeling/admin-api/internal/service/model"
	"log"
	"regexp"
	"strings"
)

type Client struct {
	c   *openai.Client
	ctx context.Context
}

func NewClient(ctx context.Context, apiKey string) *Client {
	return &Client{
		c:   openai.NewClient(apiKey),
		ctx: ctx,
	}
}

// Analyze the given text to divide it into sentences, phrases and word,
// and get POS tags for the words according to the context.
func (c *Client) Analyze(text string) (*Response, error) {
	prompt := fmt.Sprintf(`Divide the Spanish text given in input tag into sentences. Divide sentences into multiple phrases based on conjunction. Then phrases into words. Tag the words with FreeLing POS tags as per context.
Respond in the following compact JSON format:
{"sentences":[{"seq_id":0,"begin":0,"end":0,"phrases":[{"seq_id":0,"begin":0,"end":0,"phrase":"","words":[{"seq_id":0,"word":"","lemma":"","tag":""}]}]}]} \

<input>%s</input>
`, text)
	res, err := c.request(prompt)
	if err != nil {
		return nil, err
	}

	var result Response
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse response: %s\n%s\n", err, res)
	}

	return &result, err
}

// GetPatternExamples return Spanish phrases that match the given POS tags pattern
func (c *Client) GetPatternExamples(pattern []model.Token) ([]string, error) {
	var prompt string
	lastBlank := false
	i := 1
	nonBlankCount := 0
	for _, t := range pattern {
		if t.Type() == model.Blank {
			if lastBlank {
				continue
			}
			lastBlank = true
		} else {
			lastBlank = false
		}

		nonBlankCount++

		switch t.Type() {
		case model.POS:
			prompt += fmt.Sprintf("\n%d. a token matching the POS tag `%s`\\", i, *t.Pos)
		case model.ExactWord:
			prompt += fmt.Sprintf("\n%d. `%s` or it's lemma\\", i, *t.Word)
		case model.Blank:
			prompt += fmt.Sprintf("\n%d. any token (1 or more) which fit into the phrase\\", i)
		case model.Class:
			prompt += fmt.Sprintf("\n%d. a token (or its lemma) from the list: %s\\", i, t.ClassWords)
		}

		i++
	}

	prompt = fmt.Sprintf(`You'll need to know FreeLing Spanish tags. Generate three to five valid Spanish phrases meeting the input criteria.\

Output Example: ["phrase1","phrase2",...,"phraseN"]

Phrases should have exactly %d parts:\`, nonBlankCount) + prompt

	log.Println("Getting examples from GPT...")
	//log.Println(examplesRequest)
	log.Println(prompt)

	res, err := c.request(prompt)
	if err != nil {
		return nil, err
	}

	log.Println("GPT Response:")
	log.Println(res)

	var phrases []string
	for _, phrase := range examplesRegex.FindAllString(res, -1) {
		phrases = append(phrases, strings.Trim(phrase, `"`))
	}

	return phrases, nil
}

var examplesRegex = regexp.MustCompile(`"(.*?)"`)

// ListWordsInCategory will return words that belong to the given category
// e.g.
// color -> [rojo verde azul amarillo naranja morado rosa blanco negro gris marrón]
// country -> [España  México  Argentina  Colombia  Perú  Chile  Cuba  Ecuador]
func (c *Client) ListWordsInCategory(category string) ([]string, error) {
	prompt := fmt.Sprintf(`Generate Spanish words that represents "%s".\
Example Input: color
Example Output: ["rojo", "verde", "azul", ...]
`, category)

	res, err := c.request(prompt)
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse response: %s\n%s\n", err, res)
	}

	return result, err
}
