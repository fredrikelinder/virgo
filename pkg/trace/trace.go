package trace

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fredrikelinder/virgo/pkg/ints"
)

// Logf defines how to log.
type Logf = func(fmt string, args ...interface{}) (int, error)

// Enter logs the caller.
func Enter(logf Logf) (Logf, string, bool) {
	pc, file, line, ok := runtime.Caller(1)
	file = formatFile(file)

	fn := runtime.FuncForPC(pc).Name()
	fn = formatFunc(fn)

	s := fmt.Sprintf("%s:%v %s\n", file, line, fn)

	logf("> %s", s)

	return logf, s, ok
}

// Exit logs the given caller.
func Exit(logf Logf, s string, ok bool) {
	logf("< %s", s)
}

func formatFile(file string) string {
	parts := strings.FieldsFunc(file, func(r rune) bool { return r == '/' })
	n := ints.Max(0, len(parts)-2)
	return strings.Join(parts[n:], "/")
}

func formatFunc(fn string) string {
	parts := strings.FieldsFunc(fn, func(r rune) bool { return r == '.' })
	n := ints.Max(0, len(parts)-1)
	return strings.Join(parts[n:], "/")
}
