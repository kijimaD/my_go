package gen

import (
	"context"
)

type Writer interface {
	Write(id string) error
}

type WriterFactory func(ctx context.Context) (Writer, bool, context.Context)

func oneOffProgress(ctx context.Context, id string) func(err error) error {
	pw, _, _ := NewFromContext(ctx)
	_ = pw.Write(id)

	return func(err error) error {
		// TODO: set error on status
		pw.Write(id)
		return err
	}
}

func NewFromContext(ctx context.Context) (Writer, bool, context.Context) {
	return FromContext(ctx)(ctx)
}

func FromContext(ctx context.Context) WriterFactory {
	v := ctx.Value(ctxkey)
	return func(ctx context.Context) (Writer, bool, context.Context) {
		pw, ok := v.(*progressWriter)
		if ok != true {
			panic("pwでない")
		}
		pw = newWriter(pw)
		ctx = context.WithValue(ctx, ctxkey, pw)
		return pw, false, ctx
	}
}

type progressWriter struct {
	reader string
}

func newWriter(pw *progressWriter) *progressWriter {
	pw = &progressWriter{
		reader: pw.reader,
	}
	pw.reader = "123"
	return pw
}

func (pw *progressWriter) Write(id string) error {
	pw.reader = id
	return nil
}
