package cwmp

import (
	"strconv"

	"github.com/abxuz/cwmp/xmlx"
)

type DownloadResponse struct {
	baseMessage
	Status       int
	StartTime    string
	CompleteTime string
}

func NewDownloadResponse() *DownloadResponse {
	return &DownloadResponse{}
}

func (msg *DownloadResponse) GetName() string {
	return "DownloadResponse"
}

func (msg *DownloadResponse) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	statusNode := doc.SelectNode("*", "Status")
	if statusNode != nil {
		var err error
		msg.Status, err = strconv.Atoi(statusNode.GetValue())
		if err != nil {
			return err
		}
	}
	msg.StartTime = doc.SelectNode("*", "StartTime").GetValue()
	msg.CompleteTime = doc.SelectNode("*", "CompleteTime").GetValue()
	return nil
}
