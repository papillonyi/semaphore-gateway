package router

import (
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/semaphore-gateway/server"
)

func Load(addr string, middleware ...gin.HandlerFunc) (err error) {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(middleware...)

	e.GET("/version", server.Version)
	e.GET("/healthz", server.Health)

	return e.Run(addr)
}
