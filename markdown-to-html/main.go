package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func Convert(content string, scanner *bufio.Scanner) {
	var builder strings.Builder

	//Remove whitespaces
	line := bytes.TrimSpace(scanner.Bytes())

	if string(line[0]) == "#" {
		// Count how many hash signs there are
		count := bytes.Count(line, []byte("#"))

		switch count {
		case 1:
			str := strings.Replace(content, "#", "", -1)
			builder.WriteString(fmt.Sprintf("<h1>%s</h1>", strings.TrimSpace(str)))
		}
		fmt.Println(builder.String())
	}
}

func main() {
	fileName := "test.md"
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Unable to open %s: %v", fileName, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan the content line by line
	for scanner.Scan() {
		// Save the current line
		content := scanner.Text()

		if strings.Contains(content, "#") {
			Convert(content, scanner)
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("Unable to scan the file: %v", err)
	}
}
