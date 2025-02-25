package cwmp

import (
	"math/rand"
	"strconv"
)

type Message interface {
	GetID() string
	SetID(id string)
}

type Header struct {
	ID string
}

func (h *Header) RandomID() {
	h.SetID(strconv.FormatUint(rand.Uint64(), 10))
}

func (h *Header) SetID(id string) {
	h.ID = id
}

func (h *Header) GetID() string {
	return h.ID
}
