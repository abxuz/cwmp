package cwmp

type GetParameterNames struct {
	baseMessage
	getParameterNamesStruct
}

type getParameterNamesBodyStruct struct {
	Body *getParameterNamesStruct `xml:"cwmp:GetParameterNames"`
}

type getParameterNamesStruct struct {
	ParameterPath string `xml:"ParameterPath"`
	NextLevel     bool   `xml:"NextLevel"`
}

func NewGetParameterNames() *GetParameterNames {
	return &GetParameterNames{}
}

func (msg *GetParameterNames) GetName() string {
	return "GetParameterNames"
}

func (msg *GetParameterNames) CreateXML() []byte {
	return marshal(msg.GetID(), &getParameterNamesBodyStruct{&msg.getParameterNamesStruct})
}
