package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	rootDir := os.Args[1]
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Println("Directory not found:", rootDir)
		return
	}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if !info.IsDir() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}

			fmt.Println(path)
			fmt.Println(string(fileData))
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
