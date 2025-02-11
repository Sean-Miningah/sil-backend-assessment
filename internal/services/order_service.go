package services

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
	"github.com/sean-miningah/sil-backend-assessment/internal/core/ports"
	"go.opentelemetry.io/otel"
)

type orderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) ports.OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(ctx context.Context, order *domain.Order) error {
	ctx, span := otel.Tracer("").Start(ctx, "OrderService.CreateOrder")
	defer span.End()

	return s.repo.Create(ctx, order)
}

func (s *orderService) ListOrders(ctx context.Context) ([]domain.Order, error) {
	ctx, span := otel.Tracer("").Start(ctx, "OrderService.ListOrders")
	defer span.End()

	return s.repo.List(ctx)
}
