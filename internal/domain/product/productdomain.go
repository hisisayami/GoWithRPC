package product

import (
	"context"

	"example.com/go-inventory-grpc/config"
	"example.com/go-inventory-grpc/internal/model"
	"example.com/go-inventory-grpc/internal/repository/product"
)

type ProductDomain interface {
	CreateProduct(ctx context.Context, productmodel model.Product) (model.Product, error)
}

type productDomain struct {
	productRepo product.Repository
}

func New(productRepo product.Repository) *ProductDomain {
	s := &productDomain{
		productRepo: productRepo,
	}

	return s
}

func (s *productDomain) CreateProduct(ctx context.Context, catId int, productmodel model.Product) (model.Product, error) {
	entClient := config.GetClient()

	product, err := product.New(ctx, entClient).CreateProduct(ctx, catId, productmodel)
	if err != nil {
		return model.Product{}, err
	}

	productResp := model.Product{
		ProductId:          product.ProductID,
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		ProductQuantity:    product.ProductQuantity,
		UnitPrice:          product.UnitPrice,
	}

	return productResp, nil
}
