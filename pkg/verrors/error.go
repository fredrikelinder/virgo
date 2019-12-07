package verrors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fredrikelinder/virgo/pkg/ints"
)

// Errorf wraps the given error with the provided format and arguments.
// Use "%@" to insert "<package>/<file>:<line>" in the output of Error.
// Nil is returned if the given err is nil.
func Errorf(err error, fmt string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	_, fn, line, _ := runtime.Caller(1)

	return &Error{err: err, fmt: fmt, args: args, fn: fn, line: line}
}

// Error represents a wrapped error.
type Error struct {
	err      error
	fn       string
	line     int
	fmt      string
	args     []interface{}
	replaced bool
}

func (e *Error) Error() string {
	e.replaceProcentAt()
	f := e.fmt + "%s"
	a := append(e.args, e.err)
	return fmt.Sprintf(f, a...)
}

// replaceProcentAt replaces "%@" with "<package>/<file>:<line>".
func (e *Error) replaceProcentAt() {
	if e.replaced {
		return
	}

	e.replaced = true
	var i int
	var fmt []rune
	var wasProcent bool

	for _, c := range e.fmt {
		switch {
		case c == '%' && !wasProcent:
			wasProcent = true
		case c == '%' && wasProcent:
			fmt = append(fmt, '%', '%')
			wasProcent = false
		case c == '@' && wasProcent:
			e.args = append(e.args, nil)
			copy(e.args[i+1:], e.args[i:])
			e.args[i] = packageFileAndLine(e.fn, e.line)
			fmt = append(fmt, '%', 's')
			i++
			wasProcent = false
		case c != '%' && wasProcent:
			fmt = append(fmt, '%', c)
			i++
			wasProcent = false
		default:
			fmt = append(fmt, c)
			wasProcent = false
		}
	}

	e.fmt = string(fmt)
}

func packageFileAndLine(fn string, line int) string {
	parts := strings.FieldsFunc(fn, func(r rune) bool { return r == '/' })
	n := ints.IntsMax(0, len(parts)-2)
	return fmt.Sprintf("%s/%s:%v", parts[n], parts[n+1], line)
}
