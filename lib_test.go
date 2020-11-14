// Test the httpstress.Test() function.
package main

import (
	"testing"

	httpstress "github.com/chillum/httpstress/lib"
)

func TestHttpStressTest(t *testing.T) {
	up := []string{"https://google.com", "http://google.com"} // These URLs should pass.
	down := []string{"http://localhost:1234"}                 // This should fail.
	invalid := []string{"ftp://localhost"}                    // Error. Non HTTP/HTTPS URL.

	if _, err := httpstress.Test(1, 1, invalid); err == nil {
		t.Errorf("%s ok (should be an error)", invalid)
	}

	if err, _ := httpstress.Test(1, 1, up); len(err) > 0 {
		t.Errorf("%s down (should be up)", up)
	}

	if err, _ := httpstress.Test(1, 10, up); len(err) > 0 {
		t.Errorf("%s down in consecutive tests (should be up)", up)
	}

	if err, _ := httpstress.Test(1, 1, down); len(err) == 0 {
		t.Errorf("%s up (should be down)", down)
	}
}
