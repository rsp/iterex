package iterex

import (
	"reflect"
	"testing"
)

// Internal tests:

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

// Tests working on strings:

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

// Tests working on byte slices:

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
