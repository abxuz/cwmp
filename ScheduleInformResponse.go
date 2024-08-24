package cwmp

type ScheduleInformResponse struct {
	baseMessage
}

func NewScheduleInformResponse() *ScheduleInformResponse {
	return &ScheduleInformResponse{}
}

func (msg *ScheduleInformResponse) GetName() string {
	return "ScheduleInformResponse"
}
