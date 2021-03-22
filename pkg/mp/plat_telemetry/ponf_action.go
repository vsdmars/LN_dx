package plat_telemetry

import (
	"bufio"
	"context"
	"dx/pkg/mp"
	"encoding/xml"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	metaKey    = "ponfMetaData"
	featureKey = "ponfFeatures"
)

type Runner struct {
	inputFile  string
	outputFile string
}

func (r *Runner) createCFG2() ([]byte, error) {
	file, err := os.Open(r.inputFile)
	if err != nil {
		return nil, mp.ErrReadFile
	}
	defer file.Close()

	app := mp.NewXMLApp()
	values := mp.NewXMLSet()

	scanner := bufio.NewScanner(file)
	metaRead := false

	for scanner.Scan() {
		if !metaRead {
			metaRead = true
			app.ConfigSource.Properties = append([]*mp.CfgProperty(nil),
				mp.NewXMLProperty(metaKey, scanner.Text(), nil))

		} else {
			values.Values = append(values.Values, scanner.Text())
		}
	}

	app.ConfigSource.Properties = append(app.ConfigSource.Properties,
		mp.NewXMLProperty(featureKey, "", values))

	out, err := xml.MarshalIndent(app, " ", "  ")
	if err != nil {
		return nil, mp.ErrXMLgen
	}

	return out, nil
}

func (r *Runner) writeCFG2(content []byte) error {
	f, err := os.OpenFile(r.outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return mp.ErrXMLWrite

	}
	defer f.Close()

	_, err = f.WriteString(string(content))
	if err != nil {
		return mp.ErrXMLWrite
	}

	return nil
}

func (r *Runner) OutputFile(fileName string) {
	r.outputFile = fileName
}

func (r *Runner) InputFile(fileName string) {
	r.inputFile = fileName
}

func (r *Runner) Run(ctx context.Context) mp.ResultC {
	result := make(chan mp.Result)
	running := make(chan mp.Result)

	go func() {
		b, err := r.createCFG2()
		if err != nil {
			running <- mp.Result{Err: err}
			return
		}

		err = r.writeCFG2(b)
		if err != nil {
			running <- mp.Result{Err: err}
			return
		}

		running <- mp.Result{}
	}()

	go func() {
		select {
		case <-ctx.Done():
			result <- mp.Result{Err: mp.ErrTimeOut}
		case processed := <-running:
			result <- processed
		}
	}()

	return result
}

func Run(args cli.Args) error {
	print("ha\n")
	return nil
}
