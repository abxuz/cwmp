package cwmp

import (
	"encoding/xml"
	"io"
)

type EnvelopeEncoder struct {
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
	Body Message `xml:"SOAP-ENV:Body"`
}

func EncodeTo(body Message, w io.Writer) error {
	env := new(EnvelopeEncoder)
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header.ID.MustUnderstand = "1"
	env.Header.ID.Value = body.GetID()
	env.Body = body
	encoder := xml.NewEncoder(w)
	err := encoder.Encode(env)
	if err != nil {
		return err
	}
	return encoder.Flush()
}
