package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/papillonyi/semaphore-gateway/server"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(middleware...)

	e.GET("/version", server.Version)
	e.GET("/healthz", server.Health)

	return e

}