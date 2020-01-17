package main

import (
	"github.com/kuritayu/infra-tools/internal/tchat"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tchat"
	app.Usage = "chat tool by terminal"
	app.Version = "1.0"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "c",
		},
		&cli.IntFlag{
			Name:  "p",
			Value: 7777,
		},
	}

	app.Action = func(c *cli.Context) error {
		port := c.Int("p")
		if c.Bool("c") {
			tchat.ClientExecute(port)
		} else {
			tchat.ServerExecute(port)
		}

		return nil
	}
	_ = app.Run(os.Args)
	os.Exit(0)

}
