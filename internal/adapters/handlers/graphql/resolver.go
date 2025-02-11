package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/sean-miningah/sil-backend-assessment/internal/core/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productService ports.ProductService
	// categoryService ports.CategoryService
}

func NewResolver(ps ports.ProductService,

// cs ports.CategoryService
) *Resolver {
	return &Resolver{
		productService: ps,
		// categoryService: cs,
	}
}
