package cwmp

type RebootResponse struct {
	baseMessage
}

func NewRebootResponse() *RebootResponse {
	return &RebootResponse{}
}

func (msg *RebootResponse) GetName() string {
	return "RebootResponse"
}
