package cwmp

type Inform struct {
	Header        `xml:"-"`
	Manufacturer  string       `xml:"DeviceId>Manufacturer"`
	OUI           string       `xml:"DeviceId>OUI"`
	ProductClass  string       `xml:"DeviceId>ProductClass"`
	SerialNumber  string       `xml:"DeviceId>SerialNumber"`
	Events        []*Event     `xml:"Event>EventStruct"`
	ParameterList []*Parameter `xml:"ParameterList>ParameterValueStruct"`
}

type Event struct {
	EventCode  string
	CommandKey string
}

type Parameter struct {
	Name  string
	Value string
}

func NewInform() *Inform {
	m := new(Inform)
	m.Header.RandomID()
	return m
}

func (m *Inform) Response() *InformResponse {
	resp := new(InformResponse)
	resp.ID = m.ID
	resp.MaxEnvelopes = 1
	return resp
}
