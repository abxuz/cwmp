package cwmp

type Upload struct {
	Header       `xml:"-"`
	CommandKey   string `xml:"cwmp:Upload>CommandKey"`
	FileType     string `xml:"cwmp:Upload>FileType"`
	URL          string `xml:"cwmp:Upload>URL"`
	Username     string `xml:"cwmp:Upload>Username"`
	Password     string `xml:"cwmp:Upload>Password"`
	DelaySeconds int    `xml:"cwmp:Upload>DelaySeconds"`
}

func NewUpload() *Upload {
	m := new(Upload)
	m.Header.RandomID()
	return m
}

func (m *Upload) Response() *UploadResponse {
	resp := new(UploadResponse)
	resp.ID = m.ID
	return resp
}
