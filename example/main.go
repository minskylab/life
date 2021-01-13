package main

import "github.com/minskylab/life"

func main() {
	if err := life.GenerateEntities("types/*.graphql", "ent/schema", life.GenerationOptions{
		WithGoEntDirectives: true,
	}); err != nil {
		panic(err)
	}
}
