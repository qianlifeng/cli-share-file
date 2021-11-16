package util

import (
	"fmt"
	"os"
	"path"
)

func GetAppFolder() string {
	cwd, err := os.Getwd()
	if err == nil {
		return path.Join(cwd, "csf")
	}

	panic("can't find app dir")
}

func GetUploadFolder() string {
	return path.Join(GetAppFolder(), "uploads")
}

func EnsureAppFolderExist() {
	createFolderIfNotExist(GetAppFolder())
	createFolderIfNotExist(GetUploadFolder())
}

func createFolderIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(path, 0777)
		if mkdirErr != nil {
			panic(fmt.Sprintf("create folder %s failed：%s", path, mkdirErr.Error()))
		}
	}
}
