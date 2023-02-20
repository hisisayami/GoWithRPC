package product

import (
	"context"

	"example.com/go-inventory-grpc/ent"
	"example.com/go-inventory-grpc/internal/model"
)

type Repository interface{}

type repository struct{}

type ProductRepo struct {
	ctx    context.Context
	client *ent.Client
}

func New1() *repository {
	return &repository{}
}

func New(ctx context.Context, client *ent.Client) *ProductRepo {
	return &ProductRepo{
		ctx:    ctx,
		client: client,
	}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, catId int32, productmodel model.Product) (*ent.Product, error) {

	productCreated, err := r.client.Product.Create().
		SetProductID(productmodel.ProductId).
		SetProductName(productmodel.ProductName).
		SetProductDescription(productmodel.ProductDescription).
		SetProductQuantity(productmodel.ProductQuantity).
		SetUnitPrice(productmodel.UnitPrice).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return productCreated, nil
}
