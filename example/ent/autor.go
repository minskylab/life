// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/minskylab/life/example/ent/autor"
)

// Autor is the model entity for the Autor schema.
type Autor struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature string `json:"signature,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AutorQuery when eager-loading is set.
	Edges AutorEdges `json:"edges"`
}

// AutorEdges holds the relations/edges for other nodes in the graph.
type AutorEdges struct {
	// Todos holds the value of the todos edge.
	Todos []*Todo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodosOrErr returns the Todos value or an error if the edge
// was not loaded in eager-loading.
func (e AutorEdges) TodosOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todos, nil
	}
	return nil, &NotLoadedError{edge: "todos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Autor) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullString{}, // name
		&sql.NullString{}, // signature
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Autor fields.
func (a *Autor) assignValues(values ...interface{}) error {
	if m, n := len(values), len(autor.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		a.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		a.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field signature", values[1])
	} else if value.Valid {
		a.Signature = value.String
	}
	return nil
}

// QueryTodos queries the todos edge of the Autor.
func (a *Autor) QueryTodos() *TodoQuery {
	return (&AutorClient{config: a.config}).QueryTodos(a)
}

// Update returns a builder for updating this Autor.
// Note that, you need to call Autor.Unwrap() before calling this method, if this Autor
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Autor) Update() *AutorUpdateOne {
	return (&AutorClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Autor) Unwrap() *Autor {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Autor is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Autor) String() string {
	var builder strings.Builder
	builder.WriteString("Autor(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", signature=")
	builder.WriteString(a.Signature)
	builder.WriteByte(')')
	return builder.String()
}

// Autors is a parsable slice of Autor.
type Autors []*Autor

func (a Autors) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
