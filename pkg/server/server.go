package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/server/routes"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

var srv HttpServer

type Config struct {
	Port     uint16
	BasePath string
	Mode     string
}

func Setup(config *Config) *HttpServer {
	gin.SetMode(config.Mode)
	handler := gin.New()
	routes.Setup(handler, config.BasePath)
	srv.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &srv
}
func (s *HttpServer) ServeAsync() {
	log.Info().Msgf("Server is up and running at %s", s.server.Addr)
	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Failure in http.Server.ListenAndServe(): %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server Shutdown: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info().Msg("timeout of 5 seconds.")
	}

	log.Info().Msg("Server exiting")
}
