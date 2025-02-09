package ports

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProduct(ctx context.Context, id uint) (*domain.Product, error)
	ListProducts(ctx context.Context) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id uint) error
}

type CategoryService interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetCategory(ctx context.Context, id uint) (*domain.Category, error)
	ListCategories(ctx context.Context) ([]domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) error
	DeleteCategory(ctx context.Context, id uint) error
}
