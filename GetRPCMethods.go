package cwmp

type GetRPCMethods struct {
	baseMessage
	getRPCMethodsStruct
}

type getRPCMethodsBodyStruct struct {
	Body *getRPCMethodsStruct `xml:"cwmp:GetRPCMethods"`
}

type getRPCMethodsStruct struct {
}

func NewGetRPCMethods() *GetRPCMethods {
	return &GetRPCMethods{}
}

func (msg *GetRPCMethods) GetName() string {
	return "GetRPCMethods"
}

func (msg *GetRPCMethods) CreateXML() []byte {
	return marshal(msg.GetID(), &getRPCMethodsBodyStruct{&msg.getRPCMethodsStruct})
}
