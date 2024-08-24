package cwmp

import (
	"strings"

	"github.com/abxuz/cwmp/xmlx"
)

type GetParameterNamesResponse struct {
	baseMessage
	ParameterList map[string]bool
}

func NewGetParameterNamesResponse() *GetParameterNamesResponse {
	return &GetParameterNamesResponse{
		ParameterList: make(map[string]bool),
	}
}

func (msg *GetParameterNamesResponse) GetName() string {
	return "GetParameterNamesResponse"
}

func (msg *GetParameterNamesResponse) Has(name string) bool {
	_, ok := msg.ParameterList[name]
	return ok
}

func (msg *GetParameterNamesResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	paramList := doc.SelectNode("*", "ParameterList")
	if len(strings.TrimSpace(paramList.String())) > 0 {
		for _, param := range paramList.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name := param.SelectNode("", "Name").GetValue()
				writable := strings.ToLower(param.SelectNode("", "Writable").GetValue())
				msg.ParameterList[name] = writable == "1" || writable == "true"
			}
		}
	}
	return nil
}
