package intdump

import "fmt"

// Print integer representation of bytes.
func fmtInts(bytes []byte, stride int) string {
	length := len(bytes)

	output := ""
	for i := 0; i < length; i += stride {
		str := ""
		for j := 0; j < stride && i+j < length; j++ {
			str += fmt.Sprintf("% 4d", bytes[i+j])
		}

		output += fmt.Sprintf("%07d%s\n", i, str)
	}

	return output
}
