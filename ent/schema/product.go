package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return *&[]ent.Field{
		field.Int("product_id"),
		field.String("product_name"),
		field.String("product_description"),
		field.Int("product_quantity"),
		field.Int("unit_price"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category_id", Category.Type).
			Ref("products").
			Unique(),
	}
}
