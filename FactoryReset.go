package cwmp

type FactoryReset struct {
	baseMessage
	factoryResetStruct
}

type factoryResetBodyStruct struct {
	Body *factoryResetStruct `xml:"cwmp:FactoryReset"`
}

type factoryResetStruct struct {
}

func NewFactoryReset() *FactoryReset {
	return &FactoryReset{}
}

func (msg *FactoryReset) GetName() string {
	return "FactoryReset"
}

func (msg *FactoryReset) CreateXML() []byte {
	return marshal(msg.GetID(), &factoryResetBodyStruct{&msg.factoryResetStruct})
}
