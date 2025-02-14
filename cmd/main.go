package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/handlers/graphql"
	"github.com/sean-miningah/sil-backend-assessment/internal/adapters/handlers/rest"
	repo "github.com/sean-miningah/sil-backend-assessment/internal/adapters/repositories/postgres"
	"github.com/sean-miningah/sil-backend-assessment/internal/services"
	"github.com/sean-miningah/sil-backend-assessment/pkg/config"
	"github.com/sean-miningah/sil-backend-assessment/pkg/database"
	"github.com/sean-miningah/sil-backend-assessment/pkg/telemetry"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/client"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

// Application should
// Input and upload products with their various categories
// Return average product price for a category
// Make Order

// Auth using openid connect
// When order is made send customer sms alerting them using Africa's Talking API
// Send Administrator an email about order placed
// Deploy to k8s
// Write Docs

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

	// Initialize Zitadel client
	conf := zitadel.New(cfg.ZitadelIssuerURL)

	authZ, err := authorization.New(ctx, conf, oauth.DefaultAuthorization(cfg.ZitadelKey))
	if err != nil {
		log.Error("zitadel sdk could not initialize authorization", "error", err)
		os.Exit(1)
	}

	mw := middleware.New(authZ)

	c, err := client.New(ctx, conf)
	if err != nil {
		log.Error("zitadel sdk could not initialize authorization", "error", err)
		os.Exit(1)
	}

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
