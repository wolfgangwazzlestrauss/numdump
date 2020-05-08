package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Print integer representation of bytes.
func printFile(bytes []byte, stride int) {
	length := len(bytes)

	for i := 0; i < length; i += stride {
		str := ""
		for j := 0; j < stride && i+j < length; j++ {
			str += fmt.Sprintf("% 4d", bytes[i+j])
		}

		fmt.Printf("%07d %s\n", i, str)
	}
}

// Parse command line arguments.
func readArgs() string {
	if len(os.Args) < 2 {
		panic("missing file path arugment")
	} else {
		return os.Args[1]
	}
}

// Read binary file from disk.
func readFile(file_path string) []byte {
	bytes, err := ioutil.ReadFile(file_path)

	if err != nil {
		panic(err)
	}

	return bytes
}

func main() {
	file_path := readArgs()
	bytes := readFile(file_path)
	printFile(bytes, 4)
}
