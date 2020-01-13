package main

import (
	"fmt"
	"github.com/kuritayu/infra-tools/internal/lstar"
	"github.com/kuritayu/infra-tools/pkg"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lstar"
	app.Usage = "print tar information"
	app.Version = "1.0"

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

		err = pkg.ValidateTar(target)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		var t lstar.ArchiveInfo
		t = lstar.New(target)
		t.Print()

		return nil
	}
	_ = app.Run(os.Args)
	os.Exit(0)
}
