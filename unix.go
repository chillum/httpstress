// +build !windows

// This ulility takes care of `ulimit -n` on Unix systems: sets it to
// the value of `-c` option plus 6, if the current limit is smaller.
package main

import (
	"fmt"
	"os"
	"syscall"
)

// Sets Unix limits. Returns true on success, false on errors.
func setlimits(limit *int) bool {
	var old, new syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &old)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: getrlimit(2) failed:", err)
		return false
	}

	new.Cur = uint64(*limit + 6) // Magic. 1-5 does not work, 6 seems OK.
	new.Max = new.Cur
	if old.Cur < new.Cur || old.Max < new.Cur {
		err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &new)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: For testing with", *limit, "concurrent connections,")
			fmt.Fprintln(os.Stderr, "       you need the NOFILE Unix limit set to", new.Cur, "or greater.")
			fmt.Fprintln(os.Stderr, "We tried to set it, but setrlimit(2) failed:", err)
			fmt.Fprintln(os.Stderr, "Ask your system administrator to tune this limit or use sudo.")
			return false
		}
	}
	return true
}
