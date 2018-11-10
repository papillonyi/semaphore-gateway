package main

import (
	"github.com/urfave/cli"
	"github.com/Sirupsen/logrus"
	"github.com/papillonyi/semaphore-gateway/router"
	"github.com/gin-gonic/contrib/ginrus"
	"time"
	"golang.org/x/sync/errgroup"
	"net/http"
	"crypto/tls"
)


func server(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}


	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
	)

	var g errgroup.Group

	if c.String("server-cert") != "" {
		g.Go(func() error {
			serve := &http.Server{
				Addr:    ":https",
				Handler: handler,
				TLSConfig: &tls.Config{
					NextProtos: []string{"http/1.1"}, // disable h2 because Safari :(
				},
			}
			return serve.ListenAndServeTLS(
				c.String("server-cert"),
				c.String("server-key"),
			)
		})
		return g.Wait()
	}

	if !c.Bool("lets-encrypt") {
		return http.ListenAndServe(
			c.String("server-addr"),
			handler,
		)
	}
}

