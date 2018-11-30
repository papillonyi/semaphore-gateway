package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/papillonyi/semaphore-gateway/version"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "8080",
			Usage: "port for the app",
		},
	}

	app.Name = "Semaphore-gateway"
	app.Version = version.Version.String()
	app.Action = server

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
