package utils

import (
	"errors"
	"strings"
)

func GetFileExtension(name string) (string, error) {
	arr := strings.Split(name, ".")

	if len(arr) == 0 {
		return "", errors.New("file is not readable")
	}

	return arr[len(arr)-1], nil
}

func PostFileExtensionValidate(ext string) bool {
	if ext == "png" || ext == "jpg" || ext == "pdf" || ext == "mp4" || ext == "mp3" {
		return true
	}

	return false
}
