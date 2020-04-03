package errors

import (
	"bytes"
	"io"
)

type CouldNotSaveOplogOnElasticsearch struct {
	Message io.ReadCloser
}

type CouldNotReadLastTimeOnElasticsearch struct {
	Message io.ReadCloser
}

func (e *CouldNotSaveOplogOnElasticsearch) Error() string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(e.Message)
	return buf.String()
}

func (e *CouldNotReadLastTimeOnElasticsearch) Error() string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(e.Message)
	return buf.String()
}
