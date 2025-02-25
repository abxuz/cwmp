package cwmp

type DeleteObject struct {
	Header       `xml:"-"`
	ObjectName   string `xml:"cwmp:DeleteObject>ObjectName"`
	ParameterKey string `xml:"cwmp:DeleteObject>ParameterKey"`
}

func NewDeleteObject() *DeleteObject {
	m := new(DeleteObject)
	m.Header.RandomID()
	return m
}

func (m *DeleteObject) Response() *DeleteObjectResponse {
	resp := new(DeleteObjectResponse)
	resp.ID = m.ID
	return resp
}
