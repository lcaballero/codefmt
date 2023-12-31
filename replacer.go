package codefmt

import (
	"fmt"
	"regexp"
	"strings"
)

// MapReplaceRE is for finding replacements of placeholders '${...}'
var MapReplaceRE = regexp.MustCompile("\\${[^{}]*}[\t\n]*")

// Replacer is a partial application that holds a string template and
// accepts Replacements map to fill in placeholders with values
type Replacer func(m map[string]any) string

// Replace turns the given pairs into a map and executes the resulting
// Replacer
func (r Replacer) Replace(pairs ...any) string {
	return r(ToPairs(pairs...))
}

// BraceTemplate is simply a string with a template based on ${...}
// replacements that can be executed using the Replace and MapPair
// functions
type BraceTemplate string

func (b BraceTemplate) Replace(pairs ...any) string {
	return MapReplacer(string(b)).Replace(pairs...)
}

// Replace converts the pairs of strings into a map where the keys can
// be replace when found in the BraceTemplate
func (b BraceTemplate) MapPair(pairs ...string) string {
	args := make([]any, len(pairs))
	for i := 0; i < len(pairs); i++ {
		args[i] = pairs[i]
	}
	return MapReplacer(string(b)).Replace(args...)
}

// MapReplacer creates a Replacer capable of substituting placeholders
// in the input string in the form of "${...}" where the key can be
// found in the input map and should is replaced with the
// corresponding value.
//
// Placeholders can include notation for processing values.  Given a
// key/value map like so: {name:"world", greeting:"hello"} and a
// template of: '${^greeting}, ${^name}!' indicates that the values
// should be output in Pascal case.  Additionally, "${@name}" is
// notation for Camel case.  Where this notation is used the values
// any whitespace or hyphens are removed.
func MapReplacer(s string) Replacer {
	return func(m map[string]any) string {
		return MapReplaceRE.ReplaceAllStringFunc(s, func(name string) string {
			orig := name
			trimmed := strings.TrimSpace(orig)
			name = strings.TrimPrefix(trimmed, "${")
			name = strings.TrimSuffix(name, "-}")
			name = strings.TrimSuffix(name, "}")
			isPascal := strings.HasPrefix(name, "^")
			isCamel := strings.HasPrefix(name, "@")
			if isPascal {
				name = strings.TrimPrefix(name, "^")
			}
			if isCamel {
				name = strings.TrimPrefix(name, "@")
			}
			v, ok := m[name]
			if !ok {
				return name
			}
			if isPascal {
				v = ToPascal(fmt.Sprintf("%v", v))
			}
			if isCamel {
				v = ToCamel(fmt.Sprintf("%v", v))
			}
			return fmt.Sprintf("%v", v)
		})
	}
}

func camel(s string) string {
	lower := strings.ToLower(s)
	switch lower {
	case "id", "html", "url":
		return lower
	default:
		return strings.ToLower(string(s[0])) + string(s[1:])
	}
}

// ToCamel formats the string as a camel cased string
func ToCamel(s string) string {
	fn := func(i int, ks string) string {
		if i == 0 {
			return camel(ks)
		}
		return pascal(ks)
	}
	return deKebob(s, fn)
}

// deKebob splits the string on spaces and hyphens and applies the
// func to each resulting word
func deKebob(s string, fn func(int, string) string) string {
	words := strings.Split(strings.Replace(s, "-", " ", -1), " ")
	for i, w := range words {
		words[i] = fn(i, w)
	}
	return strings.Join(words, "")
}

// pascal converts a single word to pascal case
func pascal(s string) string {
	upper := strings.ToUpper(s)
	switch upper {
	case "ID", "HTML", "URL":
		return upper
	default:
		return strings.ToUpper(string(s[0])) + string(s[1:])
	}
}

// ToPascal formats the string as pascal cased string
func ToPascal(s string) string {
	fn := func(i int, str string) string {
		return pascal(str)
	}
	return deKebob(s, fn)
}

// ToPairs creates a map using the pairs as key/value pairs
func ToPairs(args ...any) map[string]any {
	m := map[string]any{}
	for i := 1; i < len(args); i = i + 2 {
		k0, v0 := args[i-1], args[i]
		key, val := fmt.Sprintf("%s", k0), fmt.Sprintf("%v", v0)
		m[key] = val
	}
	return m
}
