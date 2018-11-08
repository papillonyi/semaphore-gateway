package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dimfeld/httptreemux"
)

func Load(mux *httptreemux.ContextMux, middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(middleware...)



}