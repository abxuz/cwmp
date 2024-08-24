package cwmp

import (
	"strings"

	"github.com/abxuz/cwmp/xmlx"
)

type GetRPCMethodsResponse struct {
	baseMessage
	Methods map[string]struct{}
}

func NewGetRPCMethodsResponse() *GetRPCMethodsResponse {
	return &GetRPCMethodsResponse{
		Methods: make(map[string]struct{}),
	}
}

func (msg *GetRPCMethodsResponse) GetName() string {
	return "GetRPCMethodsResponse"
}

func (msg *GetRPCMethodsResponse) Has(name string) bool {
	_, ok := msg.Methods[name]
	return ok
}

func (msg *GetRPCMethodsResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	methodList := doc.SelectNode("*", "MethodList")
	if len(strings.TrimSpace(methodList.String())) > 0 {
		for _, param := range methodList.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name := param.GetValue()
				msg.Methods[name] = struct{}{}
			}
		}

	}
	return nil
}
