package life

import "github.com/pkg/errors"

// GenerationOptions provides options to configure the entities generation
type GenerationOptions struct {
	EntDirectivesBuiltIn bool
	AutoImportProcessor  bool
	Effects              []*EmergentEffect
}

// GenerateEntities generates your entgo entities from a graphql source.
func GenerateEntities(source string, folderOut string, opts GenerationOptions) error {
	if err := generate(source, folderOut, opts.EntDirectivesBuiltIn); err != nil {
		return errors.WithStack(err)
	}

	if opts.AutoImportProcessor {
		fixImports(folderOut)
	}

	return nil
}

// MustGenerateEntities generates your entgo entities from a graphql source.
// Panic if occur an error.
func MustGenerateEntities(source string, folderOut string, opts GenerationOptions) {
	if err := GenerateEntities(source, folderOut, opts); err != nil {
		panic(err)
	}
}
