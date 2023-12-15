// Package buffering provides a mechanism for peeking at CSV header lines.
package buffering

import (
	"github.com/jszwec/csvutil"
)

var _ HeaderAwareCsvReader = (*headerAwareCsvReader)(nil)

type HeaderAwareCsvReader interface {
	csvutil.Reader
	Header() ([]string, error)
}

type HeaderAwareCsvReaderOptions struct {
	// Header
	// if none specified: first line is header, second and later lines are data
	// If one or more are specified: all after the first line is data
	Header []string
	// HeaderAliases
	// Convert values obtained as headers (either explicitly in Header or from the first line of encoding/csv.Reader) before they are actually processed by csvutil's CSV decoder.
	HeaderAliases map[string]string
}

type headerAwareCsvReader struct {
	r                 csvutil.Reader
	header            []string
	headerReadError   error
	headerAlreadyRead bool
}

func (h *headerAwareCsvReader) Read() ([]string, error) {
	if !h.headerAlreadyRead {
		// Return cache result only the first time
		h.headerAlreadyRead = true
		return h.header, h.headerReadError
	}
	return h.r.Read()
}

func (h *headerAwareCsvReader) Header() ([]string, error) {
	return h.header, h.headerReadError
}

// NewHeaderAwareCsvReader creates a wrapper that can preload headers without affecting the call to Read.
func NewHeaderAwareCsvReader(r csvutil.Reader, opts HeaderAwareCsvReaderOptions) HeaderAwareCsvReader {
	h := &headerAwareCsvReader{r: r}
	if len(opts.Header) > 0 {
		// If a header is explicitly specified, subsequent read operations are performed as usual
		h.headerAlreadyRead = true
		h.header = opts.Header
	} else {
		// If a header needs to be detected, it is read here and cached for the first read operation
		h.header, h.headerReadError = h.r.Read()
	}
	// Replace aliases if any match
	for index := range h.header {
		if opts.HeaderAliases != nil {
			if alias, ok := opts.HeaderAliases[h.header[index]]; ok {
				h.header[index] = alias
			}
		}
	}
	return h
}
