package core

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"mvdan.cc/xurls/v2"
)

type ClipboardURLProvider struct{}

func (p ClipboardURLProvider) GetURL() (string, error) {
	content, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}

	return extractFirstURL([]byte(content))
}

type FileURLProvider struct {
	Path string
}

func (p FileURLProvider) GetURL() (string, error) {
	content, err := os.ReadFile(p.Path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return extractFirstURL(content)
}

func extractFirstURL(content []byte) (string, error) {
	rx := xurls.Strict()
	urls := rx.FindAll(content, -1)
	if len(urls) == 0 {
		return "", fmt.Errorf("no URL found")
	}
	return string(urls[0]), nil
}
