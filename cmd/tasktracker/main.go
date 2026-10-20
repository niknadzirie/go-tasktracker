package main

import (
	"errors"
	"fmt"
	"os"
	"tasktracker/internal/repository"
)

func main() {
	const jsonFile string = "./tasktracker.json"

	var file *os.File
	var err error

	if fileExist(jsonFile) {
		fmt.Println("file exist")
		file, err = os.Open(jsonFile)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
	} else {
		fmt.Println("File not exist")
		file, err = os.Create(jsonFile)
		if err != nil {
			fmt.Println("Error creating file: ", err)
			return
		}
	}

	defer file.Close()

	repo := repository.NewJSONRepository(file)

}

func fileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}
