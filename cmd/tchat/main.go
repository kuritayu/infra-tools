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
	}

	//TODO port番号は指定できるようにしたい
	app.Action = func(c *cli.Context) error {
		if c.Bool("c") {
			tchat.ClientExecute()
		} else {
			tchat.ServerExecute()
		}

		return nil
	}
	_ = app.Run(os.Args)
	os.Exit(0)

}
