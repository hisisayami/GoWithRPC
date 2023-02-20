package orderdetails

type Repository interface {
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func CreateOrderDetails() error {
	return nil
}

func GetAllOrderDetails() error {
	return nil
}

func UpdateOrderDetails() error {
	return nil
}

func DeleteOrderDetails() error {
	return nil
}
