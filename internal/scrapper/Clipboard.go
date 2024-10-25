package scraper

import "github.com/atotto/clipboard"

func CopyToClipboard(content string) error {
    return clipboard.WriteAll(content)
}
