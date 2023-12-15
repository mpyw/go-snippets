package newline

type Newline int

const (
	Lf Newline = iota
	Crlf
)

func (nl Newline) ToString() string {
	switch nl {
	case Crlf:
		return "\r\n"
	case Lf:
		fallthrough
	default:
		return "\n"
	}
}
