package cwmp

type ScheduleInform struct {
	baseMessage
	scheduleInformStruct
}

type scheduleInformBodyStruct struct {
	Body *scheduleInformStruct `xml:"cwmp:ScheduleInform"`
}

type scheduleInformStruct struct {
	CommandKey   string
	DelaySeconds int
}

func NewScheduleInform() *ScheduleInform {
	return &ScheduleInform{}
}

func (msg *ScheduleInform) GetName() string {
	return "ScheduleInform"
}

func (msg *ScheduleInform) CreateXML() []byte {
	return marshal(msg.GetID(), &scheduleInformBodyStruct{&msg.scheduleInformStruct})
}
