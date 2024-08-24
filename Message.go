package cwmp

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/abxuz/cwmp/xmlx"
)

type Message interface {
	GetID() string
	GetName() string
	CreateXML() []byte
}

type baseMessage struct {
	ID string
}

func (msg *baseMessage) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = "ID:" + strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return msg.ID
}

func (msg *baseMessage) Parse(doc *xmlx.Document) error {
	msg.ID = xmlx.GetDocNodeValue(doc, "*", "ID")
	return nil
}
func (msg *baseMessage) CreateXML() []byte { panic("not implement") }

type envelope struct {
	XMLName   xml.Name `xml:"SOAP-ENV:Envelope"`
	XmlnsEnv  string   `xml:"xmlns:SOAP-ENV,attr"`
	XmlnsEnc  string   `xml:"xmlns:SOAP-ENC,attr"`
	XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
	XmlnsCwmp string   `xml:"xmlns:cwmp,attr"`
	Header    struct {
		ID struct {
			MustUnderstand string `xml:"SOAP-ENV:mustUnderstand,attr,omitempty"`
			Value          string `xml:",chardata"`
		} `xml:"cwmp:ID"`
	} `xml:"SOAP-ENV:Header"`
	Body any `xml:"SOAP-ENV:Body"`
}

func marshal(id string, body any) []byte {
	env := &envelope{
		XmlnsEnv:  "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsEnc:  "http://schemas.xmlsoap.org/soap/encoding/",
		XmlnsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsXsd:  "http://www.w3.org/2001/XMLSchema",
		XmlnsCwmp: "urn:dslforum-org:cwmp-1-0",
		Body:      body,
	}
	env.Header.ID.MustUnderstand = "1"
	env.Header.ID.Value = id
	data, _ := xml.MarshalIndent(env, "  ", "    ")
	return data
}
