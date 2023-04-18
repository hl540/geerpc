package codec

import (
	"io"
)

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(body interface{}) error
	Write(header *Header, body interface{}) error
}

type NewCodecFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  = "application/gob"
	JsonType = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
