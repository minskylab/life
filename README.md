# Life Project

Life project is a initiative to create a toolkit to automatize the generation of
any system based in entities and their relations.

## Current state

Actually Life can be used to generate [entgo](https://entgo.io/) schemas from GraphQL schema.

**Example**

```go
package main

import "github.com/minskylab/life"

func main() {
    // read all your `types/*.graphql` and generate entities into `ent/schema`
	if err := life.GenerateEntities("types/*.graphql", "ent/schema"); err != nil {
		panic(err)
	}
}
```
