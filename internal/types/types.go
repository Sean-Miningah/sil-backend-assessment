package types

import "github.com/sean-miningah/sil-backend-assessment/internal/models"

type ProductStore interface {
	FindAll() ([]models.Product, error)
	FindById(id int) (models.Product, error)
	FindByName(name string) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(product models.Product) (models.Product, error)
}

type CategoryStore interface {
	FindAll() ([]models.Category, error)
	FindById(id int) (models.Category, error)
	FindByName(name string) (models.Category, error)
	FindByParentId(parentId int) ([]models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(category models.Category) (models.Category, error)
}
