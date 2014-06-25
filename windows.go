// +build windows

package main

func setlimits(limit *int) {
	// Windows has no user-tunable connection limits, so no-op here.
}
