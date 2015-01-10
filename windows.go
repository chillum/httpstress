// +build windows

package main

// Windows has no user-tunable connection limits, so no-op here.
func setlimits(limit *int) {
}
