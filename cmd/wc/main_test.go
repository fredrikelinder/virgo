package main

import (
	"io"
	"testing"
	"unicode/utf8"
)

func TestCount(t *testing.T) {
	cases := []struct {
		words int
		lines int
		maxLi int
		value string
	}{
		{},
		{words: 1, lines: 0, maxLi: 1, value: "h"},
		{words: 1, lines: 0, maxLi: 3, value: " h "},
		{words: 1, lines: 1, maxLi: 2, value: "\nh "},
		{words: 1, lines: 1, maxLi: 3, value: " h \n"},
		{words: 1, lines: 1, maxLi: 3, value: " ♥ \n"},
		{words: 4, lines: 2, maxLi: 10, value: " one two \nthree four\n"},
	}

	for i, c := range cases {
		c_bytes := len(c.value)
		c_chars := utf8.RuneCountInString(c.value)
		bytes, chars, words, lines, maxLi := count(newRuneReaderMock(c.value))
		switch {
		case c_bytes != bytes:
			t.Errorf("%v: expected bytes %v, got %v", i, c_bytes, bytes)
		case c_chars != chars:
			t.Errorf("%v: expected chars %v, got %v", i, c_chars, chars)
		case c.words != words:
			t.Errorf("%v: expected words %v, got %v", i, c.words, words)
		case c.lines != lines:
			t.Errorf("%v: expected lines %v, got %v", i, c.lines, lines)
		case c.maxLi != maxLi:
			t.Errorf("%v: expected max line length %v, got %v", i, c.maxLi, maxLi)
		}
	}
}

func newRuneReaderMock(value string) *runeReaderMock {
	if value == "" {
		return &runeReaderMock{}
	}

	var runes []rune
	var positions []int

	for p, r := range value {
		runes = append(runes, r)
		positions = append(positions, p)
	}

	var sizes []int
	var prevPos int
	for _, p := range positions[1:] {
		sizes = append(sizes, p-prevPos)
		prevPos = p
	}
	sizes = append(sizes, len(value)-prevPos)

	return &runeReaderMock{runes: runes, sizes: sizes}
}

type runeReaderMock struct {
	runes []rune
	sizes []int
}

func (rr *runeReaderMock) ReadRune() (r rune, s int, err error) {
	if len(rr.runes) == 0 {
		return r, s, io.EOF
	}

	r, rr.runes = rr.runes[0], rr.runes[1:]
	s, rr.sizes = rr.sizes[0], rr.sizes[1:]

	return r, s, nil
}
