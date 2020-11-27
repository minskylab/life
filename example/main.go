package main

import "github.com/minskylab/life"

func main() {
	if err := life.GenerateEntities("types/*.graphql", "ent/schema"); err != nil {
		panic(err)
	}
}
