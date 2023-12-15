// Package decoding handles decoding of CSV, TSV, JSON, and fixed length text.
package decoding

import (
	"errors"
	"io"
	"reflect"
	"strings"

	"github.com/ianlopshire/go-fixedwidth"
	"github.com/mpyw/go-snippets/stream/charset"
)

var (
	_ Decoder[any] = (*fixedWidthDecoder[any])(nil)
)

type FixedWidthDecoderOptions struct {
	// Charset is an input character encoding (default: charset.UTF8)
	Charset charset.Charset
}

// NewFixedWidthDecoder creates a Decoder that processes fixed length fields.
// See the fixedwidth documentation for the structure tags: https://github.com/ianlopshire/go-fixedwidth
func NewFixedWidthDecoder[TRow any](src io.Reader, opts FixedWidthDecoderOptions) Decoder[TRow] {
	return &fixedWidthDecoder[TRow]{fixedwidth.NewDecoder(src), opts}
}

type fixedWidthDecoder[TRow any] struct {
	fixed *fixedwidth.Decoder
	opts  FixedWidthDecoderOptions
}

func (dec *fixedWidthDecoder[TRow]) Decode() (TRow, error) {
	var row TRow
	if err := dec.fixed.Decode(&row); err != nil {
		return row, err
	}
	// Change the character encoding of the string field and trim it to include full-width spaces.
	// Since trimming full-width spaces with strings.TrimSpace() only works after character encoding conversion,
	// go-fixedwidth functionality alone cannot solve this problem.
	val := reflect.Indirect(reflect.ValueOf(&row))
	if val.Kind() != reflect.Struct {
		return row, errors.New("FixedWidthDecoder: unsupported type")
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() != reflect.String {
			continue
		}
		data, err := io.ReadAll(charset.NewDecoder(strings.NewReader(field.String()), dec.opts.Charset))
		if err != nil {
			return row, err
		}
		field.SetString(strings.TrimSpace(string(data)))
	}
	return row, nil
}
