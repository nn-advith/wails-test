package codec

import (
	b64 "encoding/base64"
)

type Codec struct {
}

func NewCodec() *Codec {
	return &Codec{}
}

func (c *Codec) Encode(msg string) string {
	return b64.StdEncoding.EncodeToString([]byte(msg))
}

func (c *Codec) Decode(msg string) string {
	dec, _ := b64.StdEncoding.DecodeString(msg)
	return string(dec)
}
