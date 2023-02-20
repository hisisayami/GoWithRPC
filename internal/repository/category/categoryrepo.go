package category

import (
	"context"

	"example.com/go-inventory-grpc/ent"
	"example.com/go-inventory-grpc/internal/model"
)

type Repository interface{}

type repository struct{}

type CategoryRepo struct {
	ctx    context.Context
	client *ent.Client
}

func New1() Repository {
	return &repository{}
}

func New(ctx context.Context, client *ent.Client) *CategoryRepo {
	return &CategoryRepo{
		ctx:    ctx,
		client: client,
	}
}

func (r *CategoryRepo) CreateCategory(ctx context.Context, categorymodel model.Category) (*ent.Category, error) {

	categoryCreated, err := r.client.Category.Create().
		SetCategoryID(categorymodel.CategoryId).
		SetCategoryName(categorymodel.CategoryName).
		SetCategoryDescription(categorymodel.CategoryDescription).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return categoryCreated, nil
}
