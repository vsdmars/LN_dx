package mp

import "encoding/xml"

// CFG2 XML namespace.
const (
	xmlNs  = "urn:com:linkedin:ns:configuration:source:1.0"
	xmlNsw = "urn:com:linkedin:ns:configuration:wildcard:1.0"
)

// CfgApplication is the base element for CFG2.
type CfgApplication struct {
	XMLName      xml.Name         `xml:"application"`
	ConfigSource *CfgConfigSource `xml:""` // pointer type goes first to save struct size due to padding
	Xmlns        string           `xml:"xmlns,attr"`
	Xmlnsw       string           `xml:"xmlns:w,attr"`
}

// CfgConfigSource is the embedded element for application element.
type CfgConfigSource struct {
	XMLName    xml.Name       `xml:"configuration-source"`
	Properties []*CfgProperty `xml:""`
}

// CfgProperty is the embedded element for configuration-source element.
type CfgProperty struct {
	XMLName xml.Name `xml:"property"`
	Set     *CfgSet  `xml:",omitempty"`           // pointer type goes first to save struct size due to padding
	Name    string   `xml:"name,attr,omitempty"`  // copy by value
	Value   string   `xml:"value,attr,omitempty"` // copy by value
}

// CfgSet is the embedded element for property element.
type CfgSet struct {
	XMLName xml.Name `xml:"set"`
	Values  []string `xml:"value"` // copy by value
}

// NewXMLApp creates CFG2 <application> element.
func NewXMLApp() *CfgApplication {
	cfgApp := &CfgApplication{}
	cfgApp.Xmlns = xmlNs
	cfgApp.Xmlnsw = xmlNsw

	cfgSource := &CfgConfigSource{}
	cfgApp.ConfigSource = cfgSource

	return cfgApp
}

// NewXMLSet creates CFG2 <set> element.
func NewXMLSet() *CfgSet {
	s := &CfgSet{}
	return s
}

// NewXMLProperty creates CFG2 <property> element with name and value.
//
// If value not provided, will create property element with sub element.
func NewXMLProperty(name string, value string, subElement interface{}) *CfgProperty {
	prop := &CfgProperty{}
	prop.Name = name

	if subElement != nil {
		switch val := subElement.(type) {
		case *CfgSet:
			prop.Set = val
		}
	} else {
		prop.Value = value
	}

	return prop
}
