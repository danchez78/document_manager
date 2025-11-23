package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Server
//
//	@title				Document Manager
//	@version			1.0.0
//	@description		Service for managing documents for users
//
//	@host				127.0.0.1:50000
//	@BasePath			/api
//
//	@authorizationUrl	http://127.0.0.1
type Server struct {
	srv *echo.Echo
}

func New() *Server {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover(), middleware.Logger())

	if os.Getenv("APP_ENV") != "prod" {
		e.GET("/docs/*", echoSwagger.WrapHandler)
	}

	return &Server{srv: e}
}

func (s *Server) BasePath() *echo.Group {
	return s.srv.Group("/api")
}

func (s *Server) Start() error {
	log.Println("Server listening...")

	log.Println("Running server in HTTP mode...")
	if err := s.srv.Start(":50000"); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	log.Println("Shutting down server...")

	return s.srv.Shutdown(ctx)
}
