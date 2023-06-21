package main

import (
	"flag"
	"fmt"
	"github.com/spongeling/admin-api/freeling-analyzer"
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/shared"
	"log"
	"os"
	"strings"
)

func main() {
	// config
	err := shared.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = analyzer.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func readConfig() (analyzer.Config, error) {
	var c analyzer.Config

	flag.StringVar(&c.Text, "text", "", "text to analyze")
	flag.StringVar(&c.File, "file", "", "file to analyze")
	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Unknown arguments:", strings.Join(remainingArgs, " "))
		os.Exit(-1)
	}

	if c.Text == "" && c.File == "" {
		return analyzer.Config{}, errors.New(errors.InvalidArgument, "`text` or `file` required")
	}

	return c, nil
}
