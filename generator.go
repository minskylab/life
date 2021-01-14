package life

import "github.com/pkg/errors"

// GenerationOptions provides options to configure the entities generation
type GenerationOptions struct {
	WithGoEntDirectives bool
}

// GenerateEntities generates your entgo entities from a graphql source.
func GenerateEntities(source string, folderOut string, opts GenerationOptions) error {
	if err := generate(source, folderOut, opts.WithGoEntDirectives); err != nil {
		return errors.WithStack(err)
	}

	fixImports(folderOut)

	return nil
}

// MustGenerateEntities generates your entgo entities from a graphql source.
// Panic if occur an error.
func MustGenerateEntities(source string, folderOut string, opts GenerationOptions) {
	if err := generate(source, folderOut, opts.WithGoEntDirectives); err != nil {
		panic(err)
	}
}
