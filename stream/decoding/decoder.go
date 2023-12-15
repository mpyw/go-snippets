// Package decoding handles decoding of CSV, TSV, JSON, and fixed length text.
package decoding

var (
	_ Decoder[any] = (*invalidDecoder[any])(nil)
)

// Decoder handles decoding of CSV, TSV, JSON, and fixed length text.
type Decoder[TRow any] interface {
	Decode() (TRow, error)
}

type invalidDecoder[TRow any] struct {
	err error
}

func (dec *invalidDecoder[TRow]) Decode() (TRow, error) {
	var empty TRow
	return empty, dec.err
}
