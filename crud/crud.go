package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type Todos []Todo

func (t *Todos) AddTask(Title string, Description string) {
	todo := Todo{
		Title:  Title,
		Status: false,
	}
	*t = append(*t, todo)
}

func (t *Todos) RemoveTask(id int) error {
	if len(*t) <= 0 {
		return errors.New("error: task list empty")
	}
	if id < 0 || id >= len(*t) {
		return errors.New("error: id not found")
	}
	fmt.Println("arr", (*t)[:id], (*t)[id+1:len(*t)])
	*t = append((*t)[:id], (*t)[id+1:len(*t)]...)
	return nil
}

func (t *Todos) MarkCompleted(id int) error {
	if len(*t) <= 0 {
		return errors.New("error: task list empty")
	}
	if id < 1 || id >= len(*t) {
		return errors.New("error: id not found")
	}
	if (*t)[id].Status == false {
		(*t)[id].Status = true
	} else {
		return errors.New("error: task already completed")
	}
	return nil
}

func (t *Todos) LoadJSON(fileName string) error {
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonData, t); err != nil {
		return err
	}
	return nil
}

func (t *Todos) SaveJSON(fileName string) error {
	jsonData, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return nil
	}
	return nil
}
