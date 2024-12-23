package tasks

import (
	"encoding/json"
	"fmt"
)

// JSONStorage interface for JSON operations
type JSONStorage interface {
	Save(tasks JSONTasks) error
	Load() (JSONTasks, error)
}

// JSONTaskRepository handles JSON-specific operations
type JSONTaskRepository struct {
	Storage        FileStorage
	DefaultContent string
}

func (jsonTaskRepository *JSONTaskRepository) Save(tasks JSONTasks) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	return jsonTaskRepository.Storage.Write(string(data))
}

func (jsonTaskRepository *JSONTaskRepository) Load() (JSONTasks, error) {
	content, err := jsonTaskRepository.Storage.Read()
	if err != nil {
		if writeErr := jsonTaskRepository.Storage.Write(jsonTaskRepository.DefaultContent); writeErr != nil {
			return JSONTasks{}, fmt.Errorf("failed to write default content: %w", writeErr)
		}
		content = jsonTaskRepository.DefaultContent
	}

	var tasks JSONTasks
	err = json.Unmarshal([]byte(content), &tasks)
	if err != nil {
		return JSONTasks{}, fmt.Errorf("failed to parse JSON content: %w", err)
	}

	return tasks, nil
}
