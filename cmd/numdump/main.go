package main

import (
	"fmt"
	"github.com/wolfgangwazzlestrauss/numdump/pkg/numdump"
	"io/ioutil"
	"os"
)

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
	text := numdump.DumpUInt16(bytes, 4, numdump.LittleEndian)
	fmt.Printf(text)
}
