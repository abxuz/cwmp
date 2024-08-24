package cwmp

import (
	"strconv"

	"github.com/abxuz/cwmp/xmlx"
)

type TransferComplete struct {
	baseMessage
	CommandKey   string
	StartTime    string
	CompleteTime string
	Fault        *FaultStruct
}

type FaultStruct struct {
	FaultCode   int
	FaultString string
}

func NewTransferComplete() *TransferComplete {
	return &TransferComplete{
		Fault: &FaultStruct{},
	}
}

func (msg *TransferComplete) GetName() string {
	return "TransferComplete"
}

func (msg *TransferComplete) Parse(doc *xmlx.Document) error {
	msg.baseMessage.Parse(doc)
	msg.CommandKey = doc.SelectNode("*", "CommandKey").GetValue()
	msg.CompleteTime = doc.SelectNode("*", "CompleteTime").GetValue()
	msg.StartTime = doc.SelectNode("*", "StartTime").GetValue()
	msg.Fault.FaultString = doc.SelectNode("*", "FaultString").GetValue()
	faultCode, err := strconv.Atoi(doc.SelectNode("*", "FaultCode").GetValue())
	if err != nil {
		return err
	}
	msg.Fault.FaultCode = faultCode
	return nil
}
