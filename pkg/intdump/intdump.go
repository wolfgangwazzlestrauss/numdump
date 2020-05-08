package intdump

import (
	"fmt"
)

// Print integer representation of bytes.
func PrintBytes(bytes []byte, stride int) {
	length := len(bytes)

	for i := 0; i < length; i += stride {
		str := ""
		for j := 0; j < stride && i+j < length; j++ {
			str += fmt.Sprintf("% 4d", bytes[i+j])
		}

		fmt.Printf("%07d %s\n", i, str)
	}
}
