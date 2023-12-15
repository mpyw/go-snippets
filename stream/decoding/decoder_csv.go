// Package decoding handles decoding of CSV, TSV, JSON, and fixed length text.
package decoding

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
	"github.com/mpyw/go-snippets/stream/charset"
	"github.com/mpyw/go-snippets/stream/decoding/adjuster"
	"github.com/mpyw/go-snippets/stream/decoding/buffering"
)

var (
	_ Decoder[any] = (*csvDecoder[any])(nil)
)

// CsvDecoderOptions is an optional argument to NewCsvDecoder.
type CsvDecoderOptions struct {
	// Charset is an input character encoding (default: charset.UTF8)
	Charset charset.Charset
	// Charset default: ','
	Separator rune
	// Header
	// If none specified: 1st line is header, 2nd and later lines are data
	// If one or more are specified: everything after the first line is data
	Header []string
	// HeaderAliases
	// Convert values obtained as headers (either explicitly in Header or from the first line of io.Reader) before they are actually processed by csvutil's CSV decoder.
	HeaderAliases map[string]string
}

// NewCsvDecoder creates a Decoder that processes CSV/TSV.
// See csvutil documentation for structure tags: https://github.com/jszwec/csvutil
func NewCsvDecoder[TRow any](src io.Reader, opts CsvDecoderOptions) Decoder[TRow] {
	r := csv.NewReader(charset.NewDecoder(src, opts.Charset))

	// In the encoding/csv.Reader layer, the number of fields in the first line does not necessarily have to match the number of fields in the second and subsequent lines.
	// Field count checking is done in the csvutil layer.
	r.FieldsPerRecord = -1

	if opts.Separator != 0 {
		r.Comma = opts.Separator
	}

	// Error if header read fails.
	// To unify the interface, we will return a Decoder that returns an error no matter what we do.
	dec, err := csvutil.NewDecoder(
		// Trim off the extra fields based on the number of fields in the header.
		// csvutil requires an exact match on the number of fields, so be sure to avoid errors for extra fields.
		adjuster.NewFieldAdjustableCsvReader(
			// We want to know in advance how many fields are in the header,
			// but reading the first line would change the result of r.Read(),
			// so wrap r to read the header without affecting it.
			// Also convert the value read from the header before inputting it into csvutil.
			buffering.NewHeaderAwareCsvReader(r, buffering.HeaderAwareCsvReaderOptions{
				Header:        opts.Header,
				HeaderAliases: opts.HeaderAliases,
			}),
		),
		opts.Header...,
	)
	if err != nil {
		return &invalidDecoder[TRow]{err}
	}
	// No missing fields allowed
	dec.DisallowMissingColumns = true
	return &csvDecoder[TRow]{dec}
}

type csvDecoder[TRow any] struct {
	csv *csvutil.Decoder
}

func (dec *csvDecoder[TRow]) Decode() (TRow, error) {
	var row TRow
	err := dec.csv.Decode(&row)
	return row, err
}
