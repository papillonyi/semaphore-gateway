package server

import (
	"fmt"
	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/semaphore-gateway/semaphore"
	"github.com/papillonyi/semaphore-gateway/version"
	"net/http"
)

func Health(c *gin.Context) {
	c.String(200, "")

}

func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": version.Version.String(),
	})
}

func Semaphore(c *gin.Context) {
	var downloadTask semaphore.DownloadTask
	if err := c.ShouldBindJSON(&downloadTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := downloadTask.GetNewTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	task.Login()
	token := task.GetToken()

	c.JSON(200, gin.H{
		"task":  task,
		"token": token,
	})
}
