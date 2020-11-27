// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AutorsColumns holds the columns for the "autors" table.
	AutorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "signature", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// AutorsTable holds the schema information for the "autors" table.
	AutorsTable = &schema.Table{
		Name:        "autors",
		Columns:     AutorsColumns,
		PrimaryKey:  []*schema.Column{AutorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "done", Type: field.TypeBool},
		{Name: "autor_todos", Type: field.TypeString, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "todos_autors_todos",
				Columns: []*schema.Column{TodosColumns[3]},

				RefColumns: []*schema.Column{AutorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AutorsTable,
		TodosTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = AutorsTable
}