package cwmp

type GetParameterNames struct {
	Header        `xml:"-"`
	ParameterPath string `xml:"cwmp:GetParameterNames>ParameterPath"`
	NextLevel     bool   `xml:"cwmp:GetParameterNames>NextLevel"`
}

func NewGetParameterNames() *GetParameterNames {
	m := new(GetParameterNames)
	m.Header.RandomID()
	return m
}

func (m *GetParameterNames) Response() *GetParameterNamesResponse {
	resp := new(GetParameterNamesResponse)
	resp.ID = m.ID
	return resp
}
