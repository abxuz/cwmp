package cwmp

type FactoryReset struct {
	Header `xml:"-"`
	Body   struct{} `xml:"cwmp:FactoryReset"`
}

func NewFactoryReset() *FactoryReset {
	m := new(FactoryReset)
	m.Header.RandomID()
	return m
}

func (m *FactoryReset) Response() *FactoryResetResponse {
	resp := new(FactoryResetResponse)
	resp.ID = m.ID
	return resp
}
