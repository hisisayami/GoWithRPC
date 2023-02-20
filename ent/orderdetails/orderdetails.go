// Code generated by ent, DO NOT EDIT.

package orderdetails

const (
	// Label holds the string label denoting the orderdetails type in the database.
	Label = "order_details"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderDetailId holds the string denoting the orderdetailid field in the database.
	FieldOrderDetailId = "order_detail_id"
	// FieldOrderId holds the string denoting the orderid field in the database.
	FieldOrderId = "order_id"
	// FieldProductId holds the string denoting the productid field in the database.
	FieldProductId = "product_id"
	// FieldUnitPrice holds the string denoting the unitprice field in the database.
	FieldUnitPrice = "unit_price"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldTotalPrice holds the string denoting the totalprice field in the database.
	FieldTotalPrice = "total_price"
	// Table holds the table name of the orderdetails in the database.
	Table = "order_details"
)

// Columns holds all SQL columns for orderdetails fields.
var Columns = []string{
	FieldID,
	FieldOrderDetailId,
	FieldOrderId,
	FieldProductId,
	FieldUnitPrice,
	FieldQuantity,
	FieldTotalPrice,
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

var (
	// DefaultUnitPrice holds the default value on creation for the "UnitPrice" field.
	DefaultUnitPrice int
	// DefaultQuantity holds the default value on creation for the "Quantity" field.
	DefaultQuantity int
	// DefaultTotalPrice holds the default value on creation for the "TotalPrice" field.
	DefaultTotalPrice int
)
