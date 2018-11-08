package main

import (
	"github.com/urfave/cli"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)


func server(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	router := gin.Default()

	router.POST("")
}

