package infrastructure

import (
	"context"
	"invoice/handler"
	"invoice/middleware"
	"invoice/repository"
	"invoice/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	invHandler := handler.NewInvoiceHandler(service.NewInvoiceService(repository.NewInvoiceRepository(db)))
	ordHandler := handler.NewOrderHandler(service.NewOrderService(repository.NewOrderRepository(db)))

	r := gin.Default()
	r.Use(middleware.ErrorCheck())

	inv := r.Group("/invoices")
	inv.GET("", invHandler.GetAllInvoices)
	inv.GET("/:id", invHandler.SelectInvoiceByID)
	inv.GET("/:id/items", ordHandler.DisplayOrders)
	inv.PUT("/:id/items/:itemid", ordHandler.EditOrder)
	inv.DELETE("/:id/items/:itemid", ordHandler.DeleteOrder)
	inv.POST("", invHandler.NewInvoice)
	inv.PUT("", invHandler.EditInvoice)

	return r
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	log.Println("Server exiting")
}
