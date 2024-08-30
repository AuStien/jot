package debug

import (
	"fmt"
	"runtime"
)

// WithFrame adds the caller file and line to the error
// (if it can be found).
func WithFrame(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}

	return fmt.Errorf("%w (%s:%d)", err, file, line)
}
