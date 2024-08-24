package cwmp

import (
	"github.com/abxuz/cwmp/xmlx"
)

type Fault struct {
	baseMessage
	Detail string
}

func NewFault() *Fault {
	return &Fault{}
}

func (msg *Fault) GetName() string {
	return "Fault"
}

func (msg *Fault) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	msg.Detail = doc.SelectNode("*", "detail").String()
	return nil
}
