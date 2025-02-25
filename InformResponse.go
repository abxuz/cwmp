package cwmp

type InformResponse struct {
	Header       `xml:"-"`
	MaxEnvelopes int `xml:"cwmp:InformResponse>MaxEnvelopes"`
}
