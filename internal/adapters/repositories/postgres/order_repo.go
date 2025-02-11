package repo

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *domain.Order) error {
	ctx, span := otel.Tracer("").Start(ctx, "OrderRepository.Create")
	defer span.End()

	return r.db.WithContext(ctx).Create(order).Error
}

func (r *OrderRepository) List(ctx context.Context) ([]domain.Order, error) {
	ctx, span := otel.Tracer("").Start(ctx, "OrderRepository.List")
	defer span.End()

	var orders []domain.Order
	err := r.db.WithContext(ctx).Find(&orders).Error
	return orders, err
}
