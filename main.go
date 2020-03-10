package main

import (
	"github.com/urfave/cli"
	"gproxy/internal/commands"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "douyacun"
	app.Version = "v0.3.10"
	app.Commands = []cli.Command{
		commands.Start,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
