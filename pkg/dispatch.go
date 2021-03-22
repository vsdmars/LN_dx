package pkg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	errMP            = "mp: [%s] does not exist"
	errAction        = "action: [%s] not defined"
	errMpHasNoAction = "mp: [%s] has no action: [%s] defined"
	errArgCount      = "missing input file or output file"
)
const runnerDuration = 3 * time.Second

func dispatch(mp string, action string, args cli.Args) error {
	imp, ok := mpToMp[mp]
	if !ok {
		return errors.New(fmt.Sprintf(errMP, mp))
	}

	iaction, ok := actionToAction[action]
	if !ok {
		return errors.New(fmt.Sprintf(errAction, action))
	}

	runner, ok := mpActionRunner[imp][iaction]
	if !ok {
		return errors.New(fmt.Sprintf(errMpHasNoAction, mp, action))
	}

	if args.Len() < 2 {
		return errors.New(errArgCount)
	}

	runner.InputFile(args.Get(0))
	runner.OutputFile(args.Get(1))

	ctx, cancel := context.WithTimeout(context.Background(), runnerDuration)
	defer cancel()

	select {
	case rc := <-runner.Run(ctx):
		return rc.Err
	}
}
