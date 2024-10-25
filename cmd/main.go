package main

import (
    "fmt"
    "codeclip/internal/scrapper"
)

func main() {
    rootDir := "./" // specify the directory to scrape
    outputFile := "codebase_dump.txt"

    fmt.Println("Scraping codebase...")

    // Collect and save the codebase
    content, err := scraper.CollectCodebaseContent(rootDir)
    if err != nil {
        fmt.Printf("Error collecting codebase: %v\n", err)
        return
    }

    // Write content to file
    err = scraper.WriteToFile(outputFile, content)
    if err != nil {
        fmt.Printf("Error writing to file: %v\n", err)
        return
    }

    fmt.Printf("Codebase content saved to %s\n", outputFile)

    // Copy content to clipboard
    err = scraper.CopyToClipboard(content)
    if err != nil {
        fmt.Printf("Error copying to clipboard: %v\n", err)
    } else {
        fmt.Println("Codebase content copied to clipboard successfully.")
    }
}
