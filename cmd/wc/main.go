package main

import (
	"bufio"
	"os"
	"unicode"
)

func main() {
	options := newOptions()

	switch {
	case options.help:
		usage()
	case options.version:
		version()
	default:
		pagesize := os.Getpagesize()
		reader := bufio.NewReaderSize(os.Stdin, pagesize)
		size, chars, words, lines, maxLineLen := count(reader)
		writeCount(size, chars, lines, words, maxLineLen, options)
	}
}

type runeReader interface {
	ReadRune() (rune, int, error)
}

func count(reader runeReader) (bytes, chars, words, lines, maxLineLen int) {
	var inWord bool
	var lineLen int

	for {
		c, s, err := reader.ReadRune()
		if err != nil {
			break
		}

		bytes += s
		chars++

		switch {
		case c == '\n':
			if maxLineLen < lineLen {
				maxLineLen = lineLen
			}
			lines++
			lineLen = 0
			inWord = false
		case unicode.IsSpace(c):
			lineLen++
			inWord = false
		case !inWord:
			words++
			lineLen++
			inWord = true
		default:
			lineLen++
		}
	}

	if maxLineLen < lineLen {
		maxLineLen = lineLen
	}

	return
}
