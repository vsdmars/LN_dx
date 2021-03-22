package mp

import (
	"context"
)

type Actioner interface {
	OutputFile(filename string)
	InputFile(filename string)
	Run(ctx context.Context) ResultC
}

type Result struct {
	Err error
}

type ResultC <-chan Result
