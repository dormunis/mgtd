package taskmanager

import (
	"errors"
	"fmt"
	"mgtd/adapters"
	"mgtd/taskmanagers/todoist"
)

type TaskManagerAdapterType string

const (
	Todoist TaskManagerAdapterType = "todoist"
)

type TaskManager interface {
	Initialize(adapters.Settings) error
	FetchTasks() ([]adapters.Task, error)
	DeleteTask(adapters.Task) error
	RevalidateTask(adapters.Task) error
	DeferSomedayTask(adapters.Task) error
}

func Initialize(taskManagerAdapterType TaskManagerAdapterType, settings adapters.Settings) (adapters.TaskManagerAdapter, error) {
	var adapter adapters.TaskManagerAdapter
	switch taskManagerAdapterType {
	case Todoist:
		var e error
		adapter, e = todoist.NewTodoistAdapter()

		if e != nil {
			return nil, e
		}
	default:
		return nil, errors.New(fmt.Sprintf("Unknown adapter type: %v", taskManagerAdapterType))
	}
	adapter.Initialize(settings)
	return adapter, nil
}

func FilterTasks(tasks *[]adapters.Task, fr *adapters.FilterRequest) ([]adapters.Task, error) {
	var filtered []adapters.Task
	for _, task := range *tasks {
		if fr.Project != nil && task.Project != *fr.Project {
			continue
		}
		if fr.Status != nil && task.Status == *fr.Status {
			continue
		}
		// TOOD: implement timespan filtering
		// TOOD: implement tags filtering
		filtered = append(filtered, task)
	}
	return filtered, nil
}
