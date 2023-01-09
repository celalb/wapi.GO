package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
	if IsEmpty(fileName) {
		return "", errors.New(" Dosya adı bos")
	}
	bytes, error := ioutil.ReadFile(fileName)
	checkError(error)
	return string(bytes), nil
}
