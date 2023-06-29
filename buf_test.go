package codefmt

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBuf(t *testing.T) {
	cases := []struct {
		name     string
		fn       func(buf *Buf)
		expected string
		checks   []func(*Buf)
	}{
		{
			name: "empty buf",
			fn: func(buf *Buf) {
				buf.Write("buf")
			},
			expected: "buf",
		},
		{
			name: "Writef",
			fn: func(buf *Buf) {
				buf.Write("%s:%d", "a", 1)
			},
			expected: "a:1",
		},
		{
			name: "Sub",
			fn: func(buf *Buf) {
				buf.Sub("${greet}, ${name}!",
					map[string]any{
						"greet": "Hello",
						"name":  "World",
					},
				)
			},
			expected: "Hello, World!",
		},
		{
			name: "Expand",
			fn: func(buf *Buf) {
				buf.Expand("${a}:${b}", "a", 1, "b", 2)
			},
			expected: "1:2",
		},
		{
			name: "Writeln",
			fn: func(buf *Buf) {
				buf.Writeln("%s, %s, %s", "x", "y", "z")
			},
			expected: "x, y, z\n",
			checks: []func(buf *Buf){
				func(buf *Buf) {
					assert.True(t, strings.HasSuffix(buf.String(), "\n"))
				},
			},
		},
		{
			name: "Preln",
			fn: func(buf *Buf) {
				buf.Preln("a:%d", 1)
				assert.True(t, strings.HasPrefix(buf.String(), "\n"))
			},
			expected: "\na:1",
		},
		{
			name: "Both",
			fn: func(buf *Buf) {
				buf.Both("b:%d", 2)
				assert.True(t,
					strings.HasPrefix(buf.String(), "\n") &&
						strings.HasSuffix(buf.String(), "\n"))
			},
			expected: "\nb:2\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := NewBuf()
			c.fn(buf)
			assert.Equal(t, c.expected, buf.String())
			assert.Equal(t, []byte(c.expected), buf.Bytes())
			if len(c.checks) > 0 {
				for _, fn := range c.checks {
					fn(buf)
				}
			}
			out := bytes.NewBufferString("")
			buf.out = out
			buf.Stdout()
			assert.Equal(t, c.expected, out.String())

			err := bytes.NewBufferString("")
			buf.err = err
			buf.Stderr()
			assert.Equal(t, c.expected, err.String())
		})
	}
}
