package server

import "io"

type EndpointResult interface {
	Status() int
	Write(writer io.Writer) error
	io.Closer
}

type Builder struct {
}

func (b *Builder) Endpoint() *Builder {
	return b
}
