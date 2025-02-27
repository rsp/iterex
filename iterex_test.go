package iterex

import (
	"reflect"
	"testing"
)

// - Internal tests:

// TestFindString tests if receiver functions from regexp package work transparently
func TestFindString(t *testing.T) {
	r := MustCompile(`b+`)
	s := "aaabbbccc"
	want := "bbb"
	got := r.FindString(s)
	if got != want {
		t.Errorf("FindString: want %s, got %s", want, got)
	}
}

// - Tests for compile functions:

// TestCompileSuccess tests Compile for correct pattern
func TestCompileSuccess(t *testing.T) {
	p := `\d+`
	r, err := Compile(p)
	if err != nil {
		t.Errorf("Compile: unexpected error: %v", err)
	}
	if r == nil {
		t.Errorf("Compile: retured nil")
	}
}

// TestCompileError tests Compile for incorrect pattern
func TestCompileError(t *testing.T) {
	p := `a+[`
	r, err := Compile(p)
	if err == nil {
		t.Errorf("Compile: should return error")
	}
	if r != nil {
		t.Errorf("Compile: should return nil")
	}
}

// TestCompilePosixSuccess tests CompilePOSIX for correct pattern
func TestCompilePosixSuccess(t *testing.T) {
	p := `a+`
	r, err := CompilePOSIX(p)
	if err != nil {
		t.Errorf("CompilePOSIX: unexpected error: %v", err)
	}
	if r == nil {
		t.Errorf("CompilePOSIX: retured nil")
	}
}

// TestCompilePosixError tests CompilePOSIX for incorrect pattern
// (which would be valid for Compile)
func TestCompilePosixError(t *testing.T) {
	p := `\d+`
	r, err := CompilePOSIX(p)
	if err == nil {
		t.Errorf("CompilePOSIX: should return error")
	}
	if r != nil {
		t.Errorf("CompilePOSIX: should return nil")
	}
}

// - Tests for difference between PCRE and POSIX:

// TestMustCompile tests MustCompile for PCRE-like behavior
func TestMustCompile(t *testing.T) {
	p := `ab|abc`
	s := "abcdef"
	want := []string{"ab"}
	var got []string
	ir := MustCompile(p)
	for m := range ir.FindEachString(s) {
		got = append(got, m)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MustCompile and FindEachString: got %v, want %v", got, want)
	}
}

// TestMustCompilePosix tests MustCompilePOSIX for POSIX-like behavior
func TestMustCompilePosix(t *testing.T) {
	p := `ab|abc`
	s := "abcdef"
	want := []string{"abc"}
	var got []string
	ir := MustCompilePOSIX(p)
	for m := range ir.FindEachString(s) {
		got = append(got, m)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MustCompilePOSIX and FindEachString: got %v, want %v", got, want)
	}
}

// - Tests working on strings:

//   - FindEachString:

// TestFindEachString tests all matches by
// FindEachString with no limit
func TestFindEachString(t *testing.T) {
	lim := -1
	ir := MustCompile(`x+`)
	s := "aaaxxaaaxxxaxxxxaaax"
	want := ir.FindAllString(s, lim)
	var got []string
	it := ir.FindEachString(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachString: got %v, want %v", got, want)
	}
}

// TestFindEachStringWithLimit tests matches by
// FindEachString with limit
func TestFindEachStringWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`x+`)
	s := "aaaxxaaaxxxaxxxxaaax"
	want := ir.FindAllString(s, lim)
	var got []string
	it := ir.FindEachString(s, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachString: got %v, want %v", got, want)
	}
}

// TestFindEachStringNoMatch tests matches by
// FindEachString with no match
func TestFindEachStringNoMatch(t *testing.T) {
	ir := MustCompile(`x+`)
	s := "aaaxxaaaxxxaxxxxaaax"
	want := []string{}
	var got []string
	it := ir.FindEachString(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachString: got %v, want %v", got, want)
	}
}

//   - FindEachStringIndex:

// TestFindEachStringIndex tests all matches by
// FindEachStringIndex with no limit
func TestFindEachStringIndex(t *testing.T) {
	lim := -1
	ir := MustCompile(`x+`)
	s := "aaaxxaaaxxxaxxxxaaax"
	want := ir.FindAllStringIndex(s, lim)
	var got [][]int
	it := ir.FindEachStringIndex(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringIndex: got %v, want %v", got, want)
	}
}

// TestFindEachStringIndexWithLimit tests matches by
// FindEachStringIndex with limit
func TestFindEachStringIndexWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`x+`)
	s := "aaaxxaaaxxxaxxxxaaax"
	want := ir.FindAllStringIndex(s, lim)
	var got [][]int
	it := ir.FindEachStringIndex(s, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringIndex: got %v, want %v", got, want)
	}
}

// TestFindEachStringIndexNoMatch tests matches by
// FindEachStringIndex with no match
func TestFindEachStringIndexNoMatch(t *testing.T) {
	ir := MustCompile(`X+`)
	s := "aaaxxaaaxxxaxxxxaaax"
	want := [][]int{}
	var got [][]int
	it := ir.FindEachStringIndex(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringIndex: got %v, want %v", got, want)
	}
}

//   - FindEachStringSubmatch:

// TestFindEachStringSubmatch tests all matches by
// FindEachStringSubmatch with no limit
func TestFindEachStringSubmatch(t *testing.T) {
	lim := -1
	ir := MustCompile(`(x+)(y+)`)
	s := "aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy"
	want := ir.FindAllStringSubmatch(s, lim)
	var got [][]string
	it := ir.FindEachStringSubmatch(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringSubmatch: got %v, want %v", got, want)
	}
}

// TestFindEachStringSubmatchWithLimit tests matches by
// FindEachStringSubmatch with limit
func TestFindEachStringSubmatchWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`(x+)(y+)`)
	s := "aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy"
	want := ir.FindAllStringSubmatch(s, lim)
	var got [][]string
	it := ir.FindEachStringSubmatch(s, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringSubmatch: got %v, want %v", got, want)
	}
}

// TestFindEachStringSubmatchNoMatch tests matches by
// FindEachStringSubmatch with no match
func TestFindEachStringSubmatchNoMatch(t *testing.T) {
	ir := MustCompile(`(x+)(Y+)`)
	s := "aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy"
	want := [][]string{}
	var got [][]string
	it := ir.FindEachStringSubmatch(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringSubmatch: got %v, want %v", got, want)
	}
}

//   - FindAllStringSubmatchIndex:

// TestFindEachStringSubmatchIndex tests all matches by
// FindEachStringSubmatchIndex with no limit
func TestFindEachStringSubmatchIndex(t *testing.T) {
	lim := -1
	ir := MustCompile(`(x+)(y+)`)
	s := "aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy"
	want := ir.FindAllStringSubmatchIndex(s, lim)
	var got [][]int
	it := ir.FindEachStringSubmatchIndex(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringSubmatchIndex: got %v, want %v", got, want)
	}
}

// TestFindEachStringSubmatchIndexWithLimit tests matches by
// FindEachStringSubmatchIndex with limit
func TestFindEachStringSubmatchIndexWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`(x+)(y+)`)
	s := "aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy"
	want := ir.FindAllStringSubmatchIndex(s, lim)
	var got [][]int
	it := ir.FindEachStringSubmatchIndex(s, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringSubmatchIndex: got %v, want %v", got, want)
	}
}

// TestFindEachStringSubmatchIndexNoMatch tests matches by
// FindEachStringSubmatchIndex with limit
func TestFindEachStringSubmatchIndexNoMatch(t *testing.T) {
	ir := MustCompile(`(x+)(y+)`)
	s := "aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy"
	want := [][]int{}
	var got [][]int
	it := ir.FindEachStringSubmatchIndex(s)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachStringSubmatchIndex: got %v, want %v", got, want)
	}
}

// - Tests working on byte slices:

//   - FindEach:

// TestFindEach tests all matches by
// FindEach with no limit
func TestFindEach(t *testing.T) {
	lim := -1
	ir := MustCompile(`x+`)
	b := []byte("aaaxxaaaxxxaxxxxaaax")
	want := ir.FindAll(b, lim)
	var got [][]byte
	it := ir.FindEach(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEach: got %v, want %v", got, want)
	}
}

// TestFindEachWithLimit tests matches by
// FindEach with limit
func TestFindEachWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`x+`)
	b := []byte("aaaxxaaaxxxaxxxxaaax")
	want := ir.FindAll(b, lim)
	var got [][]byte
	it := ir.FindEach(b, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEach: got %v, want %v", got, want)
	}
}

// TestFindEachNoMatch tests matches by
// FindEach with no match
func TestFindEachNoMatch(t *testing.T) {
	ir := MustCompile(`X+`)
	b := []byte("aaaxxaaaxxxaxxxxaaax")
	want := [][]byte{}
	var got [][]byte
	it := ir.FindEach(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEach: got %v, want %v", got, want)
	}
}

//   - FindEachIndex:

// TestFindEachIndex tests all matches by
// FindEachIndex with no limit
func TestFindEachIndex(t *testing.T) {
	lim := -1
	ir := MustCompile(`x+`)
	b := []byte("aaaxxaaaxxxaxxxxaaax")
	want := ir.FindAllIndex(b, lim)
	var got [][]int
	it := ir.FindEachIndex(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachIndex: got %v, want %v", got, want)
	}
}

// TestFindEachIndexWithLimit tests matches by
// FindEachIndex with limit
func TestFindEachIndexWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`x+`)
	b := []byte("aaaxxaaaxxxaxxxxaaax")
	want := ir.FindAllIndex(b, lim)
	var got [][]int
	it := ir.FindEachIndex(b, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachIndex: got %v, want %v", got, want)
	}
}

// TestFindEachIndexNoMatch tests matches by
// FindEachIndex with no match
func TestFindEachIndexNoMatch(t *testing.T) {
	ir := MustCompile(`X+`)
	b := []byte("aaaxxaaaxxxaxxxxaaax")
	want := [][]int{}
	var got [][]int
	it := ir.FindEachIndex(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachIndex: got %v, want %v", got, want)
	}
}

//   - FindEachSubmatch:

// TestFindEachSubmatch tests all matches by
// FindEachSubmatch with no limit
func TestFindEachSubmatch(t *testing.T) {
	lim := -1
	ir := MustCompile(`(x+)(y+)`)
	b := []byte("aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy")
	want := ir.FindAllSubmatch(b, lim)
	var got [][][]byte
	it := ir.FindEachSubmatch(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachSubmatch: got %v, want %v", got, want)
	}
}

// TestFindEachSubmatch tests matches by
// FindEachSubmatch with no limit
func TestFindEachSubmatchWithLimit(t *testing.T) {
	lim := 2
	ir := MustCompile(`(x+)(y+)`)
	b := []byte("aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy")
	want := ir.FindAllSubmatch(b, lim)
	var got [][][]byte
	it := ir.FindEachSubmatch(b, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachSubmatch: got %v, want %v", got, want)
	}
}

// TestFindEachSubmatchNoMatch tests matches by
// FindEachSubmatch with no match
func TestFindEachSubmatchNoMatch(t *testing.T) {
	ir := MustCompile(`(x+)(y+)`)
	b := []byte("aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy")
	want := [][][]byte{}
	var got [][][]byte
	it := ir.FindEachSubmatch(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachSubmatch: got %v, want %v", got, want)
	}
}

//   - FindEachSubmatchIndex:

// TestFindEachSubmatchIndex tests all matches by
// FindEachSubmatchIndex with no limit
func TestFindEachSubmatchIndex(t *testing.T) {
	lim := -1
	ir := MustCompile(`(x+)(y+)`)
	b := []byte("aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy")
	want := ir.FindAllSubmatchIndex(b, lim)
	var got [][]int
	it := ir.FindEachSubmatchIndex(b)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachSubmatchIndex: got %v, want %v", got, want)
	}
}

// TestFindEachSubmatchIndexWithIndex tests matches by
// FindEachSubmatchIndex with limit
func TestFindEachSubmatchIndexWithIndex(t *testing.T) {
	lim := 2
	ir := MustCompile(`(x+)(y+)`)
	b := []byte("aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy")
	want := ir.FindAllSubmatchIndex(b, lim)
	var got [][]int
	it := ir.FindEachSubmatchIndex(b, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachSubmatchIndex: got %v, want %v", got, want)
	}
}

// TestFindEachSubmatchIndexNoMatch tests matches by
// FindEachSubmatchIndex with limit
func TestFindEachSubmatchIndexNoMatch(t *testing.T) {
	lim := 2
	ir := MustCompile(`(x+)(y+)`)
	b := []byte("aaaxyaaaaxxyyaaaaxxxyyyaaaaaxxxxxyyyyy")
	want := ir.FindAllSubmatchIndex(b, lim)
	var got [][]int
	it := ir.FindEachSubmatchIndex(b, lim)
	for e := range it {
		got = append(got, e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindEachSubmatchIndex: got %v, want %v", got, want)
	}
}
