package cwmp

const (
	FTFireware   string = "1 Firmware Upgrade Image"
	FTWebContent string = "2 Web Content"
	FTConfig     string = "3 Vendor Configuration File"
)

type Download struct {
	baseMessage
	downloadStruct
}

type downloadBodyStruct struct {
	Body *downloadStruct `xml:"cwmp:Download"`
}

type downloadStruct struct {
	CommandKey     string
	FileType       string
	URL            string
	Username       string
	Password       string
	FileSize       int
	TargetFileName string
	DelaySeconds   int
	SuccessURL     string
	FailureURL     string
}

func NewDownload() *Download {
	return &Download{}
}

func (msg *Download) GetName() string {
	return "Download"
}

func (msg *Download) CreateXML() []byte {
	return marshal(msg.GetID(), &downloadBodyStruct{&msg.downloadStruct})
}
