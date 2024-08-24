package cwmp

import (
	"strconv"

	"github.com/abxuz/cwmp/xmlx"
)

type AddObjectResponse struct {
	baseMessage
	InstanceNumber string
	Status         int
}

func NewAddObjectResponse() *AddObjectResponse {
	return &AddObjectResponse{}
}

func (msg *AddObjectResponse) GetName() string {
	return "AddObjectResponse"
}

func (msg *AddObjectResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	msg.InstanceNumber = doc.SelectNode("*", "InstanceNumber").GetValue()
	statusNode := doc.SelectNode("*", "Status")
	if statusNode != nil {
		var err error
		msg.Status, err = strconv.Atoi(statusNode.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}
