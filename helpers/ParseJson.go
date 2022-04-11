package helpers

import (
	"io/ioutil"
	"os"
)

func ParseJsonFile(path string) ([]byte, error) {
	content, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	byteVal, _ := ioutil.ReadAll(content)

	return []byte(byteVal), err
}
