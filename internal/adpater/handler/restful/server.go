package restful

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sean-minngah/sil-backend-assessment/internal/core/util"
)

const TypeRestful = "restful"

type Server struct {
	config  util.Config
	router  *gin.Engine
	service port.Service
}

func NewServer(config util.Config, service port.Service) port.Server {
	server := &Server{
		config:  config,
		service: service,
	}

	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.use(gin.Recovery(), cors.Default())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "page not found",
		})
	})

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "no method found"})
	})

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}

func (server *Server) Start() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.config.Port),
		Handler: server.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}
