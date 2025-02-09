package services

import (
	"context"
	"testing"

	"github.com/sean-miningah/sil-backend-assessment/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, product *domain.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Get(ctx context.Context, id uint) (*domain.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func TestProductService_CreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	product := &domain.Product{
		Name:       "Test Product",
		Price:      99.99,
		CategoryID: 1,
	}

	mockRepo.On("Create", mock.Anything, product).Return(nil)

	err := service.CreateProduct(context.Background(), product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	expectedProduct := &domain.Product{
		ID:         1,
		Name:       "Test Product",
		Price:      99.99,
		CategoryID: 1,
	}

	mockRepo.On("Get", mock.Anything, uint(1)).Return(expectedProduct, nil)

	product, err := service.GetProduct(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}
