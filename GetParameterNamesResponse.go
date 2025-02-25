package cwmp

type GetParameterNamesResponse struct {
	Header        `xml:"-"`
	ParameterList []*ParameterInfo `xml:"ParameterList>ParameterInfoStruct"`
}

type ParameterInfo struct {
	Name     string
	Writable bool
}
