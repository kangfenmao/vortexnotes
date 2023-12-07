package storages

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"vortexnotes/backend/logger"
)

func CalculateFileHash(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write([]byte(filePath))
	hash.Write(fileData)
	hashValue := hash.Sum(nil)

	return fmt.Sprintf("%x", hashValue), nil
}

func ListTextFiles(dirPath string) (error, []string) {
	var textFiles []string

	// 打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		return err, textFiles
	}

	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			logger.Logger.Println("Close dir error", err)
		}
	}(dir)

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, textFiles
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

	return nil, textFiles
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
