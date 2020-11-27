package life

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
)

func generateTypes(folder string, definitions []*ast.Definition, enums map[string]*ast.Definition) error {
	_ = os.MkdirAll(folder, os.ModePerm)

	for _, def := range definitions {
		entitityFile := generateType(def, enums)

		filepath := path.Join(folder, strings.ToLower(def.Name)+".go")
		f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
		if err != nil {
			return errors.WithStack(err)
		}

		defer f.Close()

		if err = entitityFile.Render(f); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
