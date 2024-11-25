# Scripts for Generating Test Data (Go Version)

This directory contains Go programs to generate test files.
These files are **not committed** to the repository to avoid bloating its size.

## Available Scripts

### 1. `generate_sparse_file.go`
Generates sparse files of a given size.
- **Usage:** `./generate_sparse_file <output_file> <size>`
- **Example:** `./generate_sparse_file sparsefile.img 10G`

### 2. `generate_large_binary.go`
Creates large binary files with random data.
- **Usage:** `./generate_large_binary <output_file> <size_in_bytes>`
- **Example:** `./generate_large_binary largefile.bin 1073741824` (1GB)
