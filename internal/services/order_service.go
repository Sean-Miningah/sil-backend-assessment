package services

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
	"github.com/sean-miningah/sil-backend-assessment/internal/core/ports"
	"go.opentelemetry.io/otel"
)

type orderService struct {
	orderRepo   ports.OrderRepository
	productRepo ports.ProductRepository
}

func NewOrderService(orderRepo ports.OrderRepository, productRepo ports.ProductRepository) ports.OrderService {
	return &orderService{orderRepo: orderRepo, productRepo: productRepo}
}

func (s *orderService) CreateOrder(ctx context.Context, order *domain.Order) error {
	ctx, span := otel.Tracer("").Start(ctx, "OrderService.CreateOrder")
	defer span.End()

	return s.orderRepo.Create(ctx, order)
}

func (s *orderService) ListOrders(ctx context.Context) ([]domain.Order, error) {
	ctx, span := otel.Tracer("").Start(ctx, "OrderService.ListOrders")
	defer span.End()

	return s.orderRepo.List(ctx)
}

func (s *orderService) GetOrder(ctx context.Context, id uint) (*domain.Order, error) {
	ctx, span := otel.Tracer("").Start(ctx, "OrderService.GetOrder")
	defer span.End()

	return s.orderRepo.Get(ctx, id)
}

func (s *orderService) UpdateOrder(ctx context.Context, order *domain.Order) error {
	ctx, span := otel.Tracer("").Start(ctx, "OrderService.UpdateOrder")
	defer span.End()

	return s.orderRepo.Update(ctx, order)
}
