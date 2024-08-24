package cwmp

type Reboot struct {
	baseMessage
	rebootStruct
}

type rebootBodyStruct struct {
	Body *rebootStruct `xml:"cwmp:Reboot"`
}

type rebootStruct struct {
	CommandKey string
}

func NewReboot() *Reboot {
	return &Reboot{}
}

func (msg *Reboot) GetName() string {
	return "Reboot"
}

func (msg *Reboot) CreateXML() []byte {
	return marshal(msg.GetID(), &rebootBodyStruct{&msg.rebootStruct})
}
