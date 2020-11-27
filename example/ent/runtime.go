// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/minskylab/life/example/ent/autor"
	"github.com/minskylab/life/example/ent/schema"
	"github.com/minskylab/life/example/ent/todo"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	autorFields := schema.Autor{}.Fields()
	_ = autorFields
	// autorDescID is the schema descriptor for id field.
	autorDescID := autorFields[0].Descriptor()
	// autor.IDValidator is a validator for the "id" field. It is called by the builders before save.
	autor.IDValidator = autorDescID.Validators[0].(func(string) error)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.IDValidator is a validator for the "id" field. It is called by the builders before save.
	todo.IDValidator = todoDescID.Validators[0].(func(string) error)
}