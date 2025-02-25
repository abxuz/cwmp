package cwmp

type Fault struct {
	Header `xml:"-"`
	Detail string `xml:",innerxml"`
}

func NewFault() *Fault {
	m := new(Fault)
	m.Header.RandomID()
	return m
}
