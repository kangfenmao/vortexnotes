package local

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"vortex-notes/indexer/logger"
)

func CalculateFileHash(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := md5.Sum(fileData)
	return fmt.Sprintf("%x", hash), nil
}

func IsFileContentChanged(filePath string, previousHash string) (bool, error) {
	currentHash, err := CalculateFileHash(filePath)
	if err != nil {
		return false, err
	}

	return currentHash != previousHash, nil
}

func ListTextFiles(dirPath string) ([]string, error) {
	var textFiles []string

	// 打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			logger.Logger.Println("Close dir error", err)
		}
	}(dir)

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

func CreateDirectoryIfNotExists(dirPath string) error {
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
