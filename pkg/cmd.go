package pkg

import (
	"os"

	"github.com/urfave/cli/v2"
)

const (
	multiProduct = "multiproduct"
	action       = "action"
)

func newAuthors() []*cli.Author {
	return []*cli.Author{
		{Name: "shchang", Email: "shchang@linkedin.com"},
	}
}

func newFlags() []cli.Flag {
	flags := []cli.Flag(nil)
	flags = append(flags, &cli.StringFlag{
		Name:     multiProduct,
		Aliases:  []string{"p"},
		Usage:    "`multiproduct` name",
		Required: true,
	}, &cli.StringFlag{
		Name:     action,
		Aliases:  []string{"a"},
		Usage:    "`action` type",
		Required: true,
	})

	return flags
}

func extractFlag(ctx *cli.Context) (m string, a string) {
	m = ctx.String(multiProduct)
	a = ctx.String(action)
	return
}

// Cmd is the entry point for program dx
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
		mp, action := extractFlag(ctx)
		return dispatch(mp, action, ctx.Args())
	}

	return app.Run(os.Args)
}
