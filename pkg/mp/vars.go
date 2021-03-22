package mp

import "errors"

type ReadFileErr struct {
	error
}

type CSVParseErr struct {
	error
}

type XMLGenErr struct {
	error
}

type XMLWriteErr struct {
	error
}

var (
	// ErrReadFile read data file error
	ErrReadFile = &ReadFileErr{errors.New("read data file error")}
	// ErrCSVparse parse csv data error
	ErrCSVparse = &CSVParseErr{errors.New("parse csv data error")}
	// ErrXMLgen generate XML error
	ErrXMLgen = &XMLGenErr{errors.New("generate xml error")}
	// ErrXMLWrite write XML file error
	ErrXMLWrite = &XMLWriteErr{errors.New("write xml file [%s] error")}

	// ErrTimeOut processing takes too long thus timed out
	ErrTimeOut = errors.New("processing timeout.")
)
