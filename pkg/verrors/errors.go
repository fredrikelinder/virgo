package verrors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fredrikelinder/virgo/pkg/con"
)

// Errors wraps the given error with the provided format and arguments.
func Errorsf(errs []error, fmt string, args ...interface{}) error {
	var es []error

	for _, err := range errs {
		if err != nil {
			es = append(es, err)
		}
	}

	if len(es) == 0 {
		return nil
	}

	_, fn, line, _ := runtime.Caller(1)

	return &Errors{errs: es, fmt: fmt, args: args, fn: fn, line: line}
}

// Errors represents a wrapped error slice.
type Errors struct {
	errs []error
	fn   string
	line int
	fmt  string
	args []interface{}
}

func (e *Errors) Error() string {
	n := len(e.errs)
	ss := make([]string, n)

	con.For(n, func(i int) {
		ss[i] = e.errs[i].Error()
	})

	f := e.fmt + "[%s]"
	a := append(e.args, strings.Join(ss, ", "))
	return fmt.Sprintf(f, a...)
}
