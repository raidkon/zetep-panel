package executil

import (
	"errors"
	"os/exec"
)

// ExitCode returns (code, true) if err is *exec.ExitError.
func ExitCode(err error) (int, bool) {
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		return ee.ExitCode(), true
	}
	return 0, false
}
