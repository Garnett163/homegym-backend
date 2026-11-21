package main

import (
	"net/http"
	"time"

	"github.com/Garnett163/homegym-backend/internal/http_handler"
	"github.com/gin-gonic/gin"
)

func newHTTPServer() *http.Server {
	router := gin.Default()

	router.GET("/api/healthcheck", http_handler.Healthcheck)

	return &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
