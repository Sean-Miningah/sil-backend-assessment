package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/handlers/graphql"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/handlers/rest"
	repo "github.com/sean-miningah/sil-backend-assessment/internal/adapters/repositories/postgres"
	"github.com/sean-miningah/sil-backend-assessment/internal/services"
	"github.com/sean-miningah/sil-backend-assessment/pkg/config"
	"github.com/sean-miningah/sil-backend-assessment/pkg/database"
	"github.com/sean-miningah/sil-backend-assessment/pkg/telemetry"
)

func main() {
	cfg := config.Load(".env")

	//Initialize tracer
	tp, err := telemetry.InitTracer(cfg.ServiceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer: %v", err)
		}
	}()

	// Construct the PostgreSQL connection string
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	// Initialize database connection
	db, err := database.NewPostgresDB(dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository
	productRepo := repo.NewProductRepository(db)
	orderRepo := repo.NewOrderRepository(db)

	// Initialize service
	productService := services.NewProductService(productRepo, orderRepo)
	orderService := services.NewOrderService(orderRepo, productRepo)

	// Initialize handler
	productHandler := rest.NewProductHandler(productService)
	orderHandler := rest.NewOrderHandler(orderService)

	// Initialize GraphQL handler
	graphqlHandler := graphql.NewHandler(productService, orderService)

	// Gin router setup
	router := gin.Default()

	// Register routes
	api := router.Group("/api/v1")
	{
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Get)
		api.POST("/products", productHandler.Create)
		api.PUT("/products/:id", productHandler.Update)
		api.DELETE("/products/:id", productHandler.Delete)

		// Order Routes
		api.GET("/orders", orderHandler.List)
		api.GET("/orders/:id", orderHandler.Get)
		api.POST("/orders", orderHandler.Create)
		api.PUT("/orders/:id", orderHandler.Update)
		api.DELETE("/order/:id", orderHandler.Delete)
	}

	router.POST("/graphql", graphqlHandler.GraphQL())
	if cfg.Environment == "development" {
		router.GET("/playground", graphqlHandler.Playground())
	}

	// Start the server
	if err := router.Run(cfg.Address); err != nil {
		log.Fatal(err)
	}
}
