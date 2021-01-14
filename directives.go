package life

import (
	"io"

	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
)

type customToken struct {
	rawCode string
}

func (t *customToken) render(f *jen.File, w io.Writer, s *jen.Statement) error {
	_, err := w.Write([]byte(t.rawCode))
	return errors.WithStack(err)
}

func (t *customToken) isNull(f *jen.File) bool {
	return false
}

func applyFieldDirective(directive *ast.Directive, s *jen.Statement) {
	switch directive.Name {
	case "unique":
		s.Dot("Unique").Call()
	case "immutable":
		s.Dot("Immutable").Call()
	case "nillable":
		s.Dot("Nillable").Call()
	case "storageKey":
		keyArg := directive.Arguments.ForName("key")
		if keyArg == nil || keyArg.Value == nil {
			return
		}

		s.Dot("StorageKey").Call(
			jen.Lit(keyArg.Value.String()),
		)
	case "sensitive":
		s.Dot("Sensitive").Call()
	case "default":
		valueArg := directive.Arguments.ForName("value")

		if valueArg != nil {
			s.Dot("Default").Call(
				jen.Op(valueArg.Value.Raw),
			)
		}

	case "updateDefault":
		valueArg := directive.Arguments.ForName("value")

		if valueArg != nil {
			s.Dot("UpdateDefault").Call(
				jen.Op(valueArg.Value.Raw),
			)
		}
	default:
		return
	}
}

func applyEdgeDirective(directive *ast.Directive, s *jen.Statement) {
	switch directive.Name {
	case "from":
		ref := directive.Arguments.ForName("ref")
		if ref != nil {
			refValue := ref.Value.Raw
			s.Dot("Ref").Call(jen.Lit(refValue))
		}
	default:
		return
	}
}
