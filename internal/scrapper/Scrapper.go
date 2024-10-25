package scraper

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

func CollectCodebaseContent(root string) (string, error) {
    var contentBuilder strings.Builder

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            fileContent, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }
            contentBuilder.WriteString(fmt.Sprintf("\n// %s\n", path))
            contentBuilder.Write(fileContent)
            contentBuilder.WriteString("\n")
        }
        return nil
    })
    if err != nil {
        return "", err
    }
    return contentBuilder.String(), nil
}

func WriteToFile(filename, content string) error {
    return ioutil.WriteFile(filename, []byte(content), 0644)
}
