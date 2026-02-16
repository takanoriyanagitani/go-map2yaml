package map2yaml

import (
	"context"
	"io"
)

type InputMap map[string]any

type Writer interface {
	Write(ctx context.Context, imap InputMap) error
}

type WriterFn func(context.Context, InputMap) error

func (f WriterFn) Write(ctx context.Context, imap InputMap) error {
	return f(ctx, imap)
}

func (f WriterFn) AsWriter() Writer { return f }

type WriteCloser interface {
	Writer
	io.Closer
}
