package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file_path>\n", os.Args[0])
		return
	}

	filePath := os.Args[1]

	content, functions, err := readAndExtractFunctions(filePath)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}

	sortedContent := sortFunctionsInContent(content, functions)

	err = writeSortedContentToFile(filePath, sortedContent)
	if err != nil {
		fmt.Println("Error writing sorted content back to file:", err)
		return
	}

	fmt.Println("File processed and functions sorted successfully.")
}

func readAndExtractFunctions(filePath string) ([]string, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var content []string
	var functions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
		if strings.HasPrefix(strings.TrimSpace(line), "func ") {
			functions = append(functions, line)
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, nil, err
	}

	return content, functions, nil
}

func sortFunctionsInContent(content []string, functions []string) []string {
	sort.Strings(functions)
	functionMap := make(map[string]bool)
	for _, f := range functions {
		functionMap[f] = true
	}

	var sortedContent []string
	for _, line := range content {
		if _, exists := functionMap[line]; !exists {
			sortedContent = append(sortedContent, line)
		} else {
			for _, sortedFunc := range functions {
				if !functionMap[sortedFunc] {
					continue
				}
				sortedContent = append(sortedContent, sortedFunc)
				functionMap[sortedFunc] = false
				break
			}
		}
	}

	return sortedContent
}

func writeSortedContentToFile(filePath string, content []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range content {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
