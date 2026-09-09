package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadAnyFile(filename string) {

	extension := filepath.Ext(filename)
	if extension == ".json" {
		fmt.Println("Файл в формате json")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}
}
