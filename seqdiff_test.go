package seqdiff_test

import (
	"bytes"
	"context"
	"crypto/rand"
	mrand "math/rand"
	"runtime/debug"
	"testing"
	"time"

	"github.com/headlesslab/seqdiff"
)

func TestLCS(t *testing.T) {
	check := func(i int, x, y string) {
		t.Helper()

		e := func(msg ...interface{}) {
			t.Helper()
			t.Log(i, x, y)
			t.Error(msg...)
			t.FailNow()
		}

		defer func() {
			err := recover()
			if err != nil {
				debug.PrintStack()
				e(err)
			}
		}()

		xs, ys := seqdiff.NewChars(x), seqdiff.NewChars(y)

		s := xs.Sub(xs.YadLCS(context.Background(), ys))
		out := s.String()
		expected := seqdiff.StandardLCS(xs, ys).String()

		if !s.IsSubsequenceOf(xs) {
			e(s.String(), "is not subsequence of", x)
		}

		if !s.IsSubsequenceOf(ys) {
			e(s.String(), "is not subsequence of", y)
		}

		if len(out) != len(expected) {
			e("length of", out, "doesn't equal length of", expected)
		}
	}

	randStr := func() string {
		const c = 8
		b := make([]byte, c)
		for i := 0; i < c; i++ {
			b[i] = byte('a' + mrand.Intn(c))
		}
		return string(b)
	}

	check(0, "", "")
	check(0, "", "a")

	for i := 1; i < 1000; i++ {
		check(i, randStr(), randStr())
	}
}

func TestLCSLongContentSmallChange(t *testing.T) {
	eq := func(x, y, expected string) {
		t.Helper()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		xs, ys := seqdiff.NewChars(x), seqdiff.NewChars(y)
		idx := xs.YadLCS(ctx, ys)
		out := xs.Sub(idx).String()

		if out != expected {
			t.Error(out, "!=", expected)
		}
	}

	x := bytes.Repeat([]byte("x"), 100000)
	y := bytes.Repeat([]byte("y"), 100000)
	eq(string(x), string(y), "")

	x[len(x)/2] = byte('a')
	y[len(y)/2] = byte('a')
	eq(string(x), string(y), "a")

	x[len(x)/2] = byte('y')
	y[len(y)/2] = byte('x')
	eq(string(x), string(y), "xy")
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	l := seqdiff.NewChars("abc").YadLCS(ctx, seqdiff.NewChars("abc"))
	if len(l) != 0 {
		t.Fatalf("a cancelled context should yield no indices, got %v", l)
	}
}

func TestLongRandom(_ *testing.T) {
	size := 10000
	x := randStr(size)
	y := randStr(size)

	c := context.Background()

	xs := seqdiff.NewChars(x)
	ys := seqdiff.NewChars(y)
	xs.YadLCS(c, ys)
}

func randStr(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return string(b)
}
