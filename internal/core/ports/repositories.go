package ports

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	Get(ctx context.Context, id uint) (*domain.Product, error)
	List(ctx context.Context) ([]domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id uint) error
}

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	Get(ctx context.Context, id uint) (*domain.Category, error)
	List(ctx context.Context) ([]domain.Category, error)
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id uint) error
}

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	List(ctx context.Context) ([]domain.Order, error)
	Get(ctx context.Context, id uint) (*domain.Order, error)
	Update(ctx context.Context, order *domain.Order) error
	Delete(ctx context.Context, id uint) error
	DeleteOrderItems(ctx context.Context, orderID uint) error
}
