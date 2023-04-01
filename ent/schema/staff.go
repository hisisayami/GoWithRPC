package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		// ent frame work has validator build in
		// the code below indicate that we define a field with type of string
		// named "email" and we validate it as required/ not empty
		field.String("email").
			NotEmpty(),

		// Giving the field name of type string to have default value of unkown if no value is supplied
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return nil
}
