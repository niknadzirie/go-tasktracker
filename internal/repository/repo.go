package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"tasktracker/internal/model"
)

type TaskRepository interface {
	ReadAllFromJsonFile() ([]model.Task, error)
	WriteAllToJsonFile(tasks []model.Task) error
}

type jsonRepository struct {
	file *os.File
}

func NewJSONRepository(file *os.File) TaskRepository {
	return &jsonRepository{file: file}
}

func (r *jsonRepository) ReadAllFromJsonFile() ([]model.Task, error) {
	var tasks []model.Task

	decoder := json.NewDecoder(r.file)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return nil, err
	}

	return tasks, nil
}

func (r *jsonRepository) WriteAllToJsonFile(tasks []model.Task) error {
	encoder := json.NewEncoder(r.file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return err
	}

	fmt.Println("Successfull encoded to JSON file")
	return nil
}
