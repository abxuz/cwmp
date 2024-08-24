package cwmp

import (
	"strconv"

	"github.com/abxuz/cwmp/xmlx"
)

type SetParameterValuesResponse struct {
	baseMessage
	Status       int
	ParameterKey string
}

func NewSetParameterValuesResponse() *SetParameterValuesResponse {
	return &SetParameterValuesResponse{}
}

func (msg *SetParameterValuesResponse) GetName() string {
	return "SetParameterValuesResponse"
}

func (msg *SetParameterValuesResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	statusNode := doc.SelectNode("*", "Status")
	if statusNode != nil {
		var err error
		msg.Status, err = strconv.Atoi(statusNode.GetValue())
		if err != nil {
			return err
		}
	}

	paramsNode := doc.SelectNode("*", "ParameterKey")
	if paramsNode != nil {
		msg.ParameterKey = paramsNode.GetValue()
	}
	return nil
}
