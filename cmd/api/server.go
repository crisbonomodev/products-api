package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"products/internal/products/handlers"
	"products/internal/products/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeServer() *http.Server {
	router := gin.Default()
	// will initialize a handler starting from data sources
	// which inject into repository layer

	// productRepository := repository.NewProductRepository(d.DB)

	// which inject into service layer

	productService := service.NewProductService(&service.ProductConfig{})
	// which inject into handler layer
	handlers.NewHandler(&handlers.Config{
		R:              router,
		ProductService: productService,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	return srv

}
