package cwmp

import (
	"encoding/xml"
	"strconv"
)

type GetParameterValues struct {
	Header         `xml:"-"`
	ParameterNames []string `xml:"-"`
}

func NewGetParameterValues() *GetParameterValues {
	m := new(GetParameterValues)
	m.Header.RandomID()
	return m
}

type getParameterValuesBody struct {
	ParameterNames struct {
		Type   string   `xml:"SOAP-ENC:arrayType,attr"`
		Values []string `xml:"string"`
	} `xml:"cwmp:GetParameterValues>ParameterNames"`
}

func (m *GetParameterValues) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	body := new(getParameterValuesBody)
	body.ParameterNames.Type = "cwmp:string[" + strconv.Itoa(len(m.ParameterNames)) + "]"
	body.ParameterNames.Values = m.ParameterNames
	return e.EncodeElement(body, start)
}

func (m *GetParameterValues) Response() *GetParameterValuesResponse {
	resp := new(GetParameterValuesResponse)
	resp.ID = m.ID
	return resp
}
