package cwmp

import (
	"strconv"
	"strings"

	"github.com/abxuz/cwmp/xmlx"
)

type Inform struct {
	baseMessage
	Manufacturer  string
	OUI           string
	ProductClass  string
	SerialNumber  string
	Events        map[string]string
	MaxEnvelopes  int
	CurrentTime   string
	RetryCount    int
	CommandKey    string
	ParameterList map[string]string
}

func NewInform() *Inform {
	return &Inform{
		Events:        make(map[string]string),
		ParameterList: make(map[string]string),
	}
}

func (msg *Inform) GetName() string {
	return "Inform"
}

func (msg *Inform) HasEvent(event string) bool {
	_, ok := msg.Events[event]
	return ok
}

func (msg *Inform) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)

	deviceNode := doc.SelectNode("*", "DeviceId")
	if deviceNode != nil && len(strings.TrimSpace(deviceNode.String())) > 0 {
		msg.Manufacturer = xmlx.GetNodeValue(deviceNode, "", "Manufacturer")
		msg.OUI = xmlx.GetNodeValue(deviceNode, "", "OUI")
		msg.ProductClass = xmlx.GetNodeValue(deviceNode, "", "ProductClass")
		msg.SerialNumber = xmlx.GetNodeValue(deviceNode, "", "SerialNumber")
	}

	informNode := doc.SelectNode("*", "Inform")
	if informNode != nil && len(strings.TrimSpace(informNode.String())) > 0 {
		msg.CommandKey = xmlx.GetNodeValue(informNode, "", "CommandKey")
		msg.CurrentTime = xmlx.GetNodeValue(informNode, "", "CurrentTime")

		var err error
		msg.MaxEnvelopes, err = strconv.Atoi(xmlx.GetNodeValue(informNode, "", "MaxEnvelopes"))
		if err != nil {
			return err
		}

		msg.RetryCount, err = strconv.Atoi(xmlx.GetNodeValue(informNode, "", "RetryCount"))
		if err != nil {
			return err
		}
	}

	eventNode := doc.SelectNode("*", "Event")
	if eventNode != nil && len(strings.TrimSpace(eventNode.String())) > 0 {
		for _, event := range eventNode.Children {
			if event != nil && len(strings.TrimSpace(event.String())) > 0 {
				code := xmlx.GetNodeValue(event, "", "EventCode")
				msg.Events[code] = xmlx.GetNodeValue(event, "", "CommandKey")
			}
		}
	}

	paramsNode := doc.SelectNode("*", "ParameterList")
	if paramsNode != nil && len(strings.TrimSpace(paramsNode.String())) > 0 {
		var name string
		for _, param := range paramsNode.Children {
			if param != nil && len(strings.TrimSpace(param.String())) > 0 {
				name = xmlx.GetNodeValue(param, "", "Name")
				msg.ParameterList[name] = xmlx.GetNodeValue(param, "", "Value")
			}
		}
	}
	return nil
}
