package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/minskylab/life"
)

func main() {
	flag.Parse()

	source := "genesis/**/*.graphql"
	outputDir := "ent/schema"

	ws := flag.Bool("wd", true, "With builtin directives")

	args := os.Args

	fmt.Println(args)

	if len(args) == 1 {
		source = args[0]
	} else if len(args) > 1 {
		source = args[0]
		outputDir = args[1]
	}

	if err := life.GenerateEntities(source, outputDir, life.GenerationOptions{
		WithGoEntDirectives: *ws,
	}); err != nil {
		panic(err)
	}

}
