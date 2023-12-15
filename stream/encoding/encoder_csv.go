// Package encoding handles the encoding process for CSV, TSV, and fixed length text.
package encoding

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
	"github.com/mpyw/go-snippets/stream/charset"
	"github.com/mpyw/go-snippets/stream/newline"
)

var (
	_ Encoder[any] = (*csvEncoder[any])(nil)
)

// CsvEncoderOptions is an optional argument to NewCsvEncoder.
type CsvEncoderOptions struct {
	// Charset is an output character code (default: charset.UTF8)
	Charset charset.Charset
	// Separator default: ','
	Separator rune
	// OmitHeader is whether to skip generating the header on the first line from the field from the structure definition
	OmitHeader bool
	// Newline default: LF
	Newline newline.Newline
}

// NewCsvEncoder creates an Encoder that processes CSV/TSV.
// See csvutil documentation for structure tags: https://github.com/jszwec/csvutil
func NewCsvEncoder[TRow any](dst io.Writer, opts CsvEncoderOptions) Encoder[TRow] {
	w := csv.NewWriter(charset.NewEncoder(dst, opts.Charset))
	if opts.Separator != 0 {
		w.Comma = opts.Separator
	}
	if opts.Newline == newline.Crlf {
		w.UseCRLF = true
	}
	enc := csvutil.NewEncoder(w)
	if opts.OmitHeader {
		enc.AutoHeader = false
	}
	return &csvEncoder[TRow]{w, enc}
}

type csvEncoder[TRow any] struct {
	w   *csv.Writer
	csv *csvutil.Encoder
}

func (enc *csvEncoder[TRow]) Encode(row TRow) error {
	return enc.csv.Encode(row)
}

func (enc *csvEncoder[TRow]) Flush() error {
	enc.w.Flush()
	return enc.w.Error()
}
