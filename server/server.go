package server

import (
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/semaphore-gateway/version"
)

func Health(c *gin.Context)  {
	c.String(200, "")

}

func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": version.Version.String(),
	})
}