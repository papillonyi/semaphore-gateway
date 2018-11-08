package main

import (
	"github.com/papillonyi/semaphore-gateway/version"
	"github.com/urfave/cli"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "Semaphore-gateway"
	app.Version = version.Version.String()
	app.Action = server

	if err :=app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

