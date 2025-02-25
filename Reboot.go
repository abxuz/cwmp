package cwmp

type Reboot struct {
	Header     `xml:"-"`
	CommandKey string `xml:"cwmp:Reboot>CommandKey"`
}

func NewReboot() *Reboot {
	m := new(Reboot)
	m.Header.RandomID()
	return m
}

func (m *Reboot) Response() *RebootResponse {
	resp := new(RebootResponse)
	resp.ID = m.ID
	return resp
}
