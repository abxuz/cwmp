package cwmp

type UploadResponse struct {
	Header       `xml:"-"`
	Status       int
	StartTime    string
	CompleteTime string
}
