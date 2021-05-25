package server

import (
	"context"
	"covid/summary"
	"covid/wongnai"
	"crypto/tls"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	addr = ":8080"
)

func Run() {
	log.SetFormatter(&log.JSONFormatter{})

	router := setUpServer()

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Info("Starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server")

	attempGracefulShutdown(srv)
}

func setUpServer() *gin.Engine {
	router := gin.Default()
	wongnaier := wongnai.New()
	summaryer := summary.New(&wongnaier)
	h := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	router.GET("/covid/summary", summaryer.GetSummary(h))

	router.GET("/healthy", healthCheck)
	return router
}

func attempGracefulShutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("Server forced to shutdown")
	}
	log.Info("Server existing")
}

type health struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
}

func healthCheck(c *gin.Context) {
	c.JSON(200, health{
		Status:    "OK",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Service:   "covid",
	})
}
