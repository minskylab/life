// Code generated by entc, DO NOT EDIT.

package autor

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/minskylab/life/example/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Signature applies equality check predicate on the "signature" field. It's identical to SignatureEQ.
func Signature(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Autor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Autor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Autor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Autor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SignatureEQ applies the EQ predicate on the "signature" field.
func SignatureEQ(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// SignatureNEQ applies the NEQ predicate on the "signature" field.
func SignatureNEQ(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSignature), v))
	})
}

// SignatureIn applies the In predicate on the "signature" field.
func SignatureIn(vs ...string) predicate.Autor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Autor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSignature), v...))
	})
}

// SignatureNotIn applies the NotIn predicate on the "signature" field.
func SignatureNotIn(vs ...string) predicate.Autor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Autor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSignature), v...))
	})
}

// SignatureGT applies the GT predicate on the "signature" field.
func SignatureGT(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSignature), v))
	})
}

// SignatureGTE applies the GTE predicate on the "signature" field.
func SignatureGTE(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSignature), v))
	})
}

// SignatureLT applies the LT predicate on the "signature" field.
func SignatureLT(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSignature), v))
	})
}

// SignatureLTE applies the LTE predicate on the "signature" field.
func SignatureLTE(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSignature), v))
	})
}

// SignatureContains applies the Contains predicate on the "signature" field.
func SignatureContains(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSignature), v))
	})
}

// SignatureHasPrefix applies the HasPrefix predicate on the "signature" field.
func SignatureHasPrefix(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSignature), v))
	})
}

// SignatureHasSuffix applies the HasSuffix predicate on the "signature" field.
func SignatureHasSuffix(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSignature), v))
	})
}

// SignatureIsNil applies the IsNil predicate on the "signature" field.
func SignatureIsNil() predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSignature)))
	})
}

// SignatureNotNil applies the NotNil predicate on the "signature" field.
func SignatureNotNil() predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSignature)))
	})
}

// SignatureEqualFold applies the EqualFold predicate on the "signature" field.
func SignatureEqualFold(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSignature), v))
	})
}

// SignatureContainsFold applies the ContainsFold predicate on the "signature" field.
func SignatureContainsFold(v string) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSignature), v))
	})
}

// HasTodos applies the HasEdge predicate on the "todos" edge.
func HasTodos() predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TodosTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TodosTable, TodosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTodosWith applies the HasEdge predicate on the "todos" edge with a given conditions (other predicates).
func HasTodosWith(preds ...predicate.Todo) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TodosInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TodosTable, TodosColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Autor) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Autor) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Autor) predicate.Autor {
	return predicate.Autor(func(s *sql.Selector) {
		p(s.Not())
	})
}