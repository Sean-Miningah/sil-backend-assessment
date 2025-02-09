package postgres

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	ctx, span := otel.Tracer("").Start(ctx, "ProductRepository.Create")
	defer span.End()

	return r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductRepository) Get(ctx context.Context, id uint) (*domain.Product, error) {
	ctx, span := otel.Tracer("").Start(ctx, "ProductRepository.Get")
	defer span.End()

	var product domain.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) List(ctx context.Context) ([]domain.Product, error) {
	ctx, span := otel.Tracer("").Start(ctx, "ProductRepository.List")
	defer span.End()

	var products []domain.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	ctx, span := otel.Tracer("").Start(ctx, "ProductRepository.Update")
	defer span.End()

	return r.db.WithContext(ctx).Save(product).Error
}

func (r *ProductRepository) Delete(ctx context.Context, id uint) error {
	ctx, span := otel.Tracer("").Start(ctx, "ProductRepository.Delete")
	defer span.End()

	return r.db.WithContext(ctx).Delete(&domain.Product{}, id).Error
}
