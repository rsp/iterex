// Package iterex provides iterator-based regular expressions for Go.
// Change `regexp` to `iterex`, change `All` to `Each` and iterate!
// This package is like the standard `regexp` with two differences:
package iterex

import (
	"iter"
	"regexp"
)

// Iterex is the representation of a compiled regular expression.
// It encapsulates [regexp.Regexp] (not to be used directly).
type Iterex struct {
	*regexp.Regexp
}

// Compile parses a regular expression and returns, if successful,
// an [Iterex] object that can be used to match against text.
// It uses syntax inspired by RE2, PCRE, Perl.
// For details see [regexp.Compile]
func Compile(str string) (*Iterex, error) {
	re, err := regexp.Compile(str)
	if err != nil {
		return nil, err
	}
	return &Iterex{re}, nil
}

// MustCompile is a version of Compile that panics on error.
func MustCompile(str string) *Iterex {
	re := regexp.MustCompile(str)
	return &Iterex{re}
}

// CompilePOSIX parses a regular expression and returns, if successful,
// an [Iterex] object that can be used to match against text.
// It uses POSIX ERE (egrep) syntax with leftmost-longest match semantics.
// For details see [regexp.CompilePOSIX]
func CompilePOSIX(str string) (*Iterex, error) {
	re, err := regexp.CompilePOSIX(str)
	if err != nil {
		return nil, err
	}
	return &Iterex{re}, nil
}

// MustCompilePOSIX is a version of CompilePOSIX that panics on error.
func MustCompilePOSIX(str string) *Iterex {
	re := regexp.MustCompilePOSIX(str)
	return &Iterex{re}
}

// Functions working on byte slices:

// FindEach is an iterator version of [regexp.FindAll].
// It returns an iterator that internally uses [regexp.FindIndex].
func (ir *Iterex) FindEach(b []byte, n ...int) iter.Seq[[]byte] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([]byte) bool) {
		var pos, i int
		for pos < len(b) && (lim < 0 || i < lim) {
			match := ir.FindIndex(b[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			if !yield(b[start:end]) {
				return
			}
			pos = end
			i++
		}
	}
}

// FindEachIndex is an iterator version of [regexp.FindAllIndex].
// It returns an iterator that internally uses [regexp.FindIndex].
func (ir *Iterex) FindEachIndex(b []byte, n ...int) iter.Seq[[]int] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([]int) bool) {
		var pos, i int
		for pos < len(b) && (lim < 0 || i < lim) {
			match := ir.FindIndex(b[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			if !yield([]int{start, end}) {
				return
			}
			pos = end
			i++
		}
	}
}

// FindEachSubmatch is an iterator version of [regexp.FindAllSubmatch]
// It returns an iterator that internally uses [regexp.FindSubmatchIndex].
func (ir *Iterex) FindEachSubmatch(b []byte, n ...int) iter.Seq[[][]byte] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([][]byte) bool) {
		var pos, i int
		for pos < len(b) && (lim < 0 || i < lim) {
			match := ir.FindSubmatchIndex(b[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			sub := [][]byte{b[start:end]}
			for i := 2; i < len(match); i += 2 {
				sstart := pos + match[i]
				send := pos + match[i+1]
				if sstart >= 0 {
					sub = append(sub, b[sstart:send])
				}
			}
			if !yield(sub) {
				return
			}
			pos = end
			i++
		}
	}

}

// FindEachSubmatchIndex is an iterator version of [regexp.FindAllSubmatchIndex]
// It returns an iterator that internally uses [regexp.FindSubmatchIndex].
func (ir *Iterex) FindEachSubmatchIndex(b []byte, n ...int) iter.Seq[[]int] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([]int) bool) {
		var pos, i int
		for pos < len(b) && (lim < 0 || i < lim) {
			match := ir.FindSubmatchIndex(b[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			sub := []int{start, end}
			for i := 2; i < len(match); i++ {
				sub = append(sub, pos+match[i])
			}
			if !yield(sub) {
				return
			}
			pos = end
			i++
		}
	}
}

// Functions working on strings

// FindEachString is an iterator version of [regexp.FindAllString]
// with optional limit n.
// It returns an iterator that internally uses [regexp.FindStringIndex].
func (ir *Iterex) FindEachString(s string, n ...int) iter.Seq[string] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func(string) bool) {
		var pos, i int
		for pos < len(s) && (lim < 0 || i < lim) {
			match := ir.FindStringIndex(s[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			if !yield(s[start:end]) {
				return
			}
			pos = end
			i++
		}
	}
}

// FindEachStringIndex is an iterator version of [regexp.FindAllStringIndex].
// with optional limit n.
// It returns an iterator that internally uses [regexp.FindStringIndex].
func (ir *Iterex) FindEachStringIndex(s string, n ...int) iter.Seq[[]int] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([]int) bool) {
		var pos, i int
		for pos < len(s) && (lim < 0 || i < lim) {
			match := ir.FindStringIndex(s[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			if !yield([]int{start, end}) {
				return
			}
			pos = end
			i++
		}
	}
}

// FindEachStringSubmatch is an iterator version of [regexp.FindAllStringSubmatch]
// It returns an iterator that internally uses [regexp.FindStringSubmatchIndex].
func (ir *Iterex) FindEachStringSubmatch(s string, n ...int) iter.Seq[[]string] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([]string) bool) {
		var pos, i int
		for pos < len(s) && (lim < 0 || i < lim) {
			match := ir.FindStringSubmatchIndex(s[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			sub := []string{s[start:end]}
			for i := 2; i < len(match); i += 2 {
				sstart := pos + match[i]
				send := pos + match[i+1]
				if sstart >= 0 {
					sub = append(sub, s[sstart:send])
				}
			}
			if !yield(sub) {
				return
			}
			pos = end
			i++
		}
	}
}

// FindEachStringSubmatchIndex is an iterator version of [regexp.FindAllStringSubmatchIndex]
// It returns an iterator that internally uses [regexp.FindStringIndex].
func (ir *Iterex) FindEachStringSubmatchIndex(s string, n ...int) iter.Seq[[]int] {
	var lim int
	if len(n) > 0 {
		lim = n[0]
	} else {
		lim = -1
	}
	return func(yield func([]int) bool) {
		var pos, i int
		for pos < len(s) && (lim < 0 || i < lim) {
			match := ir.FindStringSubmatchIndex(s[pos:])
			if match == nil {
				break
			}
			start := pos + match[0]
			end := pos + match[1]
			sub := []int{start, end}
			for i := 2; i < len(match); i++ {
				sub = append(sub, pos+match[i])
			}
			if !yield(sub) {
				return
			}
			pos = end
			i++
		}
	}
}
