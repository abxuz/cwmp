package cwmp

type ScheduleInform struct {
	Header       `xml:"-"`
	CommandKey   string `xml:"cwmp:ScheduleInform>CommandKey"`
	DelaySeconds int    `xml:"cwmp:ScheduleInform>DelaySeconds"`
}

func NewScheduleInform() *ScheduleInform {
	m := new(ScheduleInform)
	m.Header.RandomID()
	return m
}

func (m *ScheduleInform) Response() *ScheduleInformResponse {
	resp := new(ScheduleInformResponse)
	resp.ID = m.ID
	return resp
}
