package infrastructure

import (
	"context"
	"errors"
	"invoice/handler"
	"invoice/repository"
	"invoice/service"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := newRouter()
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	gracefulShutdown(&srv)
}

func newRouter() *gin.Engine {
	invHandler := handler.NewInvoiceHandler(service.NewInvoiceService(repository.NewInvoiceRepository(newDBConnection())))

	r := gin.Default()
	r.Use()

	inv := r.Group("/invoices")
	inv.GET("/", invHandler.GetAllInvoice)
	inv.GET("/:id", invHandler.GetInvoiceByID)
	inv.POST("/", invHandler.NewInvoice)
	inv.PUT("/", invHandler.EditInvoice)

	return r
}

func gracefulShutdown(srv *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
	}

	log.Println("final")
}
