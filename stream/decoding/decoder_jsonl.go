// Package decoding handles decoding of CSV, TSV, JSON, and fixed length text.
package decoding

import (
	"encoding/json"
	"io"
)

var (
	_ Decoder[any] = (*jsonlDecoder[any])(nil)
)

// NewJSONLDecoder creates a decoder that processes JSONL.
// Note: It assumes row-oriented JSONL, not JSON with an array directly under the root.
func NewJSONLDecoder[TRow any](src io.Reader) Decoder[TRow] {
	return &jsonlDecoder[TRow]{json.NewDecoder(src)}
}

type jsonlDecoder[TRow any] struct {
	reader *json.Decoder
}

func (dec *jsonlDecoder[TRow]) Decode() (TRow, error) {
	var row TRow
	err := dec.reader.Decode(&row)
	return row, err
}
