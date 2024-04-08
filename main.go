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

	functions, err := readAndExtractFunctions(filePath)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}

	sortedContent := sortFunctionsInContent(functions)

	err = writeSortedContentToFile(filePath, sortedContent)
	if err != nil {
		fmt.Println("Error writing sorted content back to file:", err)
		return
	}

	fmt.Println("File processed and functions sorted successfully.")
}

func readAndExtractFunctions(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	functionMap := make(map[string][]string)
	currentFunction := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "func ") {
			currentFunction = line
		}
		functionMap[currentFunction] = append(functionMap[currentFunction], line)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return functionMap, nil
}

func sortFunctionsInContent(functionMap map[string][]string) []string {
	var functionNames []string
	for name := range functionMap {
		functionNames = append(functionNames, name)
	}
	sort.Strings(functionNames)

	var sortedContent []string
	for _, name := range functionNames {
		sortedContent = append(sortedContent, functionMap[name]...)
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
