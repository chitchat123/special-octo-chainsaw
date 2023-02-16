package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"primes/server"
	"time"
)

func main() {
	var (
		start  = time.Now()
		logger = log.New(log.Writer(), "[PRIMES] ", log.LstdFlags)
		srv    = server.NewServer(logger)
	)

	srv.Router.POST("", primeHandler(logger))

	srv.Router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"nums": cache})
	})

	logger.Printf("Started, %fs", time.Since(start).Seconds())

	srv.Start()
}
