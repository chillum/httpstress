// +build !windows

// This ulility takes care of `ulimit -n` on Unix systems: sets it to
// the value of `-c` option plus 6, if the current limit is smaller.
// Warns on stderr upon errors.
package main

/* Copyright 2014 Chai Chillum

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

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
