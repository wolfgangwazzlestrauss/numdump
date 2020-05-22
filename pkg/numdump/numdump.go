// Package intdump provides functions for reading numbers from bytes and
// formatting them to resemble hexdump output.
package numdump

import (
	"fmt"
	"encoding/binary"
)

// Format and align strings to resemble hexdump output.
func fmtStrings(strs []string, width uint8, stride int) string {
	length := len(strs)
	format := fmt.Sprintf("%% %ds", width)

	text := ""
	for i := 0; i < length; i += stride {
		line := ""
		for j := 0; j < stride && i+j < length; j++ {
			line += fmt.Sprintf(format, strs[i+j])
		}

		text += fmt.Sprintf("%07d%s\n", i, line)
	}

	return text
}

// Read and format 8-bit integers.
func DumpBytes(bytes []byte, stride int) string {
	strs := make([]string, len(bytes))
	for idx, byte_ := range bytes {
		strs[idx] = fmt.Sprintf("%d", byte_)
	}

	return fmtStrings(strs, 4, stride)
}

// Read and format unsigned 16-bit integers.
func DumpShorts(bytes []byte, stride int, litteEndian bool) string {
	n := len(bytes)
	var m int
	if n % 2 == 0 {
		m = n
	} else {
		m = n - 1
	}

	strs := make([]string, n / 2)

	for idx := 0; idx < m; idx += 2 {
		var short uint16
		if litteEndian {
			short = binary.LittleEndian.Uint16(bytes[idx:idx+2])
		} else {
			short = binary.BigEndian.Uint16(bytes[idx:idx+2])
		}
		strs[idx / 2] = fmt.Sprintf("%d", short)
	}

	return fmtStrings(strs, 8, stride)
}
