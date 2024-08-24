package cwmp

import (
	"strconv"
)

type SetParameterValues struct {
	baseMessage
	ParameterList map[string]*ValueStruct
	ParameterKey  string
}

type setParameterValuesBodyStruct struct {
	Body *setParameterValuesStruct `xml:"cwmp:SetParameterValues"`
}

type setParameterValuesStruct struct {
	ParamList    *ParameterListStruct `xml:"ParameterList"`
	ParameterKey string
}

type ParameterListStruct struct {
	Type            string                  `xml:"SOAP-ENC:arrayType,attr"`
	ParameterValues []*ParameterValueStruct `xml:"ParameterValueStruct"`
}

type ParameterValueStruct struct {
	Name  string       `xml:"Name"`
	Value *ValueStruct `xml:"Value"`
}

type ValueStruct struct {
	Type  string `xml:"xsi:type,attr"`
	Value string `xml:",chardata"`
}

func NewSetParameterValues() *SetParameterValues {
	return &SetParameterValues{
		ParameterList: make(map[string]*ValueStruct),
	}
}

func (msg *SetParameterValues) GetName() string {
	return "SetParameterValues"
}

func (msg *SetParameterValues) SetParameter(name string, v string, t string) {
	msg.ParameterList[name] = &ValueStruct{
		Type:  t,
		Value: v,
	}
}

func (msg *SetParameterValues) SetStringParameter(name string, v string) {
	msg.SetParameter(name, v, XsdString)
}

func (msg *SetParameterValues) SetUintParameter(name string, v string) {
	msg.SetParameter(name, v, XsdUnsignedint)
}

func (msg *SetParameterValues) CreateXML() []byte {
	body := &setParameterValuesBodyStruct{
		&setParameterValuesStruct{
			ParamList:    &ParameterListStruct{},
			ParameterKey: msg.ParameterKey,
		},
	}

	body.Body.ParamList.Type = "cwmp:ParameterValueStruct[" + strconv.Itoa(len(msg.ParameterList)) + "]"
	body.Body.ParamList.ParameterValues = make([]*ParameterValueStruct, 0)
	for k, v := range msg.ParameterList {
		value := &ParameterValueStruct{
			Name:  k,
			Value: v,
		}
		body.Body.ParamList.ParameterValues = append(body.Body.ParamList.ParameterValues, value)
	}
	return marshal(msg.GetID(), body)
}
