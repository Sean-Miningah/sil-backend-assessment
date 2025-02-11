package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/handlers/graphql"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/handlers/rest"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/repositories/postgres"
	"github.com/sean-miningah/sil-backend-assessment/internal/services"
	"github.com/sean-miningah/sil-backend-assessment/pkg/config"
	"github.com/sean-miningah/sil-backend-assessment/pkg/database"
	"github.com/sean-miningah/sil-backend-assessment/pkg/telemetry"
)

func main() {
	cfg := config.Load(".env")

	fmt.Printf("Loaded Config: %+v\n", cfg)

	// Print specific fields
	fmt.Println("Environment:", cfg.Environment)
	fmt.Println("Server Address:", cfg.Address)
	fmt.Println("Database Host:", cfg.DBHost)
	fmt.Println("Database Port:", cfg.DBPort)
	fmt.Println("Database User:", cfg.DBUser)
	fmt.Println("Database Name:", cfg.DBName)
	fmt.Println("Telemetry Service Name:", cfg.ServiceName)
	fmt.Println("Jaeger Endpoint:", cfg.JaegerEndpoint)
	fmt.Println("Prometheus Port:", cfg.PrometheusPort)

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
		api.DELETE("/products/:id", productHandler.Delete)
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
