package main

import (
	"github.com/minskylab/life"
)

func main() {
	if err := life.GenerateEntities("types/*.graphql", "ent/schema", life.GenerationOptions{
		EntDirectivesBuiltIn: true,
		// AutoImportProcessor:  true,
	}); err != nil {
		panic(err)
	}
}
