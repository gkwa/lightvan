package core

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/atotto/clipboard"
)

type ClipboardURLProvider struct{}

func (p ClipboardURLProvider) GetURL() (string, error) {
	return clipboard.ReadAll()
}

type FileURLProvider struct {
	Path string
}

func (p FileURLProvider) GetURL() (string, error) {
	file, err := os.Open(p.Path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	urlPattern := regexp.MustCompile(`https?://[^\s]+`)

	for scanner.Scan() {
		line := scanner.Text()

		match := urlPattern.FindString(line)
		if match != "" {
			return match, nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return "", fmt.Errorf("no URL found in the file")
}
