package id

import (
	"github.com/speps/go-hashids/v2"
)

type ID struct{}

func NewID() *ID {
	return &ID{}
}

var Salt = "oBT6IsLAyrY7"

func (i ID) Encoder(id int) string {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = 11
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{id})
	return e

}

func (i ID) Decoder(str string) int {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = 11
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeWithError(str)
	return d[0]
}
