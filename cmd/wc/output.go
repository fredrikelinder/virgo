package main

import (
	"flag"
	"fmt"
)

func newOptions() *options {
	bs := flag.Bool("c", false, "print byte count")
	bl := flag.Bool("bytes", false, "print byte count")
	cs := flag.Bool("m", false, "print character count")
	cl := flag.Bool("chars", false, "print character count")
	ls := flag.Bool("l", false, "print line count")
	ll := flag.Bool("lines", false, "print line count")
	ws := flag.Bool("w", false, "print word count")
	wl := flag.Bool("words", false, "print word count")
	ms := flag.Bool("L", false, "print max line length")
	ml := flag.Bool("max-line-length", false, "print max line length")
	h := flag.Bool("help", false, "print help")
	v := flag.Bool("version", false, "print version")

	flag.Parse()

	if *h {
		return &options{help: true}
	}

	if *v {
		return &options{version: true}
	}

	if !*bs && !*bl && !*cs && !*cl && !*ls && !*ll && !*ws && !*wl && !*ms && !*ml {
		return &options{bytes: true, words: true, lines: true}
	}

	return &options{
		bytes:      *bs || *bl,
		chars:      *cs || *cl,
		words:      *ws || *wl,
		lines:      *ls || *ll,
		maxLineLen: *ms || *ml,
	}
}

type options struct {
	lines      bool
	bytes      bool
	chars      bool
	words      bool
	maxLineLen bool
	help       bool
	version    bool
}

func usage() {
	flag.Usage()
}

func version() {
	fmt.Println(`1.0.0`)
}

func writeCount(size, chars, lines, words, maxLineLen int, opts *options) {
	if opts.lines {
		fmt.Printf("%8v", lines)
	}

	if opts.words {
		fmt.Printf("%8v", words)
	}

	if opts.bytes {
		fmt.Printf("%8v", size)
	} else if opts.chars {
		fmt.Printf("%8v", chars)
	}

	if opts.maxLineLen {
		fmt.Printf("%8v", maxLineLen)
	}

	fmt.Println()
}
