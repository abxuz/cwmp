package cwmp

type AddObjectResponse struct {
	Header         `xml:"-"`
	InstanceNumber string
	Status         int
}
