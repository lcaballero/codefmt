package codefmt

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter int

func (w mockWriter) Write(b []byte) (int, error) {
	return 0, fmt.Errorf("expected error")
}

func TestIndent(t *testing.T) {
	cases := []struct {
		name      string
		input     Indent
		expected  string
		hasIndent bool
		writer    io.Writer
	}{
		{
			name:     "default new",
			input:    NewIndent(),
			expected: "",
		},
		{
			name:      "2x indenting",
			input:     NewIndent().Next().Next(),
			expected:  "    ",
			hasIndent: true,
		},
		{
			name:      "2x non-ws",
			input:     Indent{Level: 0, Increment: 1, Tab: "123"}.Next().Next(),
			expected:  "123123",
			hasIndent: true,
		},
		{
			name:     "empty indented 2x",
			input:    Indent{}.Next().Next(),
			expected: "",
		},
		{
			name:      "has indent",
			input:     NewIndent().Next(),
			expected:  "  ",
			hasIndent: true,
		},
		{
			name:      "has indent",
			input:     NewIndent().Next(),
			expected:  "  ",
			hasIndent: true,
			writer:    mockWriter(0),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, c.input.String())
			assert.Equal(t, c.hasIndent, c.input.HasIndent())

			buf := bytes.NewBufferString("")
			c.input.WriteTo(buf)
			assert.Equal(t, c.expected, buf.String())

			if c.input.HasIndent() {
				level := c.input.Level
				prev := c.input.Prev()
				assert.Equal(t, level-1, prev.Level)
			}

			if c.writer != nil {
				err := c.input.WriteTo(c.writer)
				assert.NotNil(t, err)
			}

			// only de-indent there is an indent
			if c.input.Level > 0 {
				curr := c.input.String()
				indent := c.input

				t.Logf("i: -, level: %d, curr: '%s' == '%s', eq? %v",
					indent.Level, curr, indent.String(),
					curr == indent.String(),
				)

				// check each level reduces to 0
				for i := c.input.Level; i >= 0; i-- {

					t.Logf("i: %d, level: %d, curr: '%s' == '%s', eq? %v",
						i, indent.Level, curr, indent.String(),
						curr == indent.String(),
					)

					indent = indent.Prev()
					curr = strings.TrimSuffix(curr, indent.Tab)

					assert.Equal(t,
						curr, indent.String(),
						"indent curr: '%s',  does not have correct prev",
					)
				}
			}
		})
	}
}
