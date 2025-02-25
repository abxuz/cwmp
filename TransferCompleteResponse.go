package cwmp

type TransferCompleteResponse struct {
	Header `xml:"-"`
	Body   struct{} `xml:"cwmp:TransferCompleteResponse"`
}
