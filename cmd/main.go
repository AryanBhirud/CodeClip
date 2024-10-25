package main

import (
    "fmt"
    "codeclip/internal/scrapper"
)

func main() {
    rootDir := "./"
    outputFile := "codebase_dump.txt"

    fmt.Println("Scraping codebase...")
    content, err := scraper.CollectCodebaseContent(rootDir)
    if err != nil {
        fmt.Printf("Error collecting codebase: %v\n", err)
        return
    }

    err = scraper.WriteToFile(outputFile, content)
    if err != nil {
        fmt.Printf("Error writing to file: %v\n", err)
        return
    }

    fmt.Printf("Codebase content saved to %s\n", outputFile)

    err = scraper.CopyToClipboard(content)
    if err != nil {
        fmt.Printf("Error copying to clipboard: %v\n", err)
    } else {
        fmt.Println("Codebase content copied to clipboard successfully.")
    }
}
