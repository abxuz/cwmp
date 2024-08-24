package cwmp

type InformResponse struct {
	baseMessage
	informResponseStruct
}

type informResponseBodyStruct struct {
	Body *informResponseStruct `xml:"cwmp:InformResponse"`
}

type informResponseStruct struct {
	MaxEnvelopes int `xml:"MaxEnvelopes"`
}

func NewInformResponse() *InformResponse {
	return &InformResponse{}
}

func (msg *InformResponse) GetName() string {
	return "InformResponse"
}

func (msg *InformResponse) CreateXML() []byte {
	return marshal(msg.GetID(), &informResponseBodyStruct{&msg.informResponseStruct})
}
