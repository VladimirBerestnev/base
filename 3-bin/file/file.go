package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadAnyFile(filename string) ([]byte, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func CheckJsonFile(filename string) bool {
	extension := filepath.Ext(filename)
	fmt.Println("This is json")
	return extension == ".json"
}
