package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-miningah/sil-backend-assessment/internal/adpaters/handlers/rest"
	"github.com/sean-miningah/sil-backend-assessment/internal/adpaters/repositories/postgres"
	"github.com/sean-miningah/sil-backend-assessment/internal/services"
	"github.com/sean-miningah/sil-backend-assessment/pkg/config"
	"github.com/sean-miningah/sil-backend-assessment/pkg/database"
	"github.com/sean-miningah/sil-backend-assessment/pkg/telemetry"
)

func main() {
	cfg := config.Load()

	//Initialize tracer
	tp, err := telemetry.InitTracer(cfg.Telemetry.ServiceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer: %v", err)
		}
	}()

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	productRepo := postgres.NewProductRepository(db)

	// Initialize service
	productService := services.NewProductService(productRepo)

	// Initialize handler
	productHandler := rest.NewProductHandler(productService)

	// Initialize GraphQL handler
	graphqlHandler := graphql.NewHandler(productService)

	// Gin router setup
	router := gin.Default()

	// Register routes
	api := router.Group("/api/v1")
	{
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Get)
		api.POST("/products", productHandler.Create)
		api.PUT("/products/:id", productHandler.Update)
		api.Delete("/products/:id", productHandler.Delete)
	}

	router.POST("/graphql", graphqlHandler.Handle())

	// Start the server
	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatal(err)
	}
}
