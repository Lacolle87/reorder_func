# Go Function Sorter

This project is a simple Go application designed to sort functions within a Go source file alphabetically by their names. It reads a file, extracts functions, sorts them, and then writes the sorted content back to the file.

## Usage

To use Go Function Sorter, run the compiled executable with the path to the Go source file you want to sort as an argument:
```bash
./function_sorter <file_path>
```
Replace <file_path> with the path to your Go source file.

**For example:**
```bash
./function_sorter ./path/to/your/go/file.go
```
This will sort the functions within the specified file alphabetically by their names.

## Features

- **Function Extraction**: The application scans the provided file and extracts all functions, including their bodies.
- **Sorting**: Functions are sorted alphabetically by their names.
- **In-place Editing**: The sorted content is written back to the original file, replacing its content.

## Contributing

Contributions are welcome.