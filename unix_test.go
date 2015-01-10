// +build !windows

// Test the setlimits() function.
package main

import "testing"

func TestSetlimits(t *testing.T) {
	if valid := 1024; setlimits(&valid) == false {
		t.Errorf("setlimit(%v) didn't work", valid)
	}
}
