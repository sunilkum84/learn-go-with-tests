package utils

import (
	"io/ioutil"
	"os"
)

// DoesFileExist checks if a file exists.
func DoesFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// EnsureExists creates if the file does not already exist.
func EnsureExists(fileName string) {
	if DoesFileExist(fileName) {
		return
	}

	ioutil.WriteFile(fileName, []byte(""), 0644)
}
