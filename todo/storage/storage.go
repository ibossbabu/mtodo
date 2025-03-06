package storage

import (
    "encoding/json"
    "os"
)

type Task struct {
    Text    string `json:"text"`
    Checked bool   `json:"checked"`
}

const DataFile = "todos.json"

// LoadTasks reads tasks from the JSON file
func LoadTasks() ([]Task, error) {
    file, err := os.OpenFile(DataFile, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    tasks := []Task{}
    stat, err := file.Stat()
    if err != nil {
        return nil, err
    }

    if stat.Size() > 0 {
        decoder := json.NewDecoder(file)
        if err := decoder.Decode(&tasks); err != nil {
            return nil, err
        }
    }

    return tasks, nil
}

// SaveTasks writes tasks to the JSON file with indented formatting
func SaveTasks(tasks []Task) error {
    file, err := os.Create(DataFile)
    if err != nil {
        return err
    }
    defer file.Close()

    // Marshal tasks with indentation
    data, err := json.MarshalIndent(tasks, "", "    ")
    if err != nil {
        return err
    }

    _, err = file.Write(data)
    return err
}
