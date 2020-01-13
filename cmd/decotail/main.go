package main

import (
	"fmt"
	"github.com/kuritayu/infra-tools/internal/decotail"
	"github.com/kuritayu/infra-tools/pkg"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "decotail"
	app.Usage = "decorated tail"
	app.Version = "1.0"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "t",
		},
		&cli.StringFlag{
			Name:  "k",
			Usage: "add color to text",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Args().Get(0) == "" {
			help := []string{"", "--help"}
			_ = app.Run(help)
			os.Exit(1)
		}

		target := c.Args().Get(0)

		err := pkg.ValidateFile(target)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		timestamp := c.Bool("t")
		keyword := c.String("k")

		var p decotail.Printer
		p = decotail.New(target, timestamp, keyword)
		err = p.Execute()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		return nil
	}
	_ = app.Run(os.Args)
	os.Exit(0)

}
