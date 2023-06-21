package analyzer

import (
	"bufio"
	"context"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/spongeling/admin-api/internal/gpt"
	"os"
	"strings"
)

type Config struct {
	Text string `yaml:"text"`
	File string `yaml:"file"`
}

func Run(cfg Config) error {
	apiKey := os.Getenv("OPENAI_API_KEY")

	c := gpt.NewClient(context.Background(), apiKey)

	if cfg.File != "" {
		file, err := os.OpenFile(cfg.File, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)

		bar := pb.StartNew(-1)
		defer bar.Finish()

		for scanner.Scan() {
			err = analyzeText(c, scanner.Text())
			if err != nil {
				return err
			}
			bar.Add(1)
		}
	} else {
		err := analyzeText(c, cfg.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func analyzeText(c *gpt.Client, text string) error {
	resp, err := c.Analyze(text)
	if err != nil {
		return err
	}

	for _, sentence := range resp.Sentences {
		for _, phrase := range sentence.Phrases {
			fmt.Println(">", text[phrase.Begin:phrase.End])
			fmt.Println()
			fmt.Println("word\tlemma\ttag")
			fmt.Println(strings.Repeat("-", 20))
			for _, word := range phrase.Words {
				fmt.Println(word.Word, "\t", word.Lemma, "\t", word.Tag)
			}
			fmt.Println()
		}
	}

	return nil
}
