package cwmp

type GetParameterValuesResponse struct {
	Header        `xml:"-"`
	ParameterList []*Parameter `xml:"ParameterList>ParameterValueStruct"`
}

func (m *GetParameterValuesResponse) Values() map[string]string {
	values := make(map[string]string)
	for _, p := range m.ParameterList {
		values[p.Name] = p.Value
	}
	return values
}
