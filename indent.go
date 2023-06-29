package codefmt

import (
	"bytes"
	"io"
)

const (
	// DefaultLevel starts indent at column 0 without any indention
	DefaultLevel = 0

	// DefaultIncrement is 1 meaning each indention is a single set of
	// indention strings
	DefaultIncrement = 1

	// DefaultTab is 2 spaces multiplied by the Increment that is
	// typically 1
	DefaultTab = "  "
)

// Indent represents indention at a given level.
type Indent struct {
	Level     int
	Increment int
	Tab       string
}

// Returns an Indent value starting at level 0, with an increment
// of 0, and a tab of 2 spaces.
func NewIndent() Indent {
	return Indent{
		Level:     DefaultLevel,
		Increment: DefaultIncrement,
		Tab:       DefaultTab,
	}
}

// String produces the indention for the given level of the Indent.
func (n Indent) String() string {
	buf := bytes.NewBuffer([]byte{})
	n.WriteTo(buf)
	return buf.String()
}

// HasIndent returns true if the Inc is > 0 and Tab != ”.
func (n Indent) HasIndent() bool {
	return n.Level > 0
}

// WriteTo outputs the Indent to the Writer.
func (n Indent) WriteTo(w io.Writer) error {
	for i := 0; i < n.Level; i++ {
		n, err := w.Write([]byte(n.Tab))
		if n <= 0 || err != nil {
			return err
		}
	}
	return nil
}

// Next adds one level to the Indent.
func (n Indent) Next() Indent {
	return Indent{
		Level:     n.Level + n.Increment,
		Tab:       n.Tab,
		Increment: n.Increment,
	}
}

// Prev reduces Indent by one level.
func (n Indent) Prev() Indent {
	if (n.Level - n.Increment) < 0 {
		return Indent{
			Level:     0,
			Increment: n.Increment,
			Tab:       n.Tab,
		}
	}
	return Indent{
		Level:     n.Level - n.Increment,
		Increment: n.Increment,
		Tab:       n.Tab,
	}
}
