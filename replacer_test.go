package codefmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapReplacer(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
		kv       []string
	}{
		{
			name:     "hw",
			input:    "${greeting}, ${name}!",
			expected: "Hello, World!",
			kv: []string{
				"greeting", "Hello",
				"name", "World",
			},
		},
		{
			name:     "hw with pascal",
			input:    "${^greeting}, ${^name}!",
			expected: "Hello, World!",
			kv: []string{
				"greeting", "hello",
				"name", "world",
			},
		},
		{
			name:     "hw with camel",
			input:    "${^greeting}, ${@name}!",
			expected: "Hello, world!",
			kv: []string{
				"greeting", "hello",
				"name", "World",
			},
		},
		{
			name:     "hw uses key when no map to value provided",
			input:    "${^greeting}, ${@name}!",
			expected: "greeting, name!",
			kv:       []string{},
		},
		{
			name:     "all modifiers",
			input:    "${^greeting}, ${@name-}!",
			expected: "Welcome, everyone!",
			kv: []string{
				"greeting", "welcome",
				"name", "Everyone",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expected := BraceTemplate(c.input).MapPair(c.kv...)
			actual := c.expected
			assert.Equal(t, expected, actual)
		})
	}
}

func Test_ToPascal(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ToPascal: catWalk",
			input:    "catWalk",
			expected: "CatWalk",
		},
		{
			name:     "ToPascal: cat-walk",
			input:    "cat-walk",
			expected: "CatWalk",
		},
		{
			name:     "ToPascal: html-page",
			input:    "html-page",
			expected: "HTMLPage",
		},
		{
			name:     "ToPascal: html-hash-id",
			input:    "HTML-hash-id",
			expected: "HTMLHashID",
		},
		{
			name:     "ToPascal: home url",
			input:    "home url",
			expected: "HomeURL",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, ToPascal(c.input))
		})
	}
}

func Test_ToCamel(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ToCamel",
			input:    "CatWalk",
			expected: "catWalk",
		},
		{
			name:     "ToCamel",
			input:    "cat-walk",
			expected: "catWalk",
		},
		{
			name:     "ToCamel",
			input:    "html-page",
			expected: "htmlPage",
		},
		{
			name:     "ToCamel",
			input:    "html-hash-id",
			expected: "htmlHashID",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, ToCamel(c.input))
		})
	}
}
