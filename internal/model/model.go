package model

type Staff struct {
	Id    int
	Name  string
	Email string
}

type Category struct {
	CategoryId          int
	CategoryName        string
	CategoryDescription string
}

type Product struct {
	ProductId          int
	CategoryId         int
	ProductName        string
	ProductDescription string
	ProductQuantity    int
	UnitPrice          int
}
