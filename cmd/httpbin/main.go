package main

import (
	"cncamp-module12/cmd/httpbin/router"
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"cncamp-module12/internal/config"
	"cncamp-module12/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	config.SetLogLevel("info")
}

func main() {
	if c, err := config.InitConfig(); err != nil {
		log.WithError(err).Error()
	} else {
		log.WithFields(log.Fields{
			"host":    c.ServerConfig.Host,
			"port":    c.ServerConfig.Port,
			"version": c.ServerConfig.Version,
		}).Debug("Server Configurations")

		log.Info("Log level is ", c.LogLevel)

		// Create context that listens for the interrupt signal from the OS.
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		gin.SetMode(c.GinMode)
		r := gin.Default()

		m := middleware.GinMonitor()
		// set middleware for gin
		m.Use(r)

		router.Register(r)

		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%s", c.ServerConfig.Host, c.ServerConfig.Port),
			Handler: r,
		}

		// Initializing the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.WithError(err).Error()
			}
		}()

		// Listen for the interrupt signal.
		<-ctx.Done()

		// Restore default behavior on the interrupt signal and notify user of shutdown.
		stop()
		log.Info("shutting down gracefully, press Ctrl+C again to force")

		// The context is used to inform the server it has 30 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.WithError(err).Error("Server forced to shutdown")
		}

		log.Info("Server exiting")
	}
}
