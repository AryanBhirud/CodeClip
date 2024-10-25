package scraper

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

// CollectCodebaseContent collects the contents of code files in the specified root directory
// while ignoring common non-user code directories and files.
func CollectCodebaseContent(root string) (string, error) {
    var contentBuilder strings.Builder

    // Define directories and file extensions to ignore
    ignoredDirs := map[string]bool{
        "node_modules": true,
        ".git":         true,
        ".env":         true,
        "vendor":       true,
        "__pycache__":  true,
    }
    ignoredExtensions := map[string]bool{
        ".exe":   true,
        ".bin":   true,
        ".dll":   true,
        ".log":   true,
        ".tmp":   true,
        ".lock":  true,
        ".tar":   true,
        ".gz":    true,
        ".zip":   true,
        ".png":   true,
        ".jpg":   true,
        ".jpeg":  true,
        ".gif":   true,
        ".svg":   true,
        ".pdf":   true,
    }

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories in the ignored list
        if info.IsDir() {
            if _, exists := ignoredDirs[info.Name()]; exists {
                return filepath.SkipDir
            }
            return nil
        }

        // Skip files with ignored extensions
        if ext := filepath.Ext(info.Name()); ignoredExtensions[ext] {
            return nil
        }

        // Read and append file content
        fileContent, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        contentBuilder.WriteString(fmt.Sprintf("\n// %s\n", path))
        contentBuilder.Write(fileContent)
        contentBuilder.WriteString("\n")

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
