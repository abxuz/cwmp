package cwmp

import (
	"encoding/xml"
	"io"
)

type EnvelopeDecoder struct {
	ID   string      `xml:"Header>ID"`
	Body BodyDecoder `xml:"Body"`
}

type BodyDecoder struct {
	Message
}

func (b *BodyDecoder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	start, ok := t.(xml.StartElement)
	if !ok {
		return ErrInvalidCWMPXML
	}

	switch start.Name.Local {
	case "Inform":
		b.Message = new(Inform)
	case "GetParameterValuesResponse":
		b.Message = new(GetParameterValuesResponse)
	case "SetParameterValuesResponse":
		b.Message = new(SetParameterValuesResponse)
	case "AddObjectResponse":
		b.Message = new(AddObjectResponse)
	case "DeleteObjectResponse":
		b.Message = new(DeleteObjectResponse)
	case "RebootResponse":
		b.Message = new(RebootResponse)
	case "FactoryResetResponse":
		b.Message = new(FactoryResetResponse)
	case "Fault":
		b.Message = new(Fault)
	case "ScheduleInformResponse":
		b.Message = new(ScheduleInformResponse)
	case "GetRPCMethodsResponse":
		b.Message = new(GetRPCMethodsResponse)
	case "GetParameterNamesResponse":
		b.Message = new(GetParameterNamesResponse)
	case "UploadResponse":
		b.Message = new(UploadResponse)
	case "DownloadResponse":
		b.Message = new(DownloadResponse)
	case "TransferCompleteResponse":
		b.Message = new(TransferCompleteResponse)
	default:
		return ErrInvalidCWMPXML
	}

	err = d.DecodeElement(b.Message, &start)
	if err != nil {
		return err
	}

	t, err = d.Token()
	if err != nil {
		return err
	}
	end, ok := t.(xml.EndElement)
	if !ok || end.Name.Local != "Body" {
		return ErrInvalidCWMPXML
	}
	return nil
}

func DecodeFrom(r io.Reader) (Message, error) {
	decoder := xml.NewDecoder(r)

	enveloper := new(EnvelopeDecoder)
	err := decoder.Decode(enveloper)
	if err != nil {
		return nil, err
	}

	enveloper.Body.Message.SetID(enveloper.ID)
	return enveloper.Body.Message, nil
}
