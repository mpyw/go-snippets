// Package charset provides an easy-to-use wrapper implementation for charset conversion.
package charset

import (
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Charset int

const (
	UTF8 Charset = iota
	ShiftJIS
)

// NewDecoder converts the specified input to UTF-8.
func NewDecoder(src io.Reader, enc Charset) io.Reader {
	switch enc {
	case ShiftJIS:
		return transform.NewReader(src, japanese.ShiftJIS.NewDecoder())
	case UTF8:
		fallthrough
	default:
		return src
	}
}

// NewEncoder converts UTF-8 input to the specified output.
func NewEncoder(dst io.Writer, enc Charset) io.Writer {
	switch enc {
	case ShiftJIS:
		return transform.NewWriter(dst, japanese.ShiftJIS.NewEncoder())
	case UTF8:
		fallthrough
	default:
		return dst
	}
}
