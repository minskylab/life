package main

import (
	"flag"
	"fmt"

	"github.com/minskylab/life"
)

var source, outputDir string
var directives, autoImport bool

func main() {
	flag.StringVar(&source, "s", "genesis/**/*.graphql", "Define your source as a path, or as a glob")
	flag.BoolVar(&directives, "d", true, "With built-in directives")
	flag.BoolVar(&autoImport, "a", true, "Active autoimport after generation")
	flag.StringVar(&outputDir, "o", "ent/schema", "Set your output directory for your ent structures")

	flag.Parse()

	if err := life.GenerateEntities(source, outputDir, life.GenerationOptions{
		EntDirectivesBuiltIn: directives,
		AutoImportProcessor:  autoImport,
	}); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

}
