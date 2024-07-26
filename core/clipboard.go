package core

import (
   "github.com/atotto/clipboard"
)

type ClipboardURLProvider struct{}

func (p ClipboardURLProvider) GetURL() (string, error) {
   return clipboard.ReadAll()
}

