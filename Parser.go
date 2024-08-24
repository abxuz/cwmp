package cwmp

import (
	"errors"

	"github.com/abxuz/cwmp/xmlx"
)

type parser interface {
	Parse(doc *xmlx.Document) error
}

func ParseXML(data []byte) (Message, error) {
	if len(data) == 0 {
		return nil, nil
	}

	doc := xmlx.New()
	err := doc.LoadBytes(data, nil)
	if err != nil {
		return nil, err
	}
	bodyNode := doc.SelectNode("*", "Body")
	if bodyNode == nil {
		return nil, errors.New("body missing")
	}

	var name string
	if len(bodyNode.Children) > 1 {
		name = bodyNode.Children[1].Name.Local
	} else {
		name = bodyNode.Children[0].Name.Local
	}

	var msg Message
	switch name {
	case "Inform":
		msg = NewInform()
	case "GetRPCMethodsResponse":
		msg = NewGetRPCMethodsResponse()
	case "GetParameterNamesResponse":
		msg = NewGetParameterNamesResponse()
	case "GetParameterValuesResponse":
		msg = NewGetParameterValuesResponse()
	case "SetParameterValuesResponse":
		msg = NewSetParameterValuesResponse()
	case "AddObjectResponse":
		msg = NewAddObjectResponse()
	case "DeleteObjectResponse":
		msg = NewDeleteObjectResponse()
	case "RebootResponse":
		msg = NewRebootResponse()
	case "FactoryResetResponse":
		msg = NewFactoryResetResponse()
	case "UploadResponse":
		msg = NewUploadResponse()
	case "DownloadResponse":
		msg = NewDownloadResponse()
	case "TransferComplete":
		msg = NewTransferComplete()
	case "ScheduleInformResponse":
		msg = NewScheduleInformResponse()
	case "Fault":
		msg = NewFault()
	default:
		return nil, errors.New("unsupported msg type: " + name)
	}

	parser, ok := msg.(parser)
	if !ok {
		return nil, errors.New("unsupported msg type")
	}
	err = parser.Parse(doc)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
