package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: generate_large_binary <output_file> <size_in_bytes>")
		fmt.Println("Example: generate_large_binary largefile.bin 1073741824")
		os.Exit(1)
	}

	outputFile := os.Args[1]
	size, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Printf("Invalid size: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	buffer := make([]byte, 1024*1024) // 1MB buffer
	var written int64
	for written < size {
		toWrite := size - written
		if toWrite > int64(len(buffer)) {
			toWrite = int64(len(buffer))
		}

		_, err := rand.Read(buffer[:toWrite])
		if err != nil {
			fmt.Printf("Error generating random data: %v\n", err)
			os.Exit(1)
		}

		_, err = file.Write(buffer[:toWrite])
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}

		written += toWrite
	}

	fmt.Printf("Binary file created: %s (%d bytes)\n", outputFile, size)
}
