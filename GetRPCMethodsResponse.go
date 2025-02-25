package cwmp

type GetRPCMethodsResponse struct {
	Header  `xml:"-"`
	Methods []string `xml:"MethodList>string"`
}
