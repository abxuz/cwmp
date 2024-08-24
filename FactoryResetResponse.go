package cwmp

type FactoryResetResponse struct {
	baseMessage
}

func NewFactoryResetResponse() *FactoryResetResponse {
	return &FactoryResetResponse{}
}

func (msg *FactoryResetResponse) GetName() string {
	return "FactoryResetResponse"
}
