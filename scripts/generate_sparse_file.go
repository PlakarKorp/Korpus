package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: generate_sparse_file <output_file> <size>")
		fmt.Println("Example: generate_sparse_file sparsefile.img 10G")
		os.Exit(1)
	}

	outputFile := os.Args[1]
	sizeStr := os.Args[2]

	size, err := parseSize(sizeStr)
	if err != nil {
		fmt.Printf("Error parsing size: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := file.Truncate(size); err != nil {
		fmt.Printf("Error truncating file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sparse file created: %s (%d bytes)\n", outputFile, size)
}

func parseSize(sizeStr string) (int64, error) {
	var multiplier int64 = 1
	switch sizeStr[len(sizeStr)-1] {
	case 'K', 'k':
		multiplier = 1024
	case 'M', 'm':
		multiplier = 1024 * 1024
	case 'G', 'g':
		multiplier = 1024 * 1024 * 1024
		sizeStr = sizeStr[:len(sizeStr)-1]
	}
	baseSize, err := strconv.ParseInt(sizeStr, 10, 64)
	return baseSize * multiplier, err
}
