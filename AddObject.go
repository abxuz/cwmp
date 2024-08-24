package cwmp

type AddObject struct {
	baseMessage
	addObjectStruct
}

type addObjectBodyStruct struct {
	Body *addObjectStruct `xml:"cwmp:AddObject"`
}

type addObjectStruct struct {
	ObjectName   string
	ParameterKey string
}

func NewAddObject() *AddObject {
	return &AddObject{}
}

func (msg *AddObject) GetName() string {
	return "AddObject"
}

func (msg *AddObject) CreateXML() []byte {
	return marshal(msg.GetID(), &addObjectBodyStruct{&msg.addObjectStruct})
}
