// Package encoding handles the encoding process for CSV, TSV, and fixed length text.
package encoding

// Encoder handles the encoding of CSV, TSV, JSON, and fixed length text.
type Encoder[TRow any] interface {
	// Encode encodes and writes arguments.
	Encode(row TRow) error
	// Flush flushes data. Call this before the end of processing.
	Flush() error
}
