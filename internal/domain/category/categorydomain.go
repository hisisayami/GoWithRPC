package category

import (
	"context"

	"example.com/go-inventory-grpc/config"
	"example.com/go-inventory-grpc/internal/model"
	"example.com/go-inventory-grpc/internal/repository/category"
)

type CategoryDomain interface {
	CreateCategory(ctx context.Context, categoryModel model.Category) (model.Category, error)
}

type categoryDomain struct {
	categoryRepo category.Repository
}

func New(categoryRepo category.Repository) CategoryDomain {
	s := &categoryDomain{
		categoryRepo: categoryRepo,
	}
	return s
}

func (s *categoryDomain) CreateCategory(ctx context.Context, categoryModel model.Category) (model.Category, error) {
	entClient := config.GetClient()

	category, err := category.New(ctx, entClient).CreateCategory(ctx, categoryModel)
	if err != nil {
		return model.Category{}, err
	}

	categorydbResponse := model.Category{
		CategoryId:          category.CategoryID,
		CategoryName:        category.CategoryName,
		CategoryDescription: category.CategoryDescription,
	}
	return categorydbResponse, nil

}
