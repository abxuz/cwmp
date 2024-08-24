package cwmp

type TransferCompleteResponse struct {
	baseMessage
	transferCompleteResponseStruct
}

type transferCompleteResponseBodyStruct struct {
	Body *transferCompleteResponseStruct `xml:"cwmp:TransferCompleteResponse"`
}

type transferCompleteResponseStruct struct {
}

func NewTransferCompleteResponse() *TransferCompleteResponse {
	return &TransferCompleteResponse{}
}

func (msg *TransferCompleteResponse) GetName() string {
	return "TransferCompleteResponse"
}

func (msg *TransferCompleteResponse) CreateXML() []byte {
	return marshal(msg.GetID(), &transferCompleteResponseBodyStruct{&msg.transferCompleteResponseStruct})
}
