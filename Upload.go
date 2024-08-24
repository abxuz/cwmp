package cwmp

type Upload struct {
	baseMessage
	uploadStruct
}

type uploadBodyStruct struct {
	Body *uploadStruct `xml:"cwmp:Upload"`
}

type uploadStruct struct {
	CommandKey   string
	FileType     string
	URL          string
	Username     string
	Password     string
	DelaySeconds int
}

func NewUpload() *Upload {
	return &Upload{}
}

func (msg *Upload) GetName() string {
	return "Upload"
}

func (msg *Upload) CreateXML() []byte {
	return marshal(msg.GetID(), &uploadBodyStruct{&msg.uploadStruct})
}
