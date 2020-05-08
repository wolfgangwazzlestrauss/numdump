package morestrings

import (
	"testing"
	"github.com/wolfgangwazzlestrauss/intdump/pkg/intdump"
)

func TestPrintBytes(t *testing.T) {
	bytes := []byte{1, 2}

	expected := "0000000   1   2\n"
	actual := intdump.DumpBytes(bytes, 4)

	if actual != expected {
		t.Errorf("\nactual = %s;\nexpected = %s", actual, expected)
	}
}
