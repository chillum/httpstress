// +build !windows

// This ulility takes care of `ulimit -n` on Unix systems: sets it to
// the value of `-c` option plus 6, if the current limit is smaller.
// Warns on stderr upon errors.
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
		fmt.Fprintln(os.Stderr, "Unable to access Unix limits:", err)
		return false
	} else {
		new.Cur = uint64(*limit + 6) // Magic. 1-5 does not work, 6 seems OK.
		new.Max = new.Cur
		if old.Cur < new.Cur || old.Max < new.Cur {
			err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &new)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Unable to set Unix limits:", err)
				return false
			}
		}
	}
	return true
}
