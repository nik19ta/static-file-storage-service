package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	middleware "static/pkg/middleware"

	upload "static/upload"
	uploadhttp "static/upload/delivery/http"
	uploadfile "static/upload/repository"
	uploadusecase "static/upload/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server

	uploadApiLoadUC upload.UseCase
}

func NewApp() *App {

	uploadRepo := uploadfile.NewHeavyApiLoadRepository()

	return &App{
		uploadApiLoadUC: uploadusecase.NewHeavyApiLoadUseCase(uploadRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.Use(middleware.CORSMiddleware())

	router.LoadHTMLGlob("templates/*.html")

	router.StaticFS("/assets", http.Dir("templates/static"))
	router.StaticFS("/style", http.Dir("templates/style"))
	router.StaticFS("/script", http.Dir("templates/script"))

	uploadhttp.RegisterHTTPEndpoints(router, a.uploadApiLoadUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
