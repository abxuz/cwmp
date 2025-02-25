package cwmp

type TransferComplete struct {
	Header       `xml:"-"`
	CommandKey   string
	StartTime    string
	CompleteTime string
	Fault        *FaultStruct
}

type FaultStruct struct {
	FaultCode   int
	FaultString string
}

func NewTransferComplete() *TransferComplete {
	m := new(TransferComplete)
	m.Header.RandomID()
	return m
}

func (m *TransferComplete) Response() *TransferCompleteResponse {
	resp := new(TransferCompleteResponse)
	resp.ID = m.ID
	return resp
}
