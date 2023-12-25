package service

import (
	"context"

	"github.com/geraldojalves/go-grpc/internal/database"
	"github.com/geraldojalves/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		categoryDB: categoryDB,
	}
}

func (c *CategoryService) ListCategories(ctx context.Context, _ *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.categoryDB.FindAll()
	if err != nil {
		panic(err)
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{
		Categories: categoriesResponse,
	}, nil
}
