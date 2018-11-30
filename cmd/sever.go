package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/papillonyi/semaphore-gateway/router"
	"github.com/urfave/cli"
	"time"
)

func server(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	return router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
	)

}
