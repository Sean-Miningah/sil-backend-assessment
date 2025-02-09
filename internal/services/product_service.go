package services

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
	"github.com/sean-miningah/sil-backend-assessment/internal/core/ports"
	"go.opentelemetry.io/otel"
)

type productService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, product *domain.Product) error {
	ctx, span := otel.Tracer("").Start(ctx, "ProductService.CreateProduct")
	defer span.End()

	return s.repo.Create(ctx, product)
}

func (s *productService) GetProduct(ctx context.Context, id uint) (*domain.Product, error) {
	ctx, span := otel.Tracer("").Start(ctx, "ProductService.GetProduct")
	defer span.End()

	return s.repo.Get(ctx, id)
}

func (s *productService) ListProducts(ctx context.Context) ([]domain.Product, error) {
	ctx, span := otel.Tracer("").Start(ctx, "ProductService.ListProducts")
	defer span.End()

	return s.repo.List(ctx)
}

func (s *productService) UpdateProduct(ctx context.Context, product *domain.Product) error {
	ctx, span := otel.Tracer("").Start(ctx, "ProductService.UpdateProduct")
	defer span.End()

	return s.repo.Update(ctx, product)
}

func (s *productService) DeleteProduct(ctx context.Context, id uint) error {
	ctx, span := otel.Tracer("").Start(ctx, "ProductService.DeleteProduct")
	defer span.End()

	return s.repo.Delete(ctx, id)
}
