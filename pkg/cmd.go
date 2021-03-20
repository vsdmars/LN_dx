package pkg

import (
	"os"

	"github.com/urfave/cli/v2"
)

func newAuthors() []*cli.Author {
	return []*cli.Author{
		{Name: "shchang", Email: "shchang@linkedin.com"},
	}
}

func newFlags() []cli.Flag {
	flags := []cli.Flag(nil)
	flags = append(flags, &cli.StringFlag{
		Name:     "p",
		Usage:    "`multiproduct` name",
		Required: true,
	}, &cli.StringFlag{
		Name:     "a",
		Usage:    "`action` type",
		Required: true,
	})

	return flags
}

func Cmd() error {
	app := cli.NewApp()
	app.Name = "dx"

	app.Authors = newAuthors()
	app.Copyright = "2021 Linkedin Corp."

	app.ArgsUsage = "input_file output_file"
	app.Description =
		`convert data to cfg2 xml
e.g dx -p multiproduct -a action input_file output_file`
	app.Usage = "dx -p multiproduct -a action input_file output_file"
	app.UsageText = "dx -p multiproduct -a action input_file output_file"

	app.EnableBashCompletion = true

	app.Flags = newFlags()

	app.Action = func(ctx *cli.Context) error {
		mp := ctx.String("p")
		action := ctx.String("a")
		return dispatch(mp, action)
	}

	return app.Run(os.Args)
}
