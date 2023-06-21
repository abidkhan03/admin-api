package main

import (
	"flag"
	"fmt"
	"github.com/spongeling/admin-api/internal/errors"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HttpPort  int    `yaml:"httpPort"`
	GptApiKey string `yaml:"gptApiKey"`
}

func readConfig() (Config, error) {
	var c Config

	flag.IntVar(&c.HttpPort, "port", 6543, "Port for http server")
	flag.StringVar(&c.GptApiKey, "gptApiKey", "", "Migrations source")
	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Unknown arguments:", strings.Join(remainingArgs, " "))
		os.Exit(-1)
	}

	if x, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 64); err != nil {
		c.HttpPort = int(x)
	}
	if x := os.Getenv("OPENAI_API_KEY"); x != "" {
		c.GptApiKey = x
	}

	if c.HttpPort == 0 {
		return Config{}, errors.New(errors.InvalidArgument, "empty server port. -port flag required")
	}

	if c.GptApiKey == "" {
		return Config{}, errors.New(errors.InvalidArgument, "empty gpt api key. -gptApiKey flag required")
	}

	return c, nil
}
