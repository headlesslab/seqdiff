package seqdiff_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/headlesslab/seqdiff"
)

func TestSplit(t *testing.T) {
	if out := seqdiff.Split(""); len(out) != 0 {
		t.Fatalf("Split(%q) = %q, want no words", "", out)
	}

	check := func(in string, expect ...string) {
		t.Helper()
		out := seqdiff.Split(strings.Repeat("\t", 100) + in)[100:]
		if !reflect.DeepEqual(out, expect) {
			t.Errorf("Split(%q) = %q, want %q", in, out, expect)
		}
	}

	check("find a place to eat 热干面",
		"find", " ", "a", " ", "place", " ", "to", " ", "eat", " ", "热", "干", "面")

	check("	as.Equal(arr, arr) test",
		"	", "as", ".", "Equal", "(", "arr", ",", " ", "arr", ")", " ", "test")

	check("English-Words紧接着中文",
		"English", "-", "Words", "紧", "接", "着", "中", "文")

	check("123456test12345",
		"123", "456", "test", "123", "45")

	check("WordVeryVeryVeryVeryVeryVeryVerylong",
		"WordVeryVery", "VeryVeryVery", "VeryVerylong")
}

func TestIsSubsequenceOf(t *testing.T) {
	y := seqdiff.NewChars("abc")

	for _, x := range []string{"ab", "ac", "bc"} {
		if !seqdiff.NewChars(x).IsSubsequenceOf(y) {
			t.Errorf("%q should be a subsequence of %q", x, "abc")
		}
	}
	for _, x := range []string{"cb", "ba", "ca"} {
		if seqdiff.NewChars(x).IsSubsequenceOf(y) {
			t.Errorf("%q should not be a subsequence of %q", x, "abc")
		}
	}
}

func TestNew(t *testing.T) {
	lenIs := func(s seqdiff.Sequence, n int) {
		t.Helper()
		if len(s) != n {
			t.Errorf("len(%q) = %d, want %d", s.String(), len(s), n)
		}
	}

	lenIs(seqdiff.NewLines("a"), 1)
	lenIs(seqdiff.NewLines("a\n"), 2)
	lenIs(seqdiff.NewLines("a\n\n"), 3)
	lenIs(seqdiff.NewLines("\na"), 2)
	if got := seqdiff.NewLines("\nabc\nabc").String(); got != "\nabc\nabc" {
		t.Errorf("NewLines(...).String() = %q, want %q", got, "\nabc\nabc")
	}

	lenIs(seqdiff.NewWords([]string{"a", "b"}), 2)
}
