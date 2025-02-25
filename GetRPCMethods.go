package cwmp

type GetRPCMethods struct {
	Header `xml:"-"`
	Body   struct{} `xml:"cwmp:GetRPCMethods"`
}

func NewGetRPCMethods() *GetRPCMethods {
	m := new(GetRPCMethods)
	m.Header.RandomID()
	return m
}

func (m *GetRPCMethods) Response() *GetRPCMethodsResponse {
	resp := new(GetRPCMethodsResponse)
	resp.ID = m.ID
	return resp
}
