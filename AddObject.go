package cwmp

type AddObject struct {
	Header `xml:"-"`
	addObjectBody
}

type addObjectBody struct {
	ObjectName   string `xml:"cwmp:AddObject>ObjectName"`
	ParameterKey string `xml:"cwmp:AddObject>ParameterKey"`
}

func NewAddObject() *AddObject {
	msg := new(AddObject)
	msg.Header.RandomID()
	return msg
}

func (m *AddObject) Response() *AddObjectResponse {
	resp := new(AddObjectResponse)
	resp.ID = m.ID
	return resp
}
