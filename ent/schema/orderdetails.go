package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OrderDetails holds the schema definition for the OrderDetails entity.
type OrderDetails struct {
	ent.Schema
}

// Fields of the OrderDetails.
func (OrderDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Int("OrderDetailId").
			Immutable(),
		field.Int("OrderId").
			Immutable(),
		field.Int("ProductId").
			Immutable(),
		field.Int("UnitPrice").
			Default(0),
		field.Int("Quantity").
			Default(0),
		field.Int("TotalPrice").
			Default(0),
	}
}

// Edges of the OrderDetails.
func (OrderDetails) Edges() []ent.Edge {
	return nil
}
