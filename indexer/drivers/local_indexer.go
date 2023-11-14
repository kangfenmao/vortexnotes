package drivers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LocalIndexer struct {
}

func (local LocalIndexer) ListAllNotes() []string {
	const path = "./data/vortexnotes"

	err := createDirectoryIfNotExists(path)
	if err != nil {
		var list []string
		fmt.Println("Error:", err)
		return list
	}

	var notes, _ = listTextFiles(path)
	return notes
}

func (local LocalIndexer) IndexExist(path string) bool {
	fmt.Println("IndexExist: ", path)
	return false
}

func (local LocalIndexer) ParseNote(path string) string {
	fmt.Println("ParseNote: ", path)
	return ""
}

func (local LocalIndexer) AddNoteToIndex(path string, note string) bool {
	fmt.Println("AddNoteToIndex: ", path)
	return true
}

func listTextFiles(dirPath string) ([]string, error) {
	var textFiles []string

	// 打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(fileInfo.Name()))

		if ext == ".txt" || ext == ".md" {
			filePath := filepath.Join(dirPath, fileInfo.Name())
			textFiles = append(textFiles, filePath)
		}
	}

	return textFiles, nil
}

func createDirectoryIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
