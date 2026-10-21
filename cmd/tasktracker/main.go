package main

import (
	"fmt"
	"os"
	"tasktracker/internal/helper"
	"tasktracker/internal/model"
	"tasktracker/internal/repository"
	"tasktracker/internal/service"
)

func main() {
	const jsonFile string = "tasktracker.json"

	var file *os.File
	var err error

	if !helper.CheckFileExist(jsonFile) {
		err = os.WriteFile(jsonFile, []byte("[]"), 0644)
		if err != nil {
			fmt.Println("Error write file: ", err)
			return
		}
	}

	file, err = os.OpenFile(jsonFile, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	repo := repository.NewJSONRepository(file)
	svc := service.NewTaskSerivce(repo)

	svc.AddTask("test2", model.StatusPending)

	svc.TestDisplay()

}
