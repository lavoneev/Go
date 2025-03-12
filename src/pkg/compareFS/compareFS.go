package FScomparer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadFiles(filename string) (map[string]bool, error) {
	files := make(map[string]bool)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		files[str] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func CompareFiles(oldFilename, newFilename string) {
	oldFiles, err := LoadFiles(oldFilename)
	if err != nil {
		fmt.Println(err)
		return
	}

	newFile, err := os.Open(newFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)

	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		if _, exists := oldFiles[str]; !exists {
			fmt.Println("ADDED", str)
		} else {
			oldFiles[str] = false
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for str, flag := range oldFiles {
		if flag {
			fmt.Println("REMOVED", str)
		}
	}
}
