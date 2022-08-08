package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// LoadFile 加载文件
func LoadFile(name string) []byte {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil
	}

	file := fmt.Sprintf("%s/%s", dir, name)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	return bytes
}
