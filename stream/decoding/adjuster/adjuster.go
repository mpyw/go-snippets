// Package adjuster provides a mechanism to relax the requirement for the number of CSV fields.
package adjuster

import (
	"errors"

	"github.com/mpyw/go-snippets/stream/decoding/buffering"
)

var _ buffering.HeaderAwareCsvReader = (*fieldAdjustableCsvReader)(nil)

var ErrCsvHeaderLoadFailure = errors.New("failed to load csv header")

type fieldAdjustableCsvReader struct {
	buffering.HeaderAwareCsvReader
}

func (c *fieldAdjustableCsvReader) Read() ([]string, error) {
	header, err := c.Header()
	if err != nil {
		return nil, errors.Join(ErrCsvHeaderLoadFailure, err)
	}
	row, err := c.HeaderAwareCsvReader.Read()
	if err != nil {
		return nil, err
	}
	// Return as-is to trigger csvutil error handling if the number of fields is small
	if len(row) <= len(header) {
		return row, nil
	}
	// Leave fields where necessary.
	return row[:len(header)], nil
}

// NewFieldAdjustableCsvReader
// csvutil requires that the number of fields be exactly the same.
// This does not take into account the requirement that extra fields may have been added on the right side,
// so a wrapper is provided with additional processing to remove extra fields if they are present.
func NewFieldAdjustableCsvReader(h buffering.HeaderAwareCsvReader) buffering.HeaderAwareCsvReader {
	return &fieldAdjustableCsvReader{h}
}
