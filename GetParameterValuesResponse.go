package cwmp

import (
	"strings"

	"github.com/abxuz/cwmp/xmlx"
)

type GetParameterValuesResponse struct {
	baseMessage
	Values map[string]string
}

func NewGetParameterValuesResponse() *GetParameterValuesResponse {
	return &GetParameterValuesResponse{
		Values: make(map[string]string),
	}
}

func (msg *GetParameterValuesResponse) GetName() string {
	return "GetParameterValuesResponse"
}

func (msg *GetParameterValuesResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	paramsNode := doc.SelectNode("*", "ParameterList")
	if len(strings.TrimSpace(paramsNode.String())) > 0 {
		for _, param := range paramsNode.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name := param.SelectNode("", "Name").GetValue()
				value := param.SelectNode("", "Value").GetValue()
				msg.Values[name] = value
			}
		}
	}
	return nil
}
