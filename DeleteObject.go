package cwmp

type DeleteObject struct {
	baseMessage
	deleteObjectStruct
}

type deleteObjectBodyStruct struct {
	Body *deleteObjectStruct `xml:"cwmp:DeleteObject"`
}

type deleteObjectStruct struct {
	ObjectName   string
	ParameterKey string
}

func NewDeleteObject() *DeleteObject {
	return &DeleteObject{}
}

func (msg *DeleteObject) GetName() string {
	return "DeleteObject"
}

func (msg *DeleteObject) CreateXML() []byte {
	return marshal(msg.GetID(), &deleteObjectBodyStruct{&msg.deleteObjectStruct})
}
