package cwmp

import (
	"strconv"
)

type GetParameterValues struct {
	baseMessage
	ParameterNames []string
}

type getParameterValuesBodyStruct struct {
	Body *getParameterValuesStruct `xml:"cwmp:GetParameterValues"`
}

type getParameterValuesStruct struct {
	ParameterNames *parameterNamesStruct `xml:"ParameterNames"`
}

type parameterNamesStruct struct {
	Type   string   `xml:"SOAP-ENC:arrayType,attr"`
	Values []string `xml:"string"`
}

func NewGetParameterValues() *GetParameterValues {
	return &GetParameterValues{}
}

func (msg *GetParameterValues) GetName() string {
	return "GetParameterValues"
}

func (msg *GetParameterValues) CreateXML() []byte {
	body := &getParameterValuesBodyStruct{
		&getParameterValuesStruct{
			&parameterNamesStruct{
				Type:   XsdString + "[" + strconv.Itoa(len(msg.ParameterNames)) + "]",
				Values: msg.ParameterNames,
			},
		},
	}
	return marshal(msg.GetID(), body)
}
