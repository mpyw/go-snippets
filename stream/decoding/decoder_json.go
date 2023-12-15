// Package decoding handles decoding of CSV, TSV, JSON, and fixed length text.
package decoding

import (
	"encoding/json"
	"io"
)

var (
	_ Decoder[any] = (*jsonDecoder[any])(nil)
)

// NewJSONDecoder creates a decoder that processes JSON.
// Note: This is not row-oriented JSONL, but JSON with an array directly under the root.
func NewJSONDecoder[TRow any](src io.Reader) Decoder[TRow] {
	var rows []TRow
	data, err := io.ReadAll(src)
	if err != nil {
		return &invalidDecoder[TRow]{err}
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		return &invalidDecoder[TRow]{err}
	}
	return &jsonDecoder[TRow]{rows, 0}
}

type jsonDecoder[TRow any] struct {
	rows       []TRow
	currentPos int
}

func (dec *jsonDecoder[TRow]) Decode() (TRow, error) {
	var row TRow
	if dec.currentPos >= len(dec.rows) {
		return row, io.EOF
	}
	defer func() {
		dec.currentPos++
	}()
	return dec.rows[dec.currentPos], nil
}
