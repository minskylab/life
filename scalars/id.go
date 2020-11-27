package scalars

import "database/sql/driver"

// ID represents a generic GraphQL ID
type ID string

// Value implements the driver.Valuer interface
func (id ID) Value() (driver.Value, error) {
	return id, nil
}

// EmptyID is a simple empty string
var EmptyID = ID("")
