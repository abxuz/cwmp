package cwmp

import (
	"encoding/xml"
	"strconv"
)

type SetParameterValues struct {
	Header          `xml:"-"`
	ParameterValues []*ParameterValueStruct `xml:"-"`
}

func NewSetParameterValues() *SetParameterValues {
	m := new(SetParameterValues)
	m.Header.RandomID()
	return m
}

func (m *SetParameterValues) SetParameter(name string, v string, t string) {
	value := new(ParameterValueStruct)
	value.Name = name
	value.Value.Type = t
	value.Value.Value = v
	m.ParameterValues = append(m.ParameterValues, value)
}

type setParameterValuesBody struct {
	ParameterList struct {
		Type            string                  `xml:"SOAP-ENC:arrayType,attr"`
		ParameterValues []*ParameterValueStruct `xml:"ParameterValueStruct,omitempty"`
	} `xml:"cwmp:SetParameterValues>ParameterList"`
	ParameterKey string `xml:"cwmp:SetParameterValues>ParameterKey"`
}

type ParameterValueStruct struct {
	Name  string `xml:"Name"`
	Value struct {
		Type  string `xml:"xsi:type,attr"`
		Value string `xml:",chardata"`
	} `xml:"Value"`
}

func (m *SetParameterValues) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	body := new(setParameterValuesBody)
	body.ParameterList.ParameterValues = m.ParameterValues
	body.ParameterList.Type = "cwmp:ParameterValueStruct[" + strconv.Itoa(len(m.ParameterValues)) + "]"
	return e.EncodeElement(body, start)
}

func (m *SetParameterValues) Response() *SetParameterValuesResponse {
	resp := new(SetParameterValuesResponse)
	resp.ID = m.ID
	return resp
}
