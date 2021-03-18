package main

import (
	"bufio"
	"encoding/xml"
	"errors"
	"os"
)

var (
	errArgsCnt  = errors.New("usage: input_csv_file_name output_xml_file_name")
	errReadFile = errors.New("read data file error")
	errCSVRead  = errors.New("parse csv data error")
	errXMLGen   = errors.New("generate xml error")
	errXMLWrite = errors.New("write xml file error")
)

type errType int

const (
	fileErr errType = iota
	csvReadErr
	xmlGenErr
	xmlWriteErr
)

// CFG2 XML namespace
const (
	xmlNs  = "urn:com:linkedin:ns:configuration:source:1.0"
	xmlNsw = "urn:com:linkedin:ns:configuration:wildcard:1.0"
)

const (
	metaKey    = "ponfMetaData"
	featureKey = "ponfFeatures"
)

func checkErr(e error, et errType) {
	if e != nil {
		switch et {

		case fileErr:
			panic(errReadFile)

		case csvReadErr:
			panic(errCSVRead)

		case xmlGenErr:
			panic(errXMLGen)

		case xmlWriteErr:
			panic(errXMLWrite)
		}
	}
}

// CfgApplication is the base element for CFG2
type CfgApplication struct {
	XMLName      xml.Name         `xml:"application"`
	Xmlns        string           `xml:"xmlns,attr"`
	Xmlnsw       string           `xml:"xmlns:w,attr"`
	ConfigSource *CfgConfigSource `xml:""`
}

// CfgConfigSource is the embedded element for application element
type CfgConfigSource struct {
	XMLName    xml.Name       `xml:"configuration-source"`
	Properties []*CfgProperty `xml:""`
}

// CfgProperty is the embedded element for configuration-source element
type CfgProperty struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr,omitempty"`  // copy by value
	Value   string   `xml:"value,attr,omitempty"` // copy by value
	Set     *CfgSet  `xml:",omitempty"`
}

// CfgSet is the embedded element for property element
type CfgSet struct {
	XMLName xml.Name `xml:"set"`
	Values  []string `xml:"value"` // copy by value
}

func newXMLApp() *CfgApplication {
	cfgApp := &CfgApplication{}
	cfgApp.Xmlns = xmlNs
	cfgApp.Xmlnsw = xmlNsw

	cfgSource := &CfgConfigSource{}
	cfgApp.ConfigSource = cfgSource

	return cfgApp
}

func newProperty(name string, value string, cfgSet *CfgSet) *CfgProperty {
	prop := &CfgProperty{}

	if cfgSet != nil {
		prop.Name = name
		prop.Set = cfgSet
	} else {
		prop.Name = name
		prop.Value = value
	}

	return prop
}

func newSet() *CfgSet {
	s := &CfgSet{}
	return s
}

func createXML(csvFileName string) ([]byte, error) {
	file, err := os.Open(csvFileName)
	checkErr(err, fileErr)
	defer file.Close()

	cfg2 := newXMLApp()
	values := newSet()

	scanner := bufio.NewScanner(file)
	metaRead := false

	for scanner.Scan() {
		if !metaRead {
			metaRead = true
			cfg2.ConfigSource.Properties = append([]*CfgProperty(nil), newProperty(metaKey, scanner.Text(), nil))

		} else {
			values.Values = append(values.Values, scanner.Text())
		}
	}

	cfg2.ConfigSource.Properties = append(cfg2.ConfigSource.Properties, newProperty(featureKey, "", values))

	out, err := xml.MarshalIndent(cfg2, " ", "  ")

	return out, err
}

func writeXML(xmlFileName string, xmlData []byte) error {
	f, err := os.OpenFile(xmlFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	checkErr(err, xmlWriteErr)
	defer f.Close()

	_, err = f.WriteString(string(xmlData))
	return err
}

func main() {
	if len(os.Args) == 1 || len(os.Args[1:]) > 2 {
		panic(errArgsCnt)
	}

	args := os.Args[1:]
	dataFile := args[0]
	xmlFile := args[1]

	xmlBytes, err := createXML(dataFile)
	checkErr(err, xmlGenErr)

	checkErr(writeXML(xmlFile, xmlBytes), xmlWriteErr)
}
