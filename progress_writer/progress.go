package progress

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type contextKeyT string

var contextKey = contextKeyT("progress")

type WriterFactory func(ctx context.Context) (Writer, bool, context.Context)

func FromContext(ctx context.Context, opts ...WriterOption) WriterFactory {
	v := ctx.Value(contextKey)
	return func(ctx context.Context) (Writer, bool, context.Context) {
		pw, ok := v.(*progressWriter)
		if !ok {
			panic("これはprogressWriterではない")
		}
		pw = newWriter(pw)
		for _, o := range opts {
			o(pw)
		}
		ctx = context.WithValue(ctx, contextKey, pw)
		return pw, true, ctx
	}
}

type Writer interface {
	Write(id string, value interface{}) error
	Close() error
}

type Progress struct {
	ID        string
	Timestamp time.Time
	Sys       interface{}
	meta      map[string]interface{}
}

type WriterOption func(Writer)

func newWriter(pw *progressWriter) *progressWriter {
	meta := make(map[string]interface{})
	for k, v := range pw.meta {
		meta[k] = v
	}
	pw = &progressWriter{
		reader: pw.reader,
		meta:   meta,
	}
	pw.reader.append(pw)
	return pw
}

type progressWriter struct {
	done   bool
	reader *progressReader
	meta   map[string]interface{}
}

func (pw *progressWriter) Write(id string, v interface{}) error {
	if pw.done {
		return errors.Errorf("writing %s to closed progress writer", id)
	}
	return pw.writeRawProgress(&Progress{
		ID:        id,
		Timestamp: time.Now(),
		Sys:       v,
		meta:      pw.meta,
	})
}

func (pw *progressWriter) WriteRawProgress(p *Progress) error {
	meta := p.meta
	if len(pw.meta) > 0 {
		meta = map[string]interface{}{}
		for k, v := range p.meta {
			meta[k] = v
		}
		for k, v := range pw.meta {
			if _, ok := meta[k]; !ok {
				meta[k] = v
			}
		}
	}
	p.meta = meta
	return pw.writeRawProgress(p)
}

func (pw *progressWriter) writeRawProgress(p *Progress) error {
	pw.reader.mu.Lock()
	pw.reader.dirty[p.ID] = p
	pw.reader.cond.Broadcast()
	pw.reader.mu.Unlock()
	return nil
}

func (pw *progressWriter) Close() error {
	pw.reader.mu.Lock()
	delete(pw.reader.writers, pw)
	pw.reader.mu.Unlock()
	pw.reader.cond.Broadcast()
	pw.done = true
	return nil
}

type progressReader struct {
	ctx     context.Context
	cond    *sync.Cond
	mu      sync.Mutex
	writers map[*progressWriter]struct{}
	dirty   map[string]*Progress
}

func (pr *progressReader) append(pw *progressWriter) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	select {
	case <-pr.ctx.Done():
		return
	default:
		pr.writers[pw] = struct{}{}
	}
}
