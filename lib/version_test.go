package lib

import "testing"

func TestVersion(t *testing.T) {
	v := Version()
	if v != "0.2.8" {
		t.Errorf(msgFail, "Version", "0.2.8", v)
	}
}
