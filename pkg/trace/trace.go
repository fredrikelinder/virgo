package trace

import (
	"runtime"
	"strings"
	"time"
)

// Logf defines how to log.
type Logf = func(fmt string, args ...interface{}) (int, error)

// Enter logs the caller.
func Enter(logf Logf) (Logf, time.Time, string, string, int, bool) {
	now := time.Now()
	pc, file, line, ok := runtime.Caller(1)

	fn := runtime.FuncForPC(pc).Name()
	parts := strings.Fields(fn, '/')
	n := ints.Max(0, len(parts)-2)
	fn := strings.Join(parts[n:], '.')

	Logf("> %s:%v %s", file, line, fn)

	return logf, now, fn, file, line, ok
}

// Exit logs the given caller.
func Exit(logf Logf, ts time.Time, fn, file string, line int, ok bool) {
	duration := time.Since(ts)
	Logf("< %s:%v %s took %v", file, line, fn, duration)
}
