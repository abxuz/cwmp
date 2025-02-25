package cwmp

const (
	FTFireware   string = "1 Firmware Upgrade Image"
	FTWebContent string = "2 Web Content"
	FTConfig     string = "3 Vendor Configuration File"
)

type Download struct {
	Header         `xml:"-"`
	CommandKey     string `xml:"cwmp:Download>CommandKey"`
	FileType       string `xml:"cwmp:Download>FileType"`
	URL            string `xml:"cwmp:Download>URL"`
	Username       string `xml:"cwmp:Download>Username"`
	Password       string `xml:"cwmp:Download>Password"`
	FileSize       int    `xml:"cwmp:Download>FileSize"`
	TargetFileName string `xml:"cwmp:Download>TargetFileName"`
	DelaySeconds   int    `xml:"cwmp:Download>DelaySeconds"`
	SuccessURL     string `xml:"cwmp:Download>SuccessURL"`
	FailureURL     string `xml:"cwmp:Download>FailureURL"`
}

func NewDownload() *Download {
	m := new(Download)
	m.Header.RandomID()
	return m
}

func (m *Download) Response() *DownloadResponse {
	resp := new(DownloadResponse)
	resp.ID = m.ID
	return resp
}
