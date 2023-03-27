// Code generated by ent, DO NOT EDIT.

package category

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldCategoryName holds the string denoting the category_name field in the database.
	FieldCategoryName = "category_name"
	// FieldCategoryDescription holds the string denoting the category_description field in the database.
	FieldCategoryDescription = "category_description"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "category_products"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCategoryID,
	FieldCategoryName,
	FieldCategoryDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}