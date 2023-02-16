package server

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	l      *log.Logger
	Router *gin.Engine
}

func NewServer(l *log.Logger) *Server {
	gin.DefaultWriter = l.Writer()
	gin.SetMode(gin.ReleaseMode)
	var (
		router = gin.Default()
	)
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Authorization", "Origin"},
	}))

	return &Server{l, router}
}

func (s *Server) Start() { // create a new server
	srv := http.Server{
		Addr:         ":8080",                              // configure the bind address
		Handler:      s.Router,                             // set the default handler
		ErrorLog:     log.New(s.l.Writer(), "[server]", 0), // set the logger for the server
		ReadTimeout:  5 * time.Second,                      // max time to read request from the client
		WriteTimeout: 10 * time.Second,                     // max time to write response to the client
		IdleTimeout:  120 * time.Second,                    // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		s.l.Println("Starting server on port 8080")

		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 2)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.l.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.l.Panic("Server forced to shutdown:", err)
	}

	s.l.Println("Server exiting")
}
