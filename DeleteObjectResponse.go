package cwmp

import (
	"strconv"

	"github.com/abxuz/cwmp/xmlx"
)

type DeleteObjectResponse struct {
	baseMessage
	Status int
}

func NewDeleteObjectResponse() *DeleteObjectResponse {
	return &DeleteObjectResponse{}
}

func (msg *DeleteObjectResponse) GetName() string {
	return "DeleteObjectResponse"
}

func (msg *DeleteObjectResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
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
