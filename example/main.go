package main

import (
	"github.com/minskylab/life"
	"github.com/minskylab/life/effects/automutations"
)

func main() {
	if err := life.GenerateEntities("types/*.graphql", "ent/schema", life.GenerationOptions{
		EntDirectivesBuiltIn: true,
		// AutoImportProcessor:  true,
		Effects: []*life.EmergentEffect{
			automutations.NewAutoMutation("./mutations"),
		},
	}); err != nil {
		panic(err)
	}
}
