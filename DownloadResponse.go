package cwmp

type DownloadResponse struct {
	Header       `xml:"-"`
	Status       int
	StartTime    string
	CompleteTime string
}
