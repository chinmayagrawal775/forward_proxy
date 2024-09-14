package utils

import (
	"bufio"
	"log"
	"os"
)

func LoadFile(filepath string, des *[]string) {
	fileData := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Error opening file: at path- %s: %v", filepath, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fileData = append(fileData, line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: at path- %s: %v", filepath, err)
		os.Exit(1)
	}

	*des = fileData
}
