package life

// GenerateEntities generates your entgo entities from a graphql source.
func GenerateEntities(source string, folderOut string) error {
	return generate(source, folderOut)
}

// MustGenerateEntities generates your entgo entities from a graphql source.
// Panic if occur an error.
func MustGenerateEntities(source string, folderOut string) {
	if err := generate(source, folderOut); err != nil {
		panic(err)
	}
}
