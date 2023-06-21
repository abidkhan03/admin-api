package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/spongeling/admin-api/internal/freeling"
)

type Config struct {
	Text string `yaml:"text"`
	File string `yaml:"file"`
}

func Run(cfg Config) error {
	srv, err := freeling.NewServer(nil)
	if err != nil {
		return err
	}
	err = srv.Start()
	if err != nil {
		return err
	}
	defer srv.Stop()

	time.Sleep(5 * time.Second)

	c, err := freeling.NewClient()
	if err != nil {
		return err
	}
	defer c.Close()

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
		err = analyzeText(c, cfg.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func analyzeText(c *freeling.Client, text string) error {
	resp, err := c.Analyze(text)
	if err != nil {
		return err
	}

	for _, sentence := range resp.Sentences {
		//begin, _ := strconv.Atoi(sentence.Tokens[0].Begin)
		//end, _ := strconv.Atoi(sentence.Tokens[len(sentence.Tokens)-1].End)
		//fmt.Println(">", text[begin:end])
		// Unicode issue: unicode uses 2 characters
		fmt.Println()
		fmt.Println("word\tlemma\ttag")
		fmt.Println(strings.Repeat("-", 20))
		for _, token := range sentence.Tokens {
			fmt.Println(token.Form, "\t", token.Lemma, "\t", token.Tag)
		}
		fmt.Println()
	}

	return nil
}
