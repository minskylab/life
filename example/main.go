package main

import (
	"github.com/minskylab/life"
	"github.com/minskylab/life/effects/automutations"
)

func main() {
	autoMut, err := automutations.NewAutoMutationEffect("./mutations")
	if err != nil {
		panic(err)
	}

	if err := life.GenerateEntities("types/*.graphql", "ent/schema", life.GenerationOptions{
		EntDirectivesBuiltIn: true,
		// AutoImportProcessor:  true,
		Effects: []life.EmergentEffect{autoMut},
	}); err != nil {
		panic(err)
	}
}
