package intdump

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestFmtInts(t *testing.T) {
	bytes := []byte{1, 2, 201, 156, 34}

	expected := "" +
		"0000000   1   2 201 156\n" +
		"0000004  34\n"
	actual := fmtInts(bytes, 4)

	assert.Equal(t, expected, actual)
}
