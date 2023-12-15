// Package encoding handles the encoding process for CSV, TSV, and fixed length text.
package encoding

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"reflect"

	"github.com/ianlopshire/go-fixedwidth"
	"github.com/mpyw/go-snippets/stream/charset"
	"github.com/mpyw/go-snippets/stream/newline"
)

var (
	_ io.Writer    = (*flushBlocker)(nil)
	_ Encoder[any] = (*fixedWidthWriter[any])(nil)
)

type FixedWidthEncoderOptions struct {
	// Charset is an output character code (default: charset.UTF8)
	Charset charset.Charset
	// Newline default: LF
	Newline newline.Newline
}

// NewFixedWidthWriter creates an Encoder that generates a fixed length field.
// See csvutil documentation for structure tags: https://github.com/ianlopshire/go-fixedwidth
func NewFixedWidthWriter[TRow any](dst io.Writer, opts FixedWidthEncoderOptions) Encoder[TRow] {
	// fixedwidth library flushes every record, which is inefficient, so add work to make the original Flush() do nothing
	bl := &flushBlocker{bufio.NewWriter(dst)}
	enc := fixedwidth.NewEncoder(bl)
	return &fixedWidthWriter[TRow]{bl, enc, opts}
}

type fixedWidthWriter[TRow any] struct {
	bl    *flushBlocker
	fixed *fixedwidth.Encoder
	opts  FixedWidthEncoderOptions
}

func (enc *fixedWidthWriter[TRow]) Encode(row TRow) error {
	// Change the character encoding of the string field and write it to the buffer.
	// Must be taken from a pointer to make it writable.
	val := reflect.Indirect(reflect.ValueOf(&row))
	if val.Kind() != reflect.Struct {
		return errors.New("FixedWidthWriter: unsupported type")
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() != reflect.String {
			continue
		}
		buf := bytes.Buffer{}
		if _, err := io.WriteString(charset.NewEncoder(&buf, enc.opts.Charset), field.String()); err != nil {
			return err
		}
		field.SetString(buf.String())
	}
	// Write the contents of the structure
	if err := enc.fixed.Encode(row); err != nil {
		return err
	}
	// Add a newline on your own, since writing a single line does not break the line.
	_, err := enc.bl.WriteString(enc.opts.Newline.ToString())
	return err
}

func (enc *fixedWidthWriter[TRow]) Flush() error {
	return enc.bl.reallyFlush()
}

type flushBlocker struct {
	*bufio.Writer
}

func (b *flushBlocker) Flush() error {
	// go-fixedwidth library flushes every record, which is inefficient, so add work to make the original Flush() do nothing
	return nil
}

func (b *flushBlocker) reallyFlush() error {
	return b.Writer.Flush()
}
