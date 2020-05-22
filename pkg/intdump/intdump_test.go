package intdump

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestFmtStrings(t *testing.T) {
	strs := []string{"1", "2", "201", "156", "34"}

	expected := "" +
		"0000000   1   2 201 156\n" +
		"0000004  34\n"
	actual := fmtStrings(strs, 4, 4)

	assert.Equal(t, expected, actual)
}
