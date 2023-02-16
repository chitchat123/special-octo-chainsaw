package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"primes/server"
	"time"
)

func main() {
	start := time.Now()

	logger := log.New(log.Writer(), "[PRIMES] ", 0)
	srv := server.NewServer(logger)

	srv.Router.Use(func(c *gin.Context) {
		logger.Println(c.Errors.Errors())
	})

	srv.Router.POST("", primeHandler(logger))

	srv.Router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"nums": cache})
	})

	logger.Printf("Started, %fs", time.Since(start).Seconds())
	srv.Start()
}
