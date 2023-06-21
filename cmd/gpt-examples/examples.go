package main

import (
	"context"
	"fmt"
	"github.com/spongeling/admin-api/internal/gpt"
	"github.com/spongeling/admin-api/internal/service/model"
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

	// args
	if len(os.Args) < 2 {
		log.Fatalf(
			`no pattern provided...
Usage Examples:
	%[1]s "--DA0MS0-NCMS000-AQ0MS00--"
	%[1]s "--BLANK-NCMS000-AQ0MS00--"
	%[1]s "--token_el-NCMS000-AQ0MS00--"
`, os.Args[0],
		)
	}

	tokens := strings.Split(strings.Trim(os.Args[1], "-"), "-")

	var pattern []model.Token
	for _, token := range tokens {
		if token == "BLANK" {
			pattern = append(pattern, model.Token{})
		} else if t, ok := strings.CutPrefix(token, "token_"); ok {
			pattern = append(pattern, model.Token{Word: &t})
		} else {
			pattern = append(pattern, model.Token{Pos: &token})
		}
	}

	// examples
	examples, err := gpt.
		NewClient(context.Background(), os.Getenv("OPENAI_API_KEY")).
		GetPatternExamples(pattern)
	if err != nil {
		log.Fatalln(err)
	}

	for _, example := range examples {
		fmt.Println(example)
	}
}
