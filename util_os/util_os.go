package util_os

import (
	"log"
	"os"
	"path/filepath"
)

func TraverseFileSystem(targetPath string) []string {
	var files []string
	err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}
