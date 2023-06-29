package codefmt // import "github.com/lcaballero/codefmt"

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Buf is a formatting interface for buffering code.
type Buf struct {
	buf *bytes.Buffer
	out io.Writer
	err io.Writer
}

// NewBuf creates an empty Buf instance
func NewBuf() *Buf {
	return &Buf{
		buf: bytes.NewBufferString(""),
		out: os.Stdout,
		err: os.Stderr,
	}
}

// Bytes returns a slice of bytes that make up the buffer
func (b *Buf) Bytes() []byte {
	return b.buf.Bytes()
}

// String returns a string of the buffer contents
func (b *Buf) String() string {
	return b.buf.String()
}

// Stdout writes buffer to stdout
func (b *Buf) Stdout() *Buf {
	fmt.Fprint(b.out, b)
	return b
}

// Stderr write buffer to stderr
func (b *Buf) Stderr() *Buf {
	fmt.Fprint(b.err, b)
	return b
}

// Writef buffer the resulting string after applying standard Sprintf
// formating
func (b *Buf) Write(format string, args ...any) *Buf {
	s := fmt.Sprintf(format, args...)
	b.buf.WriteString(s)
	return b
}

// Sub creates a replacement template that can substitute `${key}` for
// `value` using the given map.
func (b *Buf) Sub(format string, subs map[string]any) *Buf {
	b.buf.WriteString(MapReplacer(format)(subs))
	return b
}

// Expand creates a replacement template using every two args as pairs
// to build a map that can be passed to Sub
func (b *Buf) Expand(format string, args ...any) *Buf {
	text := BraceTemplate(format).Replace(args...)
	b.buf.WriteString(text)
	return b
}

// Writeln uses the format and args to create a string to buffer the
// result.
func (b *Buf) Writeln(format string, args ...any) *Buf {
	s := fmt.Sprintf(format, args...)
	b.buf.WriteString(s)
	return b.NL()
}

// NL adds a newline to the buffer
func (b *Buf) NL() *Buf {
	b.buf.Write([]byte{'\n'})
	return b
}

// Preln adds a newline before buffering the resulting formated string
func (b *Buf) Preln(format string, args ...any) *Buf {
	return b.NL().Write(format, args...)
}

// Both adds a newline before and after the string that is buffered as
// a result of using the format and args to produce a Sprintf like
// string
func (b *Buf) Both(format string, args ...any) *Buf {
	return b.NL().Write(format, args...).NL()
}
